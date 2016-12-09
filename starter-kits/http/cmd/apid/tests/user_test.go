// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package tests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ardanlabs/gotraining/starter-kits/http/cmd/apid/routes"
	"github.com/ardanlabs/gotraining/starter-kits/http/internal/sys/web"
	"github.com/ardanlabs/gotraining/starter-kits/http/internal/user"
	"gopkg.in/mgo.v2/bson"
)

const (
	// Succeed is the Unicode codepoint for a check mark.
	Succeed = "\u2713"

	// Failed is the Unicode codepoint for an X mark.
	Failed = "\u2717"
)

//==============================================================================

// The web application state for the
var a *web.App

// TestMain helps to clean up the test data.
func TestMain(m *testing.M) {
	os.Exit(runTest(m))
}

// runTest initializes the environment for the tests and allows for
// the proper return code if the test fails or succeeds.
func runTest(m *testing.M) int {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	a = routes.API().(*web.App)

	return m.Run()
}

//==============================================================================

// TestUsers is the entry point for the users
func TestUsers(t *testing.T) {

	t.Run("getUsers200Empty", getUsers200Empty)
	t.Run("postUser400", postUser400)
	t.Run("getUser404", getUser404)
	t.Run("getUser400", getUser400)
	t.Run("deleteUser404", deleteUser404)
	t.Run("putUser404", putUser404)
	t.Run("crudUsers", crudUser)
}

// getUsers200Empty validates an empty users list can be retrieved with the endpoint.
func getUsers200Empty(t *testing.T) {
	r := httptest.NewRequest("GET", "/v1/users", nil)
	w := httptest.NewRecorder()
	a.ServeHTTP(w, r)

	t.Log("Given the need to validate an empty list of users with the users endpoint.")
	{
		t.Log("\tTest 0:\tWhen using the default user value.")
		{
			if w.Code != http.StatusOK {
				t.Fatalf("\t%s\tShould received a status code of 200 for the response : %v", Failed, w.Code)
			}
			t.Logf("\t%s\tShould received a status code of 200 for the response.", Succeed)
		}
	}
}

// postUser400 validates a user can't be created with the endpoint
// unless a valid user document is submitted.
func postUser400(t *testing.T) {
	u := user.User{
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
		t.Log("\tTest 0:\tWhen using an incomplete user value.")
		{
			if w.Code != http.StatusBadRequest {
				t.Fatalf("\t%s\tShould received a status code of 400 for the response : %v", Failed, w.Code)
			}
			t.Logf("\t%s\tShould received a status code of 400 for the response.", Succeed)

			recv := w.Body.String()
			resps := []string{
				`{
  "error": "field validation failure",
  "fields": [
    {
      "field_name": "Addresses",
      "error": "required"
    },
    {
      "field_name": "FirstName",
      "error": "required"
    }
  ]
}`,
				`{
  "error": "field validation failure",
  "fields": [
    {
      "field_name": "FirstName",
      "error": "required"
    },
    {
      "field_name": "Addresses",
      "error": "required"
    }
  ]
}`,
			}

			var found bool
			for _, resp := range resps {
				if resp == recv {
					found = true
					break
				}
			}

			if !found {
				t.Log("Got :", recv)
				t.Log("Want:", resps[0])
				t.Log("Want:", resps[1])
				t.Fatalf("\t%s\tShould get the expected result.", Failed)
			}
			t.Logf("\t%s\tShould get the expected result.", Succeed)
		}
	}
}

// getUser404 validates a user request for a malformed userid.
func getUser404(t *testing.T) {
	userID := "12345"

	r := httptest.NewRequest("GET", "/v1/users/"+userID, nil)
	w := httptest.NewRecorder()
	a.ServeHTTP(w, r)

	t.Log("Given the need to validate getting a user with a malformed userid.")
	{
		t.Logf("\tTest 0:\tWhen using the new user %s.", userID)
		{
			if w.Code != http.StatusBadRequest {
				t.Fatalf("\t%s\tShould received a status code of 400 for the response : %v", Failed, w.Code)
			}
			t.Logf("\t%s\tShould received a status code of 400 for the response.", Succeed)

			recv := w.Body.String()
			resp := `{
  "error": "ID is not in it's proper form"
}`
			if resp != recv {
				t.Log("Got :", recv)
				t.Log("Want:", resp)
				t.Fatalf("\t%s\tShould get the expected result.", Failed)
			}
			t.Logf("\t%s\tShould get the expected result.", Succeed)
		}
	}
}

// getUser400 validates a user request for a user that does not exist with the endpoint.
func getUser400(t *testing.T) {
	userID := bson.NewObjectId().Hex()

	r := httptest.NewRequest("GET", "/v1/users/"+userID, nil)
	w := httptest.NewRecorder()
	a.ServeHTTP(w, r)

	t.Log("Given the need to validate getting a user with an unknown id.")
	{
		t.Logf("\tTest 0:\tWhen using the new user %s.", userID)
		{
			if w.Code != http.StatusNotFound {
				t.Fatalf("\t%s\tShould received a status code of 404 for the response : %v", Failed, w.Code)
			}
			t.Logf("\t%s\tShould received a status code of 404 for the response.", Succeed)

			recv := w.Body.String()
			resp := `{
  "error": "Entity not found"
}`
			if resp != recv {
				t.Log("Got :", recv)
				t.Log("Want:", resp)
				t.Fatalf("\t%s\tShould get the expected result.", Failed)
			}
			t.Logf("\t%s\tShould get the expected result.", Succeed)
		}
	}
}

// deleteUser404 validates deleting a user that does not exist.
func deleteUser404(t *testing.T) {
	userID := bson.NewObjectId().Hex()

	r := httptest.NewRequest("DELETE", "/v1/users/"+userID, nil)
	w := httptest.NewRecorder()
	a.ServeHTTP(w, r)

	t.Log("Given the need to validate deleting a user that does not exist.")
	{
		t.Logf("\tTest 0:\tWhen using the new user %s.", userID)
		{
			if w.Code != http.StatusNotFound {
				t.Fatalf("\t%s\tShould received a status code of 404 for the response : %v", Failed, w.Code)
			}
			t.Logf("\t%s\tShould received a status code of 404 for the response.", Succeed)

			recv := w.Body.String()
			resp := `{
  "error": "Entity not found"
}`
			if resp != recv {
				t.Log("Got :", recv)
				t.Log("Want:", resp)
				t.Fatalf("\t%s\tShould get the expected result.", Failed)
			}
			t.Logf("\t%s\tShould get the expected result.", Succeed)
		}
	}
}

// putUser404 validates updating a user that does not exist.
func putUser404(t *testing.T) {
	u := user.User{
		UserType:  1,
		FirstName: "Bill",
		LastName:  "Kennedy",
		Email:     "bill@ardanstugios.com",
		Company:   "Ardan Labs",
		Addresses: []user.Address{
			{},
		},
	}

	userID := bson.NewObjectId().Hex()

	body, _ := json.Marshal(&u)
	r := httptest.NewRequest("PUT", "/v1/users/"+userID, bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	a.ServeHTTP(w, r)

	t.Log("Given the need to validate updating a user that does not exist.")
	{
		t.Logf("\tTest 0:\tWhen using the new user %s.", userID)
		{
			if w.Code != http.StatusNotFound {
				t.Fatalf("\t%s\tShould received a status code of 404 for the response : %v", Failed, w.Code)
			}
			t.Logf("\t%s\tShould received a status code of 404 for the response.", Succeed)

			recv := w.Body.String()
			resp := `{
  "error": "Entity not found"
}`
			if resp != recv {
				t.Log("Got :", recv)
				t.Log("Want:", resp)
				t.Fatalf("\t%s\tShould get the expected result.", Failed)
			}
			t.Logf("\t%s\tShould get the expected result.", Succeed)
		}
	}
}

// crudUser performs a complete test of CRUD against the api.
func crudUser(t *testing.T) {
	nu := postUser201(t)
	defer deleteUser204(t, nu.UserID)

	getUser200(t, nu.UserID)
	putUser204(t, nu)
}

// postUser201 validates a user can be created with the endpoint.
func postUser201(t *testing.T) user.User {
	var u = user.User{
		UserType:  1,
		FirstName: "Bill",
		LastName:  "Kennedy",
		Email:     "bill@ardanlabs.com",
		Company:   "Ardan Labs",
		Addresses: []user.Address{
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

	var newUser user.User

	body, _ := json.Marshal(&u)
	r := httptest.NewRequest("POST", "/v1/users", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	a.ServeHTTP(w, r)

	t.Log("Given the need to create a new user with the users endpoint.")
	{
		t.Log("\tTest 0:\tWhen using the declared user value.")
		{
			if w.Code != http.StatusCreated {
				t.Fatalf("\t%s\tShould received a status code of 201 for the response : %v", Failed, w.Code)
			}
			t.Logf("\t%s\tShould received a status code of 201 for the response.", Succeed)

			var u user.User
			if err := json.NewDecoder(w.Body).Decode(&u); err != nil {
				t.Fatalf("\t%s\tShould be able to unmarshal the response : %v", Failed, err)
			}

			newUser = u

			u.UserID = "1234"
			u.DateCreated = nil
			u.DateModified = nil
			u.Addresses[0].DateCreated = nil
			u.Addresses[0].DateModified = nil

			doc, err := json.Marshal(&u)
			if err != nil {
				t.Fatalf("\t%s\tShould be able to marshal the response : %v", Failed, err)
			}

			recv := string(doc)
			resp := `{"user_id":"1234","type":1,"first_name":"Bill","last_name":"Kennedy","email":"bill@ardanlabs.com","company":"Ardan Labs","addresses":[{"type":1,"line_one":"12973 SW 112th ST","line_two":"Suite 153","city":"Miami","state":"FL","zipcode":"FL","phone":"305-527-3353","date_modified":null,"date_created":null}],"date_modified":null,"date_created":null}`

			if resp != recv {
				t.Log("Got :", recv)
				t.Log("Want:", resp)
				t.Fatalf("\t%s\tShould get the expected result.", Failed)
			}
			t.Logf("\t%s\tShould get the expected result.", Succeed)
		}
	}

	return newUser
}

// deleteUser200 validates deleting a user that does exist.
func deleteUser204(t *testing.T, userID string) {
	r := httptest.NewRequest("DELETE", "/v1/users/"+userID, nil)
	w := httptest.NewRecorder()
	a.ServeHTTP(w, r)

	t.Log("Given the need to validate deleting a user that does exist.")
	{
		t.Logf("\tTest 0:\tWhen using the new user %s.", userID)
		{
			if w.Code != http.StatusNoContent {
				t.Fatalf("\t%s\tShould received a status code of 204 for the response : %v", Failed, w.Code)
			}
			t.Logf("\t%s\tShould received a status code of 204 for the response.", Succeed)
		}
	}
}

// getUser200 validates a user request for an existing userid.
func getUser200(t *testing.T, userID string) {
	r := httptest.NewRequest("GET", "/v1/users/"+userID, nil)
	w := httptest.NewRecorder()
	a.ServeHTTP(w, r)

	t.Log("Given the need to validate getting a user that exsits.")
	{
		t.Logf("\tTest 0:\tWhen using the new user %s.", userID)
		{
			if w.Code != http.StatusOK {
				t.Fatalf("\t%s\tShould received a status code of 200 for the response : %v", Failed, w.Code)
			}
			t.Logf("\t%s\tShould received a status code of 200 for the response.", Succeed)

			var u user.User
			if err := json.NewDecoder(w.Body).Decode(&u); err != nil {
				t.Fatalf("\t%s\tShould be able to unmarshal the response : %v", Failed, err)
			}

			u.UserID = "1234"
			u.DateCreated = nil
			u.DateModified = nil
			u.Addresses[0].DateCreated = nil
			u.Addresses[0].DateModified = nil

			doc, err := json.Marshal(&u)
			if err != nil {
				t.Fatalf("\t%s\tShould be able to marshal the response : %v", Failed, err)
			}

			recv := string(doc)
			resp := `{"user_id":"1234","type":1,"first_name":"Bill","last_name":"Kennedy","email":"bill@ardanlabs.com","company":"Ardan Labs","addresses":[{"type":1,"line_one":"12973 SW 112th ST","line_two":"Suite 153","city":"Miami","state":"FL","zipcode":"FL","phone":"305-527-3353","date_modified":null,"date_created":null}],"date_modified":null,"date_created":null}`

			if resp != recv {
				t.Log("Got :", recv)
				t.Log("Want:", resp)
				t.Fatalf("\t%s\tShould get the expected result.", Failed)
			}
			t.Logf("\t%s\tShould get the expected result.", Succeed)
		}
	}
}

// putUser204 validates updating a user that does exist.
func putUser204(t *testing.T, u user.User) {
	u.FirstName = "Lisa"
	u.Email = "lisa@email.com"
	u.Addresses[0].State = "NY"

	body, _ := json.Marshal(&u)
	r := httptest.NewRequest("PUT", "/v1/users/"+u.UserID, bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	a.ServeHTTP(w, r)

	t.Log("Given the need to update a user with the users endpoint.")
	{
		t.Log("\tTest 0:\tWhen using the modified user value.")
		{
			if w.Code != http.StatusNoContent {
				t.Fatalf("\t%s\tShould received a status code of 204 for the response : %v", Failed, w.Code)
			}
			t.Logf("\t%s\tShould received a status code of 204 for the response.", Succeed)

			r = httptest.NewRequest("GET", "/v1/users/"+u.UserID, nil)
			w = httptest.NewRecorder()
			a.ServeHTTP(w, r)

			if w.Code != http.StatusOK {
				t.Fatalf("\t%s\tShould received a status code of 200 for the retrieve : %v", Failed, w.Code)
			}
			t.Logf("\t%s\tShould received a status code of 200 for the retrieve.", Succeed)

			var ru user.User
			if err := json.NewDecoder(w.Body).Decode(&ru); err != nil {
				t.Fatalf("\t%s\tShould be able to unmarshal the response : %v", Failed, err)
			}

			ru.UserID = "1234"
			ru.DateCreated = nil
			ru.DateModified = nil
			ru.Addresses[0].DateCreated = nil
			ru.Addresses[0].DateModified = nil

			doc, err := json.Marshal(&ru)
			if err != nil {
				t.Fatalf("\t%s\tShould be able to marshal the response : %v", Failed, err)
			}

			recv := string(doc)
			resp := `{"user_id":"1234","type":1,"first_name":"Lisa","last_name":"Kennedy","email":"lisa@email.com","company":"Ardan Labs","addresses":[{"type":1,"line_one":"12973 SW 112th ST","line_two":"Suite 153","city":"Miami","state":"NY","zipcode":"FL","phone":"305-527-3353","date_modified":null,"date_created":null}],"date_modified":null,"date_created":null}`

			if resp != recv {
				t.Log("Got :", recv)
				t.Log("Want:", resp)
				t.Fatalf("\t%s\tShould get the expected result.", Failed)
			}
			t.Logf("\t%s\tShould get the expected result.", Succeed)
		}
	}
}
