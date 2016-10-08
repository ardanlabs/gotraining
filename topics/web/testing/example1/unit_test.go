package unit

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func MyHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello World!")
}

func Test_MyHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	res := httptest.NewRecorder()
	MyHandler(res, req)

	exp := "Hello World!"
	act := res.Body.String()
	if act != exp {
		t.Fatalf("expected %s got %s", exp, act)
	}
}
