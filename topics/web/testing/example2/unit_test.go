package unit

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

type App struct{}

func (a App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello World!")
}

func Test_App(t *testing.T) {
	r := require.New(t)

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	res := httptest.NewRecorder()

	a := App{}
	a.ServeHTTP(res, req)

	r.Equal("Hello World!", res.Body.String())
}
