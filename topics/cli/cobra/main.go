// This program provides a sample building cli tooling.
package main

import (
	"github.com/ardanlabs/gotraining/topics/cli/cobra/cmduser"

	"github.com/spf13/cobra"
)

var shelf = &cobra.Command{
	Use:   "binary",
	Short: "binary provides what...",
}

func main() {
	// TODO: Add more commands here.
	shelf.AddCommand(cmduser.GetCommands())

	// Execute the program and process the flags.
	shelf.Execute()
}
