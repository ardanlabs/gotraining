// Package endpointtests implements users tests for the API layer.
package endpointtests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/ArdanStudios/gotraining/11-http/api/app"
	"github.com/ArdanStudios/gotraining/11-http/api/models"
	"github.com/ArdanStudios/gotraining/11-http/api/routes"
	"github.com/ArdanStudios/gotraining/11-http/api/tests"
	"gopkg.in/mgo.v2/bson"
)

var u = models.User{
	UserType:  1,
	FirstName: "Bill",
	LastName:  "Kennedy",
	Email:     "bill@ardanstugios.com",
	Company:   "Ardan Labs",
	Addresses: []models.UserAddress{
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

// Test_Users is the entry point for the users tests.
func Test_Users(t *testing.T) {
	c := &app.Context{
		Session:   app.GetSession(),
		SessionID: "TESTING",
	}
	defer c.Session.Close()

	usersList204(t, c)
	usersCreate200(t, c)
	usersCreate409(t, c)
	us := usersList200(t, c)
	usersRetrieve200(t, c, us[0].UserID)
	usersRetrieve404(t, c, bson.NewObjectId().Hex())
	usersRetrieve409(t, c, "123")
	usersDelete200(t, c, us[0].UserID)
}

// usersList204 validates an empty users list can be retrieved with the endpoint.
func usersList204(t *testing.T, c *app.Context) {
	r := tests.NewRequest("GET", "/v1/users", nil)
	w := httptest.NewRecorder()
	routes.TM.ServeHTTP(w, r)

	t.Log("Given the need to validate an empty list of users with the users endpoint.")
	{
		if w.Code != 204 {
			t.Fatalf("\tShould received a status code of 204 for the response. Received[%d] %s", w.Code, tests.Failed)
		}
		t.Log("\tShould received a status code of 204 for the response.", tests.Succeed)
	}
}

// usersCreate200 validates a user can be created with the endpoint.
func usersCreate200(t *testing.T, c *app.Context) {
	var response struct {
		UserID string `json:"user_id"`
	}

	body, _ := json.Marshal(&u)
	r := tests.NewRequest("POST", "/v1/users", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	routes.TM.ServeHTTP(w, r)

	t.Log("Given the need to add a new user with the users endpoint.")
	{
		if w.Code != 200 {
			t.Fatalf("\tShould received a status code of 200 for the response. Received[%d] %s", w.Code, tests.Failed)
		}
		t.Log("\tShould received a status code of 200 for the response.", tests.Succeed)

		if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
			t.Fatal("\tShould be able to unmarshal the response.", tests.Failed)
		}
		t.Log("\tShould be able to unmarshal the response.", tests.Succeed)

		if response.UserID == "" {
			t.Fatal("\tShould have a user id in the response.", tests.Failed)
		}
		t.Log("\tShould have a user id in the response.", tests.Succeed)
	}
}

// usersCreate409 validates a user can't be created with the endpoint
// unless a valid user document is submitted.
func usersCreate409(t *testing.T, c *app.Context) {
	u := models.User{
		UserType: 1,
		LastName: "Kennedy",
		Email:    "bill@ardanstugios.com",
		Company:  "Ardan Labs",
	}

	var v []app.Invalid

	body, _ := json.Marshal(&u)
	r := tests.NewRequest("POST", "/v1/users", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	routes.TM.ServeHTTP(w, r)

	t.Log("Given the need to validate a new user can't be created with an invalid document.")
	{
		if w.Code != 409 {
			t.Fatalf("\tShould received a status code of 409 for the response. Received[%d] %s", w.Code, tests.Failed)
		}
		t.Log("\tShould received a status code of 409 for the response.", tests.Succeed)

		if err := json.NewDecoder(w.Body).Decode(&v); err != nil {
			t.Fatal("\tShould be able to unmarshal the response.", tests.Failed)
		}
		t.Log("\tShould be able to unmarshal the response.", tests.Succeed)

		if len(v) == 0 {
			t.Fatal("\tShould have validation errors in the response.", tests.Failed)
		}
		t.Log("\tShould have validation errors in the response.", tests.Succeed)

		if v[0].Fld != "FirstName" {
			t.Fatalf("\tShould have a FirstName validation error in the response. Received[%s] %s", v[0].Fld, tests.Failed)
		}
		t.Log("\tShould have a FirstName validation error in the response.", tests.Succeed)

		if v[1].Fld != "Addresses" {
			t.Fatalf("\tShould have an Addresses validation error in the response. Received[%s] %s", v[0].Fld, tests.Failed)
		}
		t.Log("\tShould have an Addresses validation error in the response.", tests.Succeed)
	}
}

// usersList200 validates a users list can be retrieved with the endpoint.
func usersList200(t *testing.T, c *app.Context) []models.User {
	var us []models.User

	r := tests.NewRequest("GET", "/v1/users", nil)
	w := httptest.NewRecorder()
	routes.TM.ServeHTTP(w, r)

	t.Log("Given the need to retrieve a list of users with the users endpoint.")
	{
		if w.Code != 200 {
			t.Fatalf("\tShould received a status code of 200 for the response. Received[%d] %s", w.Code, tests.Failed)
		}
		t.Log("\tShould received a status code of 200 for the response.", tests.Succeed)

		if err := json.NewDecoder(w.Body).Decode(&us); err != nil {
			t.Fatal("\tShould be able to unmarshal the response.", tests.Failed)
		}
		t.Log("\tShould be able to unmarshal the response.", tests.Succeed)

		if len(us) == 0 {
			t.Fatal("\tShould have users in the response.", tests.Failed)
		}
		t.Log("\tShould have a users in the response.", tests.Succeed)

		var failed bool
		marks := make([]string, len(us))
		for i, u := range us {
			if u.DateCreated == nil || u.DateModified == nil {
				marks[i] = tests.Failed
				failed = true
			} else {
				marks[i] = tests.Succeed
			}
		}

		if failed {
			t.Fatalf("\tShould have dates in all the user documents. %+v", marks)
		}
		t.Logf("\tShould have dates in all the user documents. %+v", marks)
	}

	return us
}

// usersList200 validates a users list can be retrieved with the endpoint.
func usersRetrieve200(t *testing.T, c *app.Context, id string) {
	r := tests.NewRequest("GET", "/v1/users/"+id, nil)
	w := httptest.NewRecorder()
	routes.TM.ServeHTTP(w, r)

	var ur models.User

	t.Log("Given the need to retrieve an individual user with the users endpoint.")
	{
		if w.Code != 200 {
			t.Fatalf("\tShould received a status code of 200 for the response. Received[%d] %s", w.Code, tests.Failed)
		}
		t.Log("\tShould received a status code of 200 for the response.", tests.Succeed)

		if err := json.NewDecoder(w.Body).Decode(&ur); err != nil {
			t.Fatal("\tShould be able to unmarshal the response.", tests.Failed)
		}
		t.Log("\tShould be able to unmarshal the response.", tests.Succeed)

		if inv, err := ur.Compare(&u); err != nil {
			fmt.Println(inv)
			t.Fatal("\tShould have the document specified by id.", tests.Failed)
		}
		t.Log("\tShould have the document specified by id", tests.Succeed)
	}
}

// usersRetrieve404 validates a user request for a user that does not exist with the endpoint.
func usersRetrieve404(t *testing.T, c *app.Context, id string) {
	r := tests.NewRequest("GET", "/v1/users/"+id, nil)
	w := httptest.NewRecorder()
	routes.TM.ServeHTTP(w, r)

	t.Log("Given the situation of retrieving an individual user that does not exist with the users endpoint.")
	{
		if w.Code != 404 {
			t.Fatalf("\tShould received a status code of 404 for the response. Received[%d] %s", w.Code, tests.Failed)
		}
		t.Log("\tShould received a status code of 404 for the response.", tests.Succeed)
	}
}

// usersRetrieve409 validates a user request with an invalid id with the endpoint.
func usersRetrieve409(t *testing.T, c *app.Context, id string) {
	r := tests.NewRequest("GET", "/v1/users/"+id, nil)
	w := httptest.NewRecorder()
	routes.TM.ServeHTTP(w, r)

	t.Log("Given the situation of retrieving an individual user with an invalid id with the users endpoint.")
	{
		if w.Code != 409 {
			t.Fatalf("\tShould received a status code of 409 for the response. Received[%d] %s", w.Code, tests.Failed)
		}
		t.Log("\tShould received a status code of 409 for the response.", tests.Succeed)
	}
}

// usersDelete200 validates a user can be deleted with the endpoint.
func usersDelete200(t *testing.T, c *app.Context, id string) {
	var response struct {
		Message string `json:"message"`
	}

	r := tests.NewRequest("DELETE", "/v1/users/"+id, nil)
	w := httptest.NewRecorder()
	routes.TM.ServeHTTP(w, r)

	t.Log("Given the need to delete a new user with the users endpoint.")
	{
		if w.Code != 200 {
			t.Fatalf("\tShould received a status code of 200 for the response. Received[%d] %s", w.Code, tests.Failed)
		}
		t.Log("\tShould received a status code of 200 for the response.", tests.Succeed)

		if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
			t.Fatal("\tShould be able to unmarshal the response.", tests.Failed)
		}
		t.Log("\tShould be able to unmarshal the response.", tests.Succeed)

		if response.Message != fmt.Sprintf("User with ID %s has been removed.", id) {
			t.Fatal("\tShould have an expected message in the response.", tests.Failed)
		}
		t.Log("\tShould have an expected message in the response.", tests.Succeed)
	}
}
