// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample test to show how to write a unit test in BDD style
package unit

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

// MyHandler is the application handler we want to test. It wouldn't
// be in this file, it would be in another file in the same package.
func MyHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello World!")
}

func Test_MyHandler(t *testing.T) {

	Convey("Given a HTTP request for http://example.com/foo", t, func() {

		// Create a new request.
		req := httptest.NewRequest("GET", "http://example.com/foo", nil)

		// Create a new ResponseRecorder which implements
		// the ResponseWriter interface.
		res := httptest.NewRecorder()

		// Execute the handler function directly.
		MyHandler(res, req)

		// Validate we received the expected response.
		Convey("Then the response should be \"Hello World!\" ", func() {
			So(res.Body.String(), ShouldEqual, "Hello World!")
		})
	})
}
