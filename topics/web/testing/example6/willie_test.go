package unit

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/markbates/willie"
	"github.com/stretchr/testify/require"
)

type App struct{}

func (a App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello World!")
}

func Test_MyHandler(t *testing.T) {
	r := require.New(t)

	w := willie.New(App{})
	res := w.Request("/").Get()

	r.Equal("Hello World!", res.Body.String())
}
