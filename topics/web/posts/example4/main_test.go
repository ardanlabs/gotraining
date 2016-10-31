// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Tests for the sample program to show how to handle
// the uploading of file content in a request.
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

// uploadDir is the physical location on disk
// for the uploaded test file.
var uploadDir string

func init() {

	// Capture the location for the uploads directory.
	uploadDir = path.Join(currentDirectory(), "uploads")

	// If the directory exists, remove it.
	os.RemoveAll(uploadDir)
}

func TestApp(t *testing.T) {

	// Startup a server to handle processing these routes.
	ts := httptest.NewServer(App())
	defer ts.Close()

	// Create a sub-test for each verb.
	t.Run("GET", testGet(ts))
	t.Run("POST", testPost(ts))
}

// testGet validates the GET verb.
func testGet(ts *httptest.Server) func(*testing.T) {

	// Test function for execution as a sub-test.
	tf := func(t *testing.T) {

		// Perform a GET call against the url.
		res, err := http.Get(ts.URL)
		if err != nil {
			t.Fatal(err)
		}

		// Read in the response from the call.
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		// Validate we received the correct document.
		got := string(b)
		want := "Continue ->"
		if !strings.Contains(got, want) {
			t.Log("Wanted:", want)
			t.Log("Got   :", got)
			t.Fatal("Mismatch")
		}
	}

	return tf
}

// testPost validates the POST verb.
func testPost(ts *httptest.Server) func(*testing.T) {

	// Test function for execution as a sub-test.
	tf := func(t *testing.T) {

		// Capture the location of the image.
		img := path.Join(currentDirectory(), "gopher.png")

		// Open the image for our test file.
		file, err := os.Open(img)
		if err != nil {
			t.Fatal(err)
		}
		defer file.Close()

		// Create a bytes buffer and use it as the writer
		// for the multipart writer.
		var bb bytes.Buffer
		writer := multipart.NewWriter(&bb)

		// Create a new writer, bound to the bytes buffer, that
		// will help write the form-data header with the specified
		// field and file name when we add the image data.
		part, err := writer.CreateFormFile("myFile", path.Base(img))
		if err != nil {
			t.Fatal(err)
		}

		// Write the contents of the image data into the multipart writer
		// which will fill the bytes buffer with a properly formatted
		// form-data header and file content.
		if _, err := io.Copy(part, file); err != nil {
			t.Fatal(err)
		}

		// You must close the multipart writer else the request
		// will be missing the terminating boundary.
		writer.Close()

		// Create a POST request using the bytes buffer as
		// the post data.
		req, err := http.NewRequest("POST", ts.URL+"/upload", &bb)
		if err != nil {
			t.Fatal(err)
		}

		// Set the content-type properly for this call.
		req.Header.Set("Content-Type", writer.FormDataContentType())

		// Create a client and perform the request. This should
		// save a copy of the file to /uploads.
		var c http.Client
		res, err := c.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		// If we don't get a 200, there is a problem.
		if res.StatusCode != 200 {
			t.Fatalf("Want status 200 got %d", res.StatusCode)
		}

		// Remove this folder after the test.
		defer os.RemoveAll(uploadDir)

		// Open the original image file again.
		orig, err := os.Open(img)
		if err != nil {
			t.Fatal(err)
		}
		defer orig.Close()

		// Read the image data into the slice.
		var origBytes []byte
		orig.Read(origBytes)

		// Open the file we wrote to disk after the request.
		upFile := path.Join(uploadDir, path.Base(img))
		uploaded, err := os.Open(upFile)
		if err != nil {
			t.Fatal(err)
		}
		defer uploaded.Close()

		// Read the image data into the slice.
		var upBytes []byte
		uploaded.Read(upBytes)

		// Validate we received the correct document.
		got := len(upBytes)
		want := len(origBytes)
		if got != want {
			t.Log("Wanted:", want)
			t.Log("Got   :", got)
			t.Fatal("Mismatch")
		}
	}

	return tf
}
