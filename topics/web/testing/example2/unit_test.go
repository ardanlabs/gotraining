package unit

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type App struct{}

func (a App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello World!")
}

func Test_App(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	res := httptest.NewRecorder()

	a := App{}
	a.ServeHTTP(res, req)

	exp := "Hello World!"
	act := res.Body.String()
	if act != exp {
		t.Fatalf("expected %s got %s", exp, act)
	}
}
