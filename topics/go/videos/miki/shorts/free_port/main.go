package main

import (
	"fmt"
	"net"
)

// FreePort return random free port to use.
func FreePort() (int, error) {
	lis, err := net.Listen("tcp", "")
	if err != nil {
		return 0, err
	}
	lis.Close()

	port := lis.Addr().(*net.TCPAddr).Port
	return port, nil
}

func main() {
	fmt.Println(FreePort())
}
