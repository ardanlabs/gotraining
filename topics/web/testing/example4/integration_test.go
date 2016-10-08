package unit

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func MyHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello World!")
}

func Test_MyHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(MyHandler))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	exp := "Hello World!"
	act := string(body)
	if act != exp {
		t.Fatalf("expected %s got %s", exp, act)
	}
}
