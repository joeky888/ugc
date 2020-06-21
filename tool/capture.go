package tool

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"regexp"
	"strings"
	"sync"
	"syscall"

	"github.com/mattn/go-colorable"
)

type Conf struct {
	Regex  *regexp.Regexp
	Colors []string
}

// genRegexGroup Generate regex group string
// genRegexGroup(1) = panic()
// genRegexGroup(2) = "$1$2"
// genRegexGroup(5) = "$1$2$3$4$5"
func genRegexGroup(n int) string {
	if n <= 1 {
		log.Fatalln("genRegexGroup is only used for regex group case")
	}
	rg := ""
	for i := 1; i < n+1; i++ {
		rg += fmt.Sprintf("$%d", i)
	}
	return rg
}

// copyAndCapture is a modified version
// of https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html
func copyAndCapture(w io.Writer, r io.Reader, config []Conf) error {
	reader := bufio.NewReader(r)
	for {
		// Read line by line
		buf, err := reader.ReadBytes('\n')
		if len(buf) == 0 && err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				return nil
			}
			return err
		}

		for _, conf := range config {
			// Case 1: Normal regex case
			if len(conf.Colors) == 1 {
				buf = conf.Regex.ReplaceAllFunc(buf, func(m []byte) []byte {
					return []byte(fmt.Sprintf(conf.Colors[0], conf.Regex.Find(m)))
				})
				continue
			}

			// Case 2: Regex group case
			color := genRegexGroup(len(conf.Colors))
			for i := range conf.Colors {
				// Replace $1 with RedBegin$1RedEnd etc.
				sign := fmt.Sprintf("$%d", i+1)
				color = strings.Replace(
					color,
					sign,
					fmt.Sprintf(conf.Colors[i], sign),
					1,
				)
			}
			buf = conf.Regex.ReplaceAll(buf, []byte(color))
		}
		if _, err := w.Write(buf); err != nil {
			log.Fatalf("bufn Write() with error %v\n", err)
		}
	}
	// never reached
	// panic(true)
	// return nil
}

func CaptureWorker(config []Conf) {
	bin := os.Args[1]
	arg := os.Args[2:]
	cmd := exec.Command(bin, arg...)
	ctrlc := make(chan os.Signal)
	signal.Notify(ctrlc, os.Interrupt, syscall.SIGTERM)

	stdoutIn, outpipeErr := cmd.StdoutPipe()
	if outpipeErr != nil {
		log.Fatalf("cmd.StdoutPipe() outpipeErr failed with %v\n", outpipeErr)
	}
	stderrIn, errpipeErr := cmd.StderrPipe()
	if errpipeErr != nil {
		log.Fatalf("cmd.StderrPipe() errpipeErr failed with %v\n", errpipeErr)
	}

	var wg sync.WaitGroup
	var errStdout, errStderr error
	colorStdout := colorable.NewColorableStdout()
	colorStderr := colorable.NewColorableStderr()

	defer func() {
		if err := stdoutIn.Close(); err != nil {
			log.Fatalf("stdoutIn.Close() failed with %v\n", err)
		}
	}()

	wg.Add(2)
	go func() {
		// errStdout = copyAndCapture(os.Stdout, stdoutIn, config)
		errStdout = copyAndCapture(colorStdout, stdoutIn, config)
		wg.Done()
	}()
	go func() {
		// errStderr = copyAndCapture(os.Stderr, stderrIn, config)
		errStderr = copyAndCapture(colorStderr, stderrIn, config)
		wg.Done()
	}()

	go func() {
		// User sends ctrl-c to the program
		<-ctrlc
		// Use Process.Release() instead of Process.Kill()
		// Release() waits until the cmd exit
		// Kill() Does not wait
		if err := cmd.Process.Release(); err != nil {
			log.Fatalf("failed to kill process: %v\n", err)
		}
	}()

	if err := cmd.Start(); err != nil {
		log.Fatalf("cmd.Start() failed with %v\n", err)
	}

	wg.Wait()
	statusCode := 0
	if err := cmd.Wait(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				statusCode = status.ExitStatus()
			}
		} else {
			statusCode = -1
		}
		// log.Fatalf("cmd.Wait() failed with %v\n", err)
	}
	if errStdout != nil || errStderr != nil {
		log.Fatalf("failed to capture stdout or stderr std err, %v\n, %v\n", errStdout, errStderr)
	}

	os.Exit(statusCode)
}
