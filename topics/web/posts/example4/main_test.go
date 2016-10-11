package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"strings"
	"testing"
)

var uploadDir string

func init() {
	uploadDir = path.Join(currentDir(), "test_uploads")
	os.RemoveAll(uploadDir)
	uploadDirectoryName = func() string {
		return uploadDir
	}
}

func Test_App(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()
	t.Run("GET", test_Get(ts))
	t.Run("POST", test_Post(ts))
}

func test_Get(ts *httptest.Server) func(*testing.T) {
	return func(t *testing.T) {
		res, err := http.Get(ts.URL)
		if err != nil {
			t.Fatal(err)
		}

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		act := string(b)
		exp := "Continue ->"

		if !strings.Contains(act, exp) {
			t.Fatalf("expected %s to contain %s", act, exp)
		}
	}
}

func test_Post(ts *httptest.Server) func(*testing.T) {
	return func(t *testing.T) {
		img := path.Join(currentDir(), "gopher.png")
		file, err := os.Open(img)
		if err != nil {
			t.Fatal(err)
		}

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, _ := writer.CreateFormFile("myFile", path.Base(img))
		io.Copy(part, file)

		req, err := http.NewRequest("POST", ts.URL, body)
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Content-Type", writer.FormDataContentType())

		c := http.Client{}
		res, err := c.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		if res.StatusCode != 200 {
			t.Fatalf("Expected %d to equal 200", res.StatusCode)
		}

		orig, _ := os.Open(img)
		origBytes := []byte{}
		orig.Read(origBytes)

		upFile := path.Join(uploadDir, path.Base(img))
		uploaded, _ := os.Open(upFile)
		upBytes := []byte{}
		uploaded.Read(upBytes)

		act := string(upBytes)
		exp := string(origBytes)

		if act != exp {
			t.Fatal("Original file bytes and uploaded file bytes aren't equal")
		}
	}
}
