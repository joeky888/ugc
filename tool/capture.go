package tool

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

// copyAndCapture is a modified version
// of https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html
func copyAndCapture(w io.Writer, r io.Reader, config []Conf) error {
	buf := make([]byte, 20480) // 20K buffer to read from stdout/stderr
	for {
		n, err := r.Read(buf[:])
		if n == 0 && err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				// err = nil
				return nil
			}
			return err
		}

		bufn := buf[:n]
		for _, conf := range config {
			// Normal regex case
			if len(conf.Colors) == 1 {
				bufn = conf.Regex.ReplaceAllFunc(bufn, func(m []byte) []byte {
					return []byte(fmt.Sprintf(conf.Colors[0], conf.Regex.Find(m)))
				})
				continue
			}
			// Regex group case
			color := conf.RegexReplace
			for i := range conf.Colors {
				sign := fmt.Sprintf("$%d", i+1)
				// Replace $1 with RedBegin$1RedEnd etc.
				color = strings.Replace(
					color,
					sign,
					fmt.Sprintf(conf.Colors[i], sign),
					1,
				)
			}
			bufn = conf.Regex.ReplaceAll(bufn, []byte(color))
		}
		if _, err := w.Write(bufn); err != nil {
			log.Fatalf("bufn Write() with error %v", err)
		}
	}
	// never reached
	panic(true)
	return nil
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

	wg.Add(2)
	go func() {
		errStdout = copyAndCapture(os.Stdout, stdoutIn, config)
		wg.Done()
	}()
	go func() {
		errStderr = copyAndCapture(os.Stderr, stderrIn, config)
		wg.Done()
	}()

	go func() {
		// User Send ctrlc to the program
		<-ctrlc
		// Use Process.Release() instead of Process.Kill()
		// Release() waits until the cmd exit
		// Kill() Does not wait
		if err := cmd.Process.Release(); err != nil {
			log.Fatalf("failed to kill process: %v", err)
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
		log.Fatal("failed to capture stdout or stderr\n")
	}

	os.Exit(statusCode)
}
