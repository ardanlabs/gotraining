// Package endpointTests implements tests for the buoy endpoints.
package tests

import (
	"testing"
	"time"

	"github.com/ArdanStudios/gotraining/11-http/api/app"
	"github.com/ArdanStudios/gotraining/11-http/api/models"
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

// Test_CreateUser validates a user can be created in the system.
func Test_CreateUser(t *testing.T) {
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

	t.Log("Attempt to create a user.")
	if err := u.Create(c); err != nil {
		t.Fatalf("Unable to create a user.", err)
	}

	t.Log("User value should contain a unique ID")
	if u.ID.Hex() == "" {
		t.Fatalf("User value does not contain a unique ID.")
	}
}
