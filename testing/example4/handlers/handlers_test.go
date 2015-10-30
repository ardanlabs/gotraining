// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample test to show how to test the execution of an internal endpoint.
package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ardanlabs/gotraining/10-testing/01-testing/example4/handlers"
)

const succeed = "\u2713"
const failed = "\u2717"

func init() {
	handlers.Routes()
}

// TestSendJSON testing the sendjson internal endpoint.
func TestSendJSON(t *testing.T) {
	t.Log("Given the need to test the SendJSON endpoint.")
	{
		r, _ := http.NewRequest("GET", "/sendjson", nil)
		w := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(w, r)

		if w.Code != 200 {
			t.Fatalf("\tShould receive a status code of \"200\" for the response. Received[%d] %s", w.Code, failed)
		}
		t.Log("\tShould receive a status code of \"200\" for the response.", succeed)

		u := struct {
			Name  string
			Email string
		}{}
		if err := json.NewDecoder(w.Body).Decode(&u); err != nil {
			t.Fatal("\tShould be able to decode the response.", failed)
		}
		t.Log("\tShould be able to decode the response.", succeed)

		if u.Name == "Bill" {
			t.Log("\tShould have \"Bill\" for Name in the response.", succeed)
		} else {
			t.Error("\tShould have \"Bill\" for Name in the response.", failed, u.Name)
		}

		if u.Email == "bill@ardanlabs.com" {
			t.Log("\tShould have \"bill@ardanlabs.com\" for Email in the response.", succeed)
		} else {
			t.Error("\tShould have \"bill@ardanlabs.com\" for Email in the response.", failed, u.Email)
		}
	}
}
