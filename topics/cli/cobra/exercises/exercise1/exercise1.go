package cmduser

import (
	"encoding/json"

	"github.com/spf13/cobra"
)

var getLong = `Use get to retrieve a user from the system.

Example:
  ./shelf user get -e "bill@ardanlabs.com"
`

// get contains the state for this command.
var get struct {
	email string
}

// addCreate handles the creation of users.
func addCreate() {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get retrieve a user from the system.",
		Long:  getLong,
		Run:   runGet,
	}

	cmd.Flags().StringVarP(&get.email, "email", "e", "", "Email for the user.")

	// THIS WILL BE AVAILABLE WHEN IN THE SAME PACKAGE.
	//userCmd.AddCommand(cmd)
}

// runGet is the code that implements the create command.
func runGet(cmd *cobra.Command, args []string) {
	cmd.Printf("Getting User : Email[%s]\n", get.email)

	if get.email == "" {
		cmd.Help()
		return
	}

	u, err := getUser(get.email)
	if err != nil {
		cmd.Println("Getting User : ", err)
		return
	}

	data, err := json.MarshalIndent(&u, "", "    ")
	if err != nil {
		cmd.Println("Getting User : ", err)
		return
	}

	cmd.Printf("\n%s\n\n", string(data))
}

//==============================================================================

// User represents a sample user model.
// NEED THIS FOR THE EXERCISE TO BUILD.
type User struct {
	Status   int
	Name     string
	Email    string
	Password string
}

// createUser is a sample function to simulate a user creation.
func getUser(email string) (*User, error) {
	u := User{
		Status:   1,
		Name:     "Bill",
		Email:    "bill@ardanlabs.com",
		Password: "my passoword",
	}

	return &u, nil
}
