package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"runtime"
)

func main() {
	http.HandleFunc("/upload", Upload)
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(html))
	})
	http.ListenAndServe(":3000", nil)
}

func Upload(res http.ResponseWriter, req *http.Request) {
	file, handler, err := req.FormFile("myFile")
	if err != nil {
		fmt.Fprint(res, err)
	}

	dirname := uploadDirectoryName()
	os.MkdirAll(dirname, 0777)
	filename := fmt.Sprintf("%s/%s", dirname, handler.Filename)

	outfile, err := os.Create(filename)
	defer outfile.Close()

	_, err = io.Copy(outfile, file)

	if err != nil {
		fmt.Fprint(res, err)
		return
	}

	fmt.Fprintln(res, filename)
}

var uploadDirectoryName = func() string {
	return path.Join(currentDir(), "uploads")
}

func currentDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

var html = `<html>
<head>
    <title>Ultimate Web</title>
    <meta charset="utf-8" />
</head>
<body>
<form action="/upload" method="POST" accept-charset="utf-8" enctype="multipart/form-data">
  <p><input type="file" name="myFile"></p>
  <p><input type="submit" value="Continue →"></p>
</form>
</body>
</html>
`
