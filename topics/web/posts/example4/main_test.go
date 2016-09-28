package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"
)

func TestUpload(t *testing.T) {
	img := path.Join(currentDir(), "gopher.png")
	file, err := os.Open(img)
	if err != nil {
		t.Fatal(err)
	}

	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("myFile", path.Base(img))
	io.Copy(part, file)

	writer.Close()

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	dirname := path.Join(currentDir(), "test_uploads")
	defer os.RemoveAll(dirname)
	os.RemoveAll(dirname)
	uploadDirectoryName = func() string {
		return dirname
	}

	Upload(res, req)

	if res.Code != 200 {
		t.Errorf("Expected %d to equal 200", res.Code)
		t.FailNow()
	}
	upFile := path.Join(dirname, path.Base(img))
	os.Stat(upFile)

	orig, _ := os.Open(img)
	defer orig.Close()
	origBytes := []byte{}
	orig.Read(origBytes)

	uploaded, _ := os.Open(upFile)
	defer uploaded.Close()
	upBytes := []byte{}
	uploaded.Read(upBytes)

	if string(origBytes) != string(upBytes) {
		t.Error("Original file bytes and uploaded file bytes aren't equal")
		t.FailNow()
	}
}
