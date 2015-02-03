// Package endpointTests implements tests for the buoy endpoints.
package tests

import (
	"testing"
	"time"

	"github.com/ArdanStudios/gotraining/11-http/api/app"
	"github.com/ArdanStudios/gotraining/11-http/api/models"
	"github.com/ArdanStudios/gotraining/11-http/api/services"
)

var c *app.Context

// TestMain is the entry point for the test. Used to create a context
// before the tests are run and can then perform cleanup.
func TestMain(m *testing.M) {
	c = &app.Context{
		Session:   app.GetSession(),
		SessionID: "TESTING",
	}

	m.Run()

	c.Session.Close()
}

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
		DateModified: now,
		DateCreated:  now,
		Addresses: []models.UserAddress{
			{
				Type:         1,
				LineOne:      "12973 SW 112th ST",
				LineTwo:      "Suite 153",
				City:         "Miami",
				State:        "FL",
				Zipcode:      "33172",
				Phone:        "305-527-3353",
				DateModified: now,
				DateCreated:  now,
			},
		},
	}

	t.Log("Given the need to add a new user, retrieve and remove that user from the system.")

	if _, err := u.Validate(); err != nil {
		t.Fatal("\tShould be able to validate the user data.", failed)
	}
	t.Log("\tShould be able to validate the user data.", succeed)

	if err := services.UsersCreate(c, &u); err != nil {
		t.Fatal("\tShould be able to create a user in the system.", failed)
	}
	t.Log("\tShould be able to create a user in the system..", succeed)

	if u.ID.Hex() == "" {
		t.Fatal("\tShould have an ID for the user.", failed)
	}
	t.Log("\tShould have an ID for the user.", succeed)

	ur, err := services.UsersRetrieve(c, u.ID)
	if err != nil {
		t.Fatal("\tShould be able to retrieve the user back from the system.", failed)
	}
	t.Log("\tShould be able to retrieve the user back from the system.", succeed)

	if ur == nil || u.ID.Hex() != ur.ID.Hex() {
		t.Fatal("\tShould have a match between the created user and the one retrieved.", failed)
	}
	t.Log("\tShould have a match between the created user and the one retrieved.", succeed)

	if err := services.UsersDelete(c, u.ID); err != nil {
		t.Fatal("\tShould be able to remove the user from the system.", failed)
	}
	t.Log("\tShould be able to remove the user from the system", succeed)

	if _, err := services.UsersRetrieve(c, u.ID); err == nil {
		t.Fatal("\tShould NOT be able to retrieve the user back from the system.", failed)
	}
	t.Log("\tShould NOT be able to retrieve the user back from the system.", succeed)
}
