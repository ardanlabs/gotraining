// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package endpointtests

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/ardanlabs/gotraining/topics/packages/http/api/app"
	"github.com/ardanlabs/gotraining/topics/packages/http/api/models"
	"github.com/ardanlabs/gotraining/topics/packages/http/api/routes"
	"gopkg.in/mgo.v2/bson"
)

// Succeed is the Unicode codepoint for a check mark.
const Succeed = "\u2713"

// Failed is the Unicode codepoint for an X mark.
const Failed = "\u2717"

var u = models.User{
	UserType:  1,
	FirstName: "Bill",
	LastName:  "Kennedy",
	Email:     "bill@ardanlabs.com",
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

// TestUsers is the entry point for the users
func TestUsers(t *testing.T) {
	a := routes.API().(*app.App)

	usersList404(t, a)
	usersCreate200(t, a)
	usersCreate400(t, a)
	us := usersList200(t, a)
	usersRetrieve200(t, a, us[0].UserID)
	usersRetrieve404(t, a, bson.NewObjectId().Hex())
	usersRetrieve400(t, a, "123")
	usersUpdate200(t, a)
	usersRetrieve200(t, a, us[0].UserID)
	usersDelete200(t, a, us[0].UserID)
	usersDelete404(t, a, us[0].UserID)
}

// usersList404 validates an empty users list can be retrieved with the endpoint.
func usersList404(t *testing.T, a *app.App) {
	r := httptest.NewRequest("GET", "/v1/users", nil)
	w := httptest.NewRecorder()
	a.ServeHTTP(w, r)

	t.Log("Given the need to validate an empty list of users with the users endpoint.")
	{
		if w.Code != 404 {
			t.Fatalf("\tShould received a status code of 404 for the response. Received[%d] %s", w.Code, Failed)
		}
		t.Log("\tShould received a status code of 404 for the response.", Succeed)
	}
}

// usersCreate200 validates a user can be created with the endpoint.
func usersCreate200(t *testing.T, a *app.App) {
	body, _ := json.Marshal(&u)
	r := httptest.NewRequest("POST", "/v1/users", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	a.ServeHTTP(w, r)

	t.Log("Given the need to add a new user with the users endpoint.")
	{
		if w.Code != 200 {
			t.Fatalf("\tShould received a status code of 200 for the response. Received[%d] %s", w.Code, Failed)
		}
		t.Log("\tShould received a status code of 200 for the response.", Succeed)

		var resp models.User
		if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
			t.Fatal("\tShould be able to unmarshal the response.", Failed)
		}
		t.Log("\tShould be able to unmarshal the response.", Succeed)

		if resp.UserID == "" {
			t.Fatal("\tShould have a user id in the response.", Failed)
		}
		t.Log("\tShould have a user id in the response.", Succeed)

		// Save for future calls.
		u.UserID = resp.UserID
	}
}

// usersCreate400 validates a user can't be created with the endpoint
// unless a valid user document is submitted.
func usersCreate400(t *testing.T, a *app.App) {
	u := models.User{
		UserType: 1,
		LastName: "Kennedy",
		Email:    "bill@ardanstugios.com",
		Company:  "Ardan Labs",
	}

	body, _ := json.Marshal(&u)
	r := httptest.NewRequest("POST", "/v1/users", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	a.ServeHTTP(w, r)

	t.Log("Given the need to validate a new user can't be created with an invalid document.")
	{
		if w.Code != 400 {
			t.Fatalf("\tShould received a status code of 400 for the response. Received[%d] %s", w.Code, Failed)
		}
		t.Log("\tShould received a status code of 400 for the response.", Succeed)

		v := struct {
			Error  string `json:"error"`
			Fields []struct {
				Fld string `json:"field_name"`
				Err string `json:"error"`
			} `json:"fields,omitempty"`
		}{}

		if err := json.NewDecoder(w.Body).Decode(&v); err != nil {
			t.Fatal("\tShould be able to unmarshal the response.", Failed)
		}
		t.Log("\tShould be able to unmarshal the response.", Succeed)

		if len(v.Fields) == 0 {
			t.Fatal("\tShould have validation errors in the response.", Failed)
		}
		t.Log("\tShould have validation errors in the response.", Succeed)

		if v.Fields[0].Fld != "FirstName" {
			t.Fatalf("\tShould have a FirstName validation error in the response. Received[%s] %s", v.Fields[0].Fld, Failed)
		}
		t.Log("\tShould have a FirstName validation error in the response.", Succeed)

		if v.Fields[1].Fld != "Addresses" {
			t.Fatalf("\tShould have an Addresses validation error in the response. Received[%s] %s", v.Fields[0].Fld, Failed)
		}
		t.Log("\tShould have an Addresses validation error in the response.", Succeed)
	}
}

// usersList200 validates a users list can be retrieved with the endpoint.
func usersList200(t *testing.T, a *app.App) []models.User {
	r := httptest.NewRequest("GET", "/v1/users", nil)
	w := httptest.NewRecorder()
	a.ServeHTTP(w, r)

	t.Log("Given the need to retrieve a list of users with the users endpoint.")
	{
		if w.Code != 200 {
			t.Fatalf("\tShould received a status code of 200 for the response. Received[%d] %s", w.Code, Failed)
		}
		t.Log("\tShould received a status code of 200 for the response.", Succeed)

		var us []models.User
		if err := json.NewDecoder(w.Body).Decode(&us); err != nil {
			t.Fatal("\tShould be able to unmarshal the response.", Failed)
		}
		t.Log("\tShould be able to unmarshal the response.", Succeed)

		if len(us) == 0 {
			t.Fatal("\tShould have users in the response.", Failed)
		}
		t.Log("\tShould have a users in the response.", Succeed)

		var failed bool
		marks := make([]string, len(us))
		for i, u := range us {
			if u.DateCreated == nil || u.DateModified == nil {
				marks[i] = Failed
				failed = true
			} else {
				marks[i] = Succeed
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
func usersRetrieve200(t *testing.T, a *app.App, id string) {
	r := httptest.NewRequest("GET", "/v1/users/"+id, nil)
	w := httptest.NewRecorder()
	a.ServeHTTP(w, r)

	t.Log("Given the need to retrieve an individual user with the users endpoint.")
	{
		if w.Code != 200 {
			t.Fatalf("\tShould received a status code of 200 for the response. Received[%d] %s", w.Code, Failed)
		}
		t.Log("\tShould received a status code of 200 for the response.", Succeed)

		var ur models.User
		if err := json.NewDecoder(w.Body).Decode(&ur); err != nil {
			t.Fatal("\tShould be able to unmarshal the response.", Failed)
		}
		t.Log("\tShould be able to unmarshal the response.", Succeed)

		if ur.UserID != id {
			t.Fatal("\tShould have the document specified by id.", Failed)
		}
		t.Log("\tShould have the document specified by id", Succeed)
	}
}

// usersRetrieve404 validates a user request for a user that does not exist with the endpoint.
func usersRetrieve404(t *testing.T, a *app.App, id string) {
	r := httptest.NewRequest("GET", "/v1/users/"+id, nil)
	w := httptest.NewRecorder()
	a.ServeHTTP(w, r)

	t.Log("Given the situation of retrieving an individual user that does not exist with the users endpoint.")
	{
		if w.Code != 404 {
			t.Fatalf("\tShould received a status code of 404 for the response. Received[%d] %s", w.Code, Failed)
		}
		t.Log("\tShould received a status code of 404 for the response.", Succeed)
	}
}

// usersRetrieve400 validates a user request with an invalid id with the endpoint.
func usersRetrieve400(t *testing.T, a *app.App, id string) {
	r := httptest.NewRequest("GET", "/v1/users/"+id, nil)
	w := httptest.NewRecorder()
	a.ServeHTTP(w, r)

	t.Log("Given the situation of retrieving an individual user with an invalid id with the users endpoint.")
	{
		if w.Code != 400 {
			t.Fatalf("\tShould received a status code of 400 for the response. Received[%d] %s", w.Code, Failed)
		}
		t.Log("\tShould received a status code of 400 for the response.", Succeed)
	}
}

// usersUpdate200 validates a user can be updated with the endpoint.
func usersUpdate200(t *testing.T, a *app.App) {
	u.FirstName = "Lisa"

	body, _ := json.Marshal(&u)
	r := httptest.NewRequest("PUT", "/v1/users/"+u.UserID, bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	a.ServeHTTP(w, r)

	t.Log("Given the need to validate a user can be updated with the users endpoint.")
	{
		if w.Code != 200 {
			t.Fatalf("\tShould received a status code of 200 for the response. Received[%d] %s", w.Code, Failed)
		}
		t.Log("\tShould received a status code of 200 for the response.", Succeed)

		var resp models.User
		if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
			t.Fatal("\tShould be able to unmarshal the response.", Failed)
		}
		t.Log("\tShould be able to unmarshal the response.", Succeed)

		if resp.UserID != u.UserID {
			t.Fatal("\tShould have an a user value with the same id.", Failed)
		}
		t.Log("\tShould have an a user value with the same id.", Succeed)
	}
}

// usersDelete200 validates a user can be deleted with the endpoint.
func usersDelete200(t *testing.T, a *app.App, id string) {
	r := httptest.NewRequest("DELETE", "/v1/users/"+id, nil)
	w := httptest.NewRecorder()
	a.ServeHTTP(w, r)

	t.Log("Given the need to delete a new user with the users endpoint.")
	{
		if w.Code != 200 {
			t.Fatalf("\tShould received a status code of 200 for the response. Received[%d] %s", w.Code, Failed)
		}
		t.Log("\tShould received a status code of 200 for the response.", Succeed)

		var resp models.User
		if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
			t.Fatal("\tShould be able to unmarshal the response.", Failed)
		}
		t.Log("\tShould be able to unmarshal the response.", Succeed)

		if resp.UserID != id {
			t.Fatal("\tShould have an a user value with the same id.", Failed)
		}
		t.Log("\tShould have an a user value with the same id.", Succeed)
	}
}

// usersDelete404 validates a user that has been deleted is deleted.
func usersDelete404(t *testing.T, a *app.App, id string) {
	r := httptest.NewRequest("DELETE", "/v1/users/"+id, nil)
	w := httptest.NewRecorder()
	a.ServeHTTP(w, r)

	t.Log("Given the need to verify a deleted user is deleted.")
	{
		if w.Code != 404 {
			t.Fatalf("\tShould received a status code of 404 for the response. Received[%d] %s", w.Code, Failed)
		}
		t.Log("\tShould received a status code of 404 for the response.", Succeed)
	}
}
