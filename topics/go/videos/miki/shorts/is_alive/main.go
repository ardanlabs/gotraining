// Send signal 0 to check if a process is alive.

package main

import (
	"fmt"
	"os"
	"syscall"
)

func isAlive(pid int) bool {
	return syscall.Kill(pid, 0) == nil
}

func main() {
	fmt.Println(isAlive(os.Getpid())) // true
	fmt.Println(isAlive(666))         // false
}
