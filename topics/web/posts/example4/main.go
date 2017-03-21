// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to handle the uploading
// of file content in a request.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"runtime"
)

// html contains the index page document for the
// web application.
var html = `
<html>
<head>
	<title>Ultimate Web</title>
	<meta charset="utf-8" />
</head>
<body>
<form action="/upload" method="POST" accept-charset="utf-8" enctype="multipart/form-data">
	<p><input type="file" name="myFile"></p>
	<p><input type="submit" value="Continue ->"></p>
</form>
</body>
</html>`

// App provides a handler to handle GET and POST calls
// for every request.
func App() http.Handler {

	// Create a new mux and bind the index page and
	// the upload route.
	m := http.NewServeMux()
	m.HandleFunc("/upload", UploadHandler)
	m.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(html))
	})

	// Return the mux to handle these routes.
	return m
}

// UploadHandler handles the uploading of content.
func UploadHandler(res http.ResponseWriter, req *http.Request) {

	// Retrieve the first file for the provided form key.
	file, handler, err := req.FormFile("myFile")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Capture the location for the uploads directory and
	// set the appropriate rights.
	uploadDir := path.Join(currentDirectory(), "uploads")
	os.MkdirAll(uploadDir, 0777)

	// Generate a file name for the uploaded file.
	filename := path.Join(uploadDir, handler.Filename)

	// Create the new file we need for the uploaded file.
	outfile, err := os.Create(filename)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	defer outfile.Close()

	// Copy the file from the request into the file on disk.
	if _, err = io.Copy(outfile, file); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the physical location on disk where the
	// file was stored.
	fmt.Fprintln(res, filename)
}

func main() {

	// Start the http server to handle the request for
	// both versions of the API.
	log.Fatal(http.ListenAndServe(":3000", App()))
}

// Returns the current directory we are running in.
func currentDirectory() string {

	// Locate the current directory for the site.
	_, fn, _, _ := runtime.Caller(1)
	return path.Dir(fn)
}
