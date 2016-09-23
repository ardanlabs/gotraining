package unit

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func MyHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello World!")
}

func Test_MyHandler(t *testing.T) {
	r := require.New(t)

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	res := httptest.NewRecorder()

	m := http.NewServeMux()
	m.HandleFunc("/", MyHandler)
	m.ServeHTTP(res, req)

	r.Equal("Hello World!", res.Body.String())
}
