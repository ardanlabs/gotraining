package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type AddrVal struct {
	addr *string
}

func (a AddrVal) String() string {
	if a.addr == nil {
		return ""
	}

	return *a.addr
}

func (a AddrVal) Set(s string) error {
	if err := validateAddr(s); err != nil {
		return err
	}

	*a.addr = s
	return nil
}

func validateAddr(addr string) error {
	i := strings.Index(addr, ":")
	if i == -1 {
		return fmt.Errorf("no : in %q", addr)
	}

	port, err := strconv.Atoi(addr[i+1:])
	if err != nil {
		return fmt.Errorf("bad port in %q: - %w", addr, err)
	}

	const minPort, maxPort = 0, 65535
	if port < minPort || port > maxPort {
		return fmt.Errorf("port %d out of range [%d:%d]", i, minPort, maxPort)
	}

	return nil
}

func main() {
	addr := ":8080" // default
	flag.Var(AddrVal{&addr}, "addr", "address to listen on")
	flag.Parse()

	fmt.Println("address:", addr)
}
