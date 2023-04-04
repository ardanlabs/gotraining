// You can pipe output from one exec.Command to the other.

package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	grep := exec.Command("grep", "func", "pipe.go")
	out, err := grep.StdoutPipe()
	if err != nil {
		log.Fatalf("erro: %s", err)
	}

	if err := grep.Start(); err != nil {
		log.Fatalf("erro: %s", err)
	}

	wc := exec.Command("wc", "-l")
	wc.Stdin = out
	data, err := wc.CombinedOutput()
	if err != nil {
		log.Fatalf("erro: %s", err)
	}
	fmt.Println(string(data))
}
