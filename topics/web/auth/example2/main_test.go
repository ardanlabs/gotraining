package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/faux"
)

func TestIndex(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	want := "Log in with Github"
	got := string(b)
	if !strings.Contains(got, want) {
		t.Logf("Wanted: %s", want)
		t.Logf("Got   : %s", got)
		t.Fail()
	}
}

func TestCallback(t *testing.T) {
	gothic.Store = sessions.NewFilesystemStore(os.TempDir(), []byte("goth-test"))
	goth.UseProviders(&faux.Provider{})

	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/auth/faux/callback", nil)
	if err != nil {
		t.Fatal(err)
	}

	sess := faux.Session{Name: "Mary Jane", Email: "mary@example.com"}
	session, _ := gothic.Store.Get(req, gothic.SessionName)
	session.Values[gothic.SessionName] = sess.Marshal()
	err = session.Save(req, res)
	if err != nil {
		t.Fatal(err)
	}

	App().ServeHTTP(res, req)

	want := "Mary Jane"
	got := res.Body.String()
	if !strings.Contains(got, want) {
		t.Logf("Wanted: %s", want)
		t.Logf("Got   : %s", got)
		t.Fail()
	}
}
