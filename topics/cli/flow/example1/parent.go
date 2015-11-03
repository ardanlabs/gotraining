package main

import (
	"fmt"
	"os"
	"os/exec"
)

func fail(err error) {
	// output errors to stderr
	fmt.Fprintln(os.Stderr, err)
	// exit with non-zero status to indicate command failure
	os.Exit(1)
}

func main() {
	// os.Args[0] is _this_ program's name
	args := os.Args[1:]
	if len(args) < 1 {
		// nothing to do
		return
	}

	cmd := exec.Command(args[0], args[1:]...)

	// These streams are closed if we don't set them
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fail(err)
	}
}
