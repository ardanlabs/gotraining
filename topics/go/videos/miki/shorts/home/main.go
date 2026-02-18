package main

import (
	"fmt"
	"os"
	"os/user"
	"path"
)

func main() {
	u, err := user.Current()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: current user - %v\n", err)
		os.Exit(1)
	}

	configFile := path.Join(u.HomeDir, ".config", "app.yml")
	fmt.Println("config file:", configFile)
}
