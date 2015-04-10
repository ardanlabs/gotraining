package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	c1 := exec.Command("df", "-h")
	c2 := exec.Command("sort")

	r, err := c1.StdoutPipe()
	if err != nil {
		fmt.Println("failed to create pipe:", err)
		return
	}
	defer r.Close()

	c1.Stdin = os.Stdin
	c1.Stderr = os.Stderr

	// the stdout of c1 feeds the stdin of c2
	c2.Stdin = r
	c2.Stdout = os.Stdout
	c2.Stderr = os.Stderr

	err = c1.Start()
	if err != nil {
		fmt.Println("process #1 failed to start:", err)
	}

	// we Wait on c1 to clean up its process entry when done
	// (otherwise, it'll remain as a zombie process until this process exits)
	defer c1.Wait()

	err = c2.Run()
	if err != nil {
		fmt.Println("process #2 failed:", err)
	}
}
