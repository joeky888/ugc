package tool

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"sync"
	"syscall"
)

func copyAndCapture(w io.Writer, r io.Reader, config []Conf) (error) {
	// var out []byte
	buf := make([]byte, 20480)
	for {
		n, err := r.Read(buf[:])
		if n > 0 {
			d := buf[:n]
			// out = append(out, d...)

			for _, conf := range config {
				d = conf.Regex.ReplaceAllFunc(d, func(m []byte) []byte {
					return []byte(fmt.Sprintf(conf.Color, conf.Regex.Find(m)))
				})
			}
			// fmt.Println(outStr)
			os.Stdout.Write(d)
		}
		if err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				// err = nil
				return nil
			}
			return err
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

	stdoutIn, outpipeErr := cmd.StdoutPipe()
	if outpipeErr != nil {
		log.Fatalf("outpipeErr failed with %v\n", outpipeErr)
	}
	stderrIn, errpipeErr := cmd.StderrPipe()
	if errpipeErr != nil {
		log.Fatalf("errpipeErr failed with %v\n", errpipeErr)
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

	if err := cmd.Start(); err != nil {
		log.Fatalf("cmd.Start() failed with %v\n", err)
	}

	wg.Wait()
	if err := cmd.Wait(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				os.Exit(status.ExitStatus())
			}
		} else {
			log.Fatalf("cmd.Wait: %v", err)
		}
		log.Fatalf("cmd.Run() failed with %v\n", err)
	}
	if errStdout != nil || errStderr != nil {
		log.Fatal("failed to capture stdout or stderr\n")
	}
}
