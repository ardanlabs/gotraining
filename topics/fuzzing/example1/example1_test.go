// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Tests to validate the api endpoints.
package api_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ardanlabs/gotraining/topics/fuzzing/example1"
)

const succeed = "\u2713"
const failed = "\u2717"

func init() {
	api.Routes()
}

// TestProcess tests the Process endpoint with proper data.
func TestProcess(t *testing.T) {
	tests := []struct {
		url    string
		status int
		val    []byte
		resp   string
	}{
		{"/process", http.StatusOK, []byte("ADM46Bill,ADM42Lisa,DEV35John,USR46Eduardo"), `[{"Type":"ADM","Name":"Bill","Age":46},{"Type":"ADM","Name":"Lisa","Age":42},{"Type":"DEV","Name":"John","Age":35},{"Type":"USR","Name":"Eduardo","Age":46}]`},
	}

	t.Log("Given the need to test the Process endpoint.")
	{
		for i, tt := range tests {
			t.Logf("\tTest %d:\tWhen checking %q for status code %d with data %s", i, tt.url, tt.status, tt.val)
			{
				r, _ := http.NewRequest("POST", tt.url, bytes.NewBuffer(tt.val))
				w := httptest.NewRecorder()
				http.DefaultServeMux.ServeHTTP(w, r)

				if w.Code != tt.status {
					t.Fatalf("\t%s\tShould receive a status code of %d for the response. Received[%d].", failed, tt.status, w.Code)
				}
				t.Logf("\t%s\tShould receive a status code of %d for the response.", succeed, tt.status)

				recv := w.Body.String()

				if tt.resp != recv[:len(recv)-1] {
					t.Log("GOT:", recv)
					t.Log("EXP:", tt.resp)
					t.Fatalf("\t%s\tShould get the expected result.", failed)
				}
				t.Logf("\t%s\tShould get the expected result.", succeed)
			}
		}
	}
}
