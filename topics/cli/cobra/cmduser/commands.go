package cmduser

import "github.com/spf13/cobra"

// userCmd represents the parent for all user cli commands.
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "user provides a shelf CLI for managing user records.",
}

// GetCommands returns the user commands.
func GetCommands() *cobra.Command {
	// TODO: Add all the commands here.
	addCreate()

	return userCmd
}
