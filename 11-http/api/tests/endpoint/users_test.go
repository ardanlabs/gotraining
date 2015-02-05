// Package endpointtests implements users tests for the API layer.
package endpointtests

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ArdanStudios/gotraining/11-http/api/app"
	"github.com/ArdanStudios/gotraining/11-http/api/models"
	"github.com/ArdanStudios/gotraining/11-http/api/routes"
	"github.com/ArdanStudios/gotraining/11-http/api/tests"
)

// Test_Users is the entry point for the users tests.
func Test_Users(t *testing.T) {
	c := &app.Context{
		Session:   app.GetSession(),
		SessionID: "TESTING",
	}
	defer c.Session.Close()

	usersCreate200(t, c)
}

// usersCreate200 validates a user can be created with the Create endpoint.
func usersCreate200(t *testing.T, c *app.Context) {
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

	var response struct {
		ID string
	}

	body, _ := json.Marshal(&u)
	r := tests.NewRequest("POST", "/v1/users", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	routes.TM.ServeHTTP(w, r)

	t.Log("Given the need to add a new user from the users endpoint.")
	{
		if w.Code != 200 {
			t.Fatalf("\tShould received a status code of 200 for the response. Received[%d] %s", w.Code, tests.Failed)
		}
		t.Log("\tShould received a status code of 200 for the response.", tests.Succeed)

		if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
			t.Fatal("\tShould be able to unmarshal the response.", tests.Failed)
		}
		t.Log("\tShould be able to unmarshal the response.", tests.Succeed)

		if response.ID == "" {
			t.Fatal("\tShould have a user id in the response.", tests.Failed)
		}
		t.Log("\tShould have a user id in the response.", tests.Succeed)
	}
}
