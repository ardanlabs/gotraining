package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	c := color{os.Stdout}
	io.Copy(&c, os.Stdin)
}

func (c *color) Write(buf []byte) (int, error) {
	for _, b := range buf {
		switch {
		case b >= 'A' && b <= 'Z':
			fmt.Fprintf(c.w, "\033[32m%c\033[39m", b)
		case b >= 'a' && b <= 'z':
			fmt.Fprintf(c.w, "\033[31m%c\033[39m", b)
		default:
			fmt.Fprintf(c.w, "%c", b)
		}
	}

	return len(buf), nil
}

type color struct {
	w io.Writer
}
