// Package endpointTests implements tests for the buoy endpoints.
package tests

import (
	"testing"
	"time"

	"github.com/ArdanStudios/gotraining/11-http/api/models"
	"github.com/ArdanStudios/gotraining/11-http/api/services"
)

// Test_UsersCreateRetrieveRemove validates a user can be created, retrieved and
// then removed from the system.
func Test_UsersCreateRetrieveRemove(t *testing.T) {
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

	if _, err := u.Validate(); err != nil {
		t.Fatal("\tShould be able to validate the user data.", failed)
	}
	t.Log("\tShould be able to validate the user data.", succeed)

	if err := services.Users.Create(c, &u); err != nil {
		t.Fatal("\tShould be able to create a user in the system.", failed)
	}
	t.Log("\tShould be able to create a user in the system.", succeed)

	if u.ID.Hex() == "" {
		t.Fatal("\tShould have an ID for the user.", failed)
	}
	t.Log("\tShould have an ID for the user.", succeed)

	ur, err := services.Users.Retrieve(c, u.ID.Hex())
	if err != nil {
		t.Fatal("\tShould be able to retrieve the user back from the system.", failed)
	}
	t.Log("\tShould be able to retrieve the user back from the system.", succeed)

	if _, err := u.Compare(ur); err != nil {
		t.Fatal("\tShould find both the original and retrieved value are identical.", failed)
	}
	t.Log("\tShould find both the original and retrieved value are identical.", succeed)

	if ur == nil || u.ID.Hex() != ur.ID.Hex() {
		t.Fatal("\tShould have a match between the created user and the one retrieved.", failed)
	}
	t.Log("\tShould have a match between the created user and the one retrieved.", succeed)

	if err := services.Users.Delete(c, u.ID); err != nil {
		t.Fatal("\tShould be able to remove the user from the system.", failed)
	}
	t.Log("\tShould be able to remove the user from the system", succeed)

	if _, err := services.Users.Retrieve(c, u.ID.Hex()); err == nil {
		t.Fatal("\tShould NOT be able to retrieve the user back from the system.", failed)
	}
	t.Log("\tShould NOT be able to retrieve the user back from the system.", succeed)
}
