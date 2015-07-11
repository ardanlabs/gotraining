// Package endpointtests implements users tests for the API layer.
package endpointtests

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/ArdanStudios/gotraining/12-http/api/app"
	"github.com/ArdanStudios/gotraining/12-http/api/models"
	"github.com/ArdanStudios/gotraining/12-http/api/routes"
	"github.com/ArdanStudios/gotraining/12-http/api/tests"
	"gopkg.in/mgo.v2/bson"
)

var u = models.User{
	UserType:  1,
	FirstName: "Bill",
	LastName:  "Kennedy",
	Email:     "bill@ardanstudios.com",
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

// TestUsers is the entry point for the users tests.
func TestUsers(t *testing.T) {
	c := &app.Context{
		Session:   app.GetSession(),
		SessionID: "TESTING",
	}
	defer c.Session.Close()

	a := routes.API().(*app.App)

	//usersList404(t, a, c)
	usersCreate200(t, a, c)
	usersCreate400(t, a, c)
	us := usersList200(t, a, c)
	usersRetrieve200(t, a, c, us[0].UserID)
	usersRetrieve404(t, a, c, bson.NewObjectId().Hex())
	usersRetrieve400(t, a, c, "123")
	usersUpdate200(t, a, c)
	usersRetrieve200(t, a, c, us[0].UserID)
	usersDelete200(t, a, c, us[0].UserID)
	usersDelete404(t, a, c, us[0].UserID)
}

// usersList404 validates an empty users list can be retrieved with the endpoint.
func usersList404(t *testing.T, a *app.App, c *app.Context) {
	r := tests.NewRequest("GET", "/v1/users", nil)
	w := httptest.NewRecorder()
	a.ServeHTTP(w, r)

	t.Log("Given the need to validate an empty list of users with the users endpoint.")
	{
		if w.Code != 404 {
			t.Fatalf("\tShould received a status code of 404 for the response. Received[%d] %s", w.Code, tests.Failed)
		}
		t.Log("\tShould received a status code of 404 for the response.", tests.Succeed)
	}
}

// usersCreate200 validates a user can be created with the endpoint.
func usersCreate200(t *testing.T, a *app.App, c *app.Context) {
	body, _ := json.Marshal(&u)
	r := tests.NewRequest("POST", "/v1/users", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	a.ServeHTTP(w, r)

	t.Log("Given the need to add a new user with the users endpoint.")
	{
		if w.Code != 200 {
			t.Fatalf("\tShould received a status code of 200 for the response. Received[%d] %s", w.Code, tests.Failed)
		}
		t.Log("\tShould received a status code of 200 for the response.", tests.Succeed)

		var resp models.User
		if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
			t.Fatal("\tShould be able to unmarshal the response.", tests.Failed)
		}
		t.Log("\tShould be able to unmarshal the response.", tests.Succeed)

		if resp.UserID == "" {
			t.Fatal("\tShould have a user id in the response.", tests.Failed)
		}
		t.Log("\tShould have a user id in the response.", tests.Succeed)

		// Save for future calls.
		u.UserID = resp.UserID
	}
}

// usersCreate400 validates a user can't be created with the endpoint
// unless a valid user document is submitted.
func usersCreate400(t *testing.T, a *app.App, c *app.Context) {
	u := models.User{
		UserType: 1,
		LastName: "Kennedy",
		Email:    "bill@ardanstugios.com",
		Company:  "Ardan Labs",
	}

	body, _ := json.Marshal(&u)
	r := tests.NewRequest("POST", "/v1/users", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	a.ServeHTTP(w, r)

	t.Log("Given the need to validate a new user can't be created with an invalid document.")
	{
		if w.Code != 400 {
			t.Fatalf("\tShould received a status code of 400 for the response. Received[%d] %s", w.Code, tests.Failed)
		}
		t.Log("\tShould received a status code of 400 for the response.", tests.Succeed)

		v := struct {
			Error  string `json:"error"`
			Fields []struct {
				Fld string `json:"field_name"`
				Err string `json:"error"`
			} `json:"fields,omitempty"`
		}{}

		if err := json.NewDecoder(w.Body).Decode(&v); err != nil {
			t.Fatal("\tShould be able to unmarshal the response.", tests.Failed)
		}
		t.Log("\tShould be able to unmarshal the response.", tests.Succeed)

		if len(v.Fields) == 0 {
			t.Fatal("\tShould have validation errors in the response.", tests.Failed)
		}
		t.Log("\tShould have validation errors in the response.", tests.Succeed)

		if v.Fields[0].Fld != "FirstName" {
			t.Fatalf("\tShould have a FirstName validation error in the response. Received[%s] %s", v.Fields[0].Fld, tests.Failed)
		}
		t.Log("\tShould have a FirstName validation error in the response.", tests.Succeed)

		if v.Fields[1].Fld != "Addresses" {
			t.Fatalf("\tShould have an Addresses validation error in the response. Received[%s] %s", v.Fields[0].Fld, tests.Failed)
		}
		t.Log("\tShould have an Addresses validation error in the response.", tests.Succeed)
	}
}

// usersList200 validates a users list can be retrieved with the endpoint.
func usersList200(t *testing.T, a *app.App, c *app.Context) []models.User {
	r := tests.NewRequest("GET", "/v1/users", nil)
	w := httptest.NewRecorder()
	a.ServeHTTP(w, r)

	t.Log("Given the need to retrieve a list of users with the users endpoint.")
	{
		if w.Code != 200 {
			t.Fatalf("\tShould received a status code of 200 for the response. Received[%d] %s", w.Code, tests.Failed)
		}
		t.Log("\tShould received a status code of 200 for the response.", tests.Succeed)

		var us []models.User
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

		return us
	}
}

// usersList200 validates a users list can be retrieved with the endpoint.
func usersRetrieve200(t *testing.T, a *app.App, c *app.Context, id string) {
	r := tests.NewRequest("GET", "/v1/users/"+id, nil)
	w := httptest.NewRecorder()
	a.ServeHTTP(w, r)

	t.Log("Given the need to retrieve an individual user with the users endpoint.")
	{
		if w.Code != 200 {
			t.Fatalf("\tShould received a status code of 200 for the response. Received[%d] %s", w.Code, tests.Failed)
		}
		t.Log("\tShould received a status code of 200 for the response.", tests.Succeed)

		var ur models.User
		if err := json.NewDecoder(w.Body).Decode(&ur); err != nil {
			t.Fatal("\tShould be able to unmarshal the response.", tests.Failed)
		}
		t.Log("\tShould be able to unmarshal the response.", tests.Succeed)

		if ur.UserID != id {
			t.Fatal("\tShould have the document specified by id.", tests.Failed)
		}
		t.Log("\tShould have the document specified by id", tests.Succeed)
	}
}

// usersRetrieve404 validates a user request for a user that does not exist with the endpoint.
func usersRetrieve404(t *testing.T, a *app.App, c *app.Context, id string) {
	r := tests.NewRequest("GET", "/v1/users/"+id, nil)
	w := httptest.NewRecorder()
	a.ServeHTTP(w, r)

	t.Log("Given the situation of retrieving an individual user that does not exist with the users endpoint.")
	{
		if w.Code != 404 {
			t.Fatalf("\tShould received a status code of 404 for the response. Received[%d] %s", w.Code, tests.Failed)
		}
		t.Log("\tShould received a status code of 404 for the response.", tests.Succeed)
	}
}

// usersRetrieve400 validates a user request with an invalid id with the endpoint.
func usersRetrieve400(t *testing.T, a *app.App, c *app.Context, id string) {
	r := tests.NewRequest("GET", "/v1/users/"+id, nil)
	w := httptest.NewRecorder()
	a.ServeHTTP(w, r)

	t.Log("Given the situation of retrieving an individual user with an invalid id with the users endpoint.")
	{
		if w.Code != 400 {
			t.Fatalf("\tShould received a status code of 400 for the response. Received[%d] %s", w.Code, tests.Failed)
		}
		t.Log("\tShould received a status code of 400 for the response.", tests.Succeed)
	}
}

// usersUpdate200 validates a user can be updated with the endpoint.
func usersUpdate200(t *testing.T, a *app.App, c *app.Context) {
	u.FirstName = "Lisa"

	body, _ := json.Marshal(&u)
	r := tests.NewRequest("PUT", "/v1/users/"+u.UserID, bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	a.ServeHTTP(w, r)

	t.Log("Given the need to validate a user can be updated with the users endpoint.")
	{
		if w.Code != 200 {
			t.Fatalf("\tShould received a status code of 200 for the response. Received[%d] %s", w.Code, tests.Failed)
		}
		t.Log("\tShould received a status code of 200 for the response.", tests.Succeed)

		var resp models.User
		if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
			t.Fatal("\tShould be able to unmarshal the response.", tests.Failed)
		}
		t.Log("\tShould be able to unmarshal the response.", tests.Succeed)

		if resp.UserID != u.UserID {
			t.Fatal("\tShould have an a user value with the same id.", tests.Failed)
		}
		t.Log("\tShould have an a user value with the same id.", tests.Succeed)
	}
}

// usersDelete200 validates a user can be deleted with the endpoint.
func usersDelete200(t *testing.T, a *app.App, c *app.Context, id string) {
	r := tests.NewRequest("DELETE", "/v1/users/"+id, nil)
	w := httptest.NewRecorder()
	a.ServeHTTP(w, r)

	t.Log("Given the need to delete a new user with the users endpoint.")
	{
		if w.Code != 200 {
			t.Fatalf("\tShould received a status code of 200 for the response. Received[%d] %s", w.Code, tests.Failed)
		}
		t.Log("\tShould received a status code of 200 for the response.", tests.Succeed)

		var resp models.User
		if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
			t.Fatal("\tShould be able to unmarshal the response.", tests.Failed)
		}
		t.Log("\tShould be able to unmarshal the response.", tests.Succeed)

		if resp.UserID != id {
			t.Fatal("\tShould have an a user value with the same id.", tests.Failed)
		}
		t.Log("\tShould have an a user value with the same id.", tests.Succeed)
	}
}

// usersDelete404 validates a user that has been deleted is deleted.
func usersDelete404(t *testing.T, a *app.App, c *app.Context, id string) {
	r := tests.NewRequest("DELETE", "/v1/users/"+id, nil)
	w := httptest.NewRecorder()
	a.ServeHTTP(w, r)

	t.Log("Given the need to verify a deleted user is deleted.")
	{
		if w.Code != 404 {
			t.Fatalf("\tShould received a status code of 404 for the response. Received[%d] %s", w.Code, tests.Failed)
		}
		t.Log("\tShould received a status code of 404 for the response.", tests.Succeed)
	}
}
