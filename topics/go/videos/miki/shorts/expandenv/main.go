// Use os.ExpandEnv to expand environment variables in a string.
package main

import (
	"fmt"
	"os"
)

func main() {
	cfg := `
[httpd]
user = ${USER}
config_dir = ${HOME}/.config/httpd
`
	fmt.Println(os.ExpandEnv(cfg))
}
