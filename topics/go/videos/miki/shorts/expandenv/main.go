// Use os.ExpandEnv to expand environment variables in a string.
package main

import (
	"fmt"
	"os"
)

func main() {
	cfg := `
[httpd]
use = ${USER}
root = ${HOME}/.config/httpd
`
	fmt.Println(os.ExpandEnv(cfg))
}
