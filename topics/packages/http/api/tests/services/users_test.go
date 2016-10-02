// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package servicestests

import (
	"testing"
	"time"

	"github.com/ardanlabs/gotraining/topics/packages/http/api/app"
	"github.com/ardanlabs/gotraining/topics/packages/http/api/models"
	"github.com/ardanlabs/gotraining/topics/packages/http/api/services"
)

// Succeed is the Unicode codepoint for a check mark.
const Succeed = "\u2713"

// Failed is the Unicode codepoint for an X mark.
const Failed = "\u2717"

// TestUsers validates a user can be created, retrieved and
// then removed from the system.
func TestUsers(t *testing.T) {
	c := &app.Context{
		Ctx:       make(map[string]interface{}),
		SessionID: "TESTING",
	}

	ses := app.GetSession()
	c.Ctx["DB"] = ses
	defer ses.Close()

	now := time.Now()

	u := models.User{
		UserType:     1,
		FirstName:    "Bill",
		LastName:     "Kennedy",
		Email:        "bill@ardanstugios.com",
		Company:      "Ardan Labs",
		DateModified: &now,
		DateCreated:  &now,
		Addresses: []models.UserAddress{
			{
				Type:         1,
				LineOne:      "12973 SW 112th ST",
				LineTwo:      "Suite 153",
				City:         "Miami",
				State:        "FL",
				Zipcode:      "33172",
				Phone:        "305-527-3353",
				DateModified: &now,
				DateCreated:  &now,
			},
		},
	}

	t.Log("Given the need to add a new user, retrieve and remove that user from the system.")
	{
		if _, err := u.Validate(); err != nil {
			t.Fatal("\tShould be able to validate the user data.", Failed)
		}
		t.Log("\tShould be able to validate the user data.", Succeed)

		if _, err := services.Users.Create(c, &u); err != nil {
			t.Fatal("\tShould be able to create a user in the system.", Failed)
		}
		t.Log("\tShould be able to create a user in the system.", Succeed)

		if u.UserID == "" {
			t.Fatal("\tShould have an UserID for the user.", Failed)
		}
		t.Log("\tShould have an UserID for the user.", Succeed)

		ur, err := services.Users.Retrieve(c, u.UserID)
		if err != nil {
			t.Fatal("\tShould be able to retrieve the user back from the system.", Failed)
		}
		t.Log("\tShould be able to retrieve the user back from the system.", Succeed)

		if _, err := u.Compare(ur); err != nil {
			t.Fatal("\tShould find both the original and retrieved value are identical.", Failed)
		}
		t.Log("\tShould find both the original and retrieved value are identical.", Succeed)

		if ur == nil || u.UserID != ur.UserID {
			t.Fatal("\tShould have a match between the created user and the one retrieved.", Failed)
		}
		t.Log("\tShould have a match between the created user and the one retrieved.", Succeed)

		if err := services.Users.Delete(c, u.UserID); err != nil {
			t.Fatal("\tShould be able to remove the user from the system.", Failed)
		}
		t.Log("\tShould be able to remove the user from the system", Succeed)

		if _, err := services.Users.Retrieve(c, u.UserID); err == nil {
			t.Fatal("\tShould NOT be able to retrieve the user back from the system.", Failed)
		}
		t.Log("\tShould NOT be able to retrieve the user back from the system.", Succeed)
	}
}
