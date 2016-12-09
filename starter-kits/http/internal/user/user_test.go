// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package user_test

import (
	"context"
	"testing"

	"github.com/ardanlabs/gotraining/starter-kits/http/internal/sys/db"
	"github.com/ardanlabs/gotraining/starter-kits/http/internal/user"
)

const (

	// Succeed is the Unicode codepoint for a check mark.
	Succeed = "\u2713"

	// Failed is the Unicode codepoint for an X mark.
	Failed = "\u2717"
)

// TestUsers validates a user can be created, retrieved and
// then removed from the system.
func TestUsers(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dbSes := db.Init()
	defer dbSes.Close()

	traceID := "traceid"

	u := user.CreateUser{
		UserType:  1,
		FirstName: "Bill",
		LastName:  "Kennedy",
		Email:     "bill@ardanstugios.com",
		Company:   "Ardan Labs",
		Addresses: []user.CreateAddress{
			{
				Type:    1,
				LineOne: "12973 SW 112th ST",
				LineTwo: "Suite 153",
				City:    "Miami",
				State:   "FL",
				Zipcode: "33172",
				Phone:   "305-527-3353",
			},
		},
	}

	t.Log("Given the need to add a new user, retrieve and remove that user from the system.")
	{
		t.Log("\tTest 0:\tWhen using a valid CreateUser value")
		{
			cu, err := user.Create(ctx, traceID, dbSes, &u)
			if err != nil {
				t.Fatalf("\t%s\tShould be able to create a user in the system : %v", Failed, err)
			}
			t.Logf("\t%s\tShould be able to create a user in the system.", Succeed)

			ru, err := user.Retrieve(ctx, traceID, dbSes, cu.UserID)
			if err != nil {
				t.Fatalf("\t%s\tShould be able to retrieve the user back from the system : %v", Failed, err)
			}
			t.Logf("\t%s\tShould be able to retrieve the user back from the system.", Succeed)

			if ru == nil || cu.UserID != ru.UserID {
				t.Fatalf("\t%s\tShould have a match between the created user and the one retrieved : %v", Failed, err)
			}
			t.Logf("\t%s\tShould have a match between the created user and the one retrieved.", Succeed)

			if err := user.Delete(ctx, traceID, dbSes, ru.UserID); err != nil {
				t.Fatalf("\t%s\tShould be able to remove the user from the system : %v", Failed, err)
			}
			t.Logf("\t%s\tShould be able to remove the user from the system.", Succeed)

			if _, err := user.Retrieve(ctx, traceID, dbSes, ru.UserID); err == nil {
				t.Fatalf("\t%s\tShould NOT be able to retrieve the user back from the system : %v", Failed, err)
			}
			t.Logf("\t%s\tShould NOT be able to retrieve the user back from the system.", Succeed)
		}
	}
}
