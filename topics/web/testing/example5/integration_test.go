package unit

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

type App struct{}

func (a App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello World!")
}

func Test_MyHandler(t *testing.T) {
	r := require.New(t)

	ts := httptest.NewServer(App{})
	defer ts.Close()

	res, _ := http.Get(ts.URL)

	body, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	r.Equal("Hello World!", string(body))
}
