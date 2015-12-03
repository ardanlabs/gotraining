package cmduser

import (
	"errors"

	"github.com/spf13/cobra"
)

var createLong = `Use create to add a new user to the system. The user email
must be unique for every user.

Example:
  ./shelf user create -n "Bill Kennedy" -e "bill@ardanlabs.com" -p "yefc*7fdf92"
`

// create contains the state for this command.
var create struct {
	name  string
	pass  string
	email string
}

// addCreate handles the creation of users.
func addCreate() {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Add a new user to the system.",
		Long:  createLong,
		Run:   runCreate,
	}

	cmd.Flags().StringVarP(&create.name, "name", "n", "", "Full name for the user.")
	cmd.Flags().StringVarP(&create.email, "email", "e", "", "Email for the user.")
	cmd.Flags().StringVarP(&create.pass, "pass", "p", "", "Password for the user.")

	userCmd.AddCommand(cmd)
}

// runCreate is the code that implements the create command.
func runCreate(cmd *cobra.Command, args []string) {
	cmd.Printf("Creating User : Name[%s] Email[%s] Pass[%s]\n", create.name, create.email, create.pass)

	if create.name == "" && create.email == "" && create.pass == "" {
		cmd.Help()
		return
	}

	u := User{
		Status:   1,
		Name:     "Bill",
		Email:    "bill@ardanlabs.com",
		Password: "my passoword",
	}

	if err := createUser(&u); err != nil {
		cmd.Println("Creating User : ", err)
		return
	}

	cmd.Println("Creating User : Created")
}

//==============================================================================

// User represents a sample user model.
type User struct {
	Status   int
	Name     string
	Email    string
	Password string
}

// createUser is a sample function to simulate a user creation.
func createUser(u *User) error {
	if u.Status == 0 {
		return errors.New("Invalid user value")
	}

	return nil
}
