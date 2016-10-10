package main

import (
	"log"
	"net/http"
)

func App() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "GET":
			GetHandler(res, req)
		case "POST":
			PostHandler(res, req)
		default:
			res.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}

func GetHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	res.Write([]byte(`
<form action="/" method="POST">
<input type="submit" value="CLICK ME!!" />
</form>
	`))
}

func PostHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Thank you for clicking me!"))
}

func main() {
	log.Fatal(http.ListenAndServe(":3000", App()))
}
