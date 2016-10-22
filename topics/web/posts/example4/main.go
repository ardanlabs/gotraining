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

var html = `<html>
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
</html>
`

func App() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/upload", UploadHandler)
	m.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(html))
	})
	return m
}

func UploadHandler(res http.ResponseWriter, req *http.Request) {
	file, handler, err := req.FormFile("myFile")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	dirname := uploadDirectoryName()
	os.MkdirAll(dirname, 0777)
	filename := fmt.Sprintf("%s/%s", dirname, handler.Filename)

	outfile, err := os.Create(filename)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	defer outfile.Close()

	_, err = io.Copy(outfile, file)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(res, filename)
}

func main() {
	log.Panic(http.ListenAndServe(":3000", App()))
}

var uploadDirectoryName = func() string {
	return path.Join(currentDir(), "uploads")
}

func currentDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}
