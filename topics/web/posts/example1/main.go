package main

import (
	"log"
	"net/http"
)

func router(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		getHandler(res, req)
	case "POST":
		postHandler(res, req)
	default:
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	res.Write([]byte(`
<form action="/" method="POST">
<input type="submit" value="CLICK ME!!" />
</form>
	`))
}

func postHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Thank you for clicking me!"))
}

func main() {
	http.HandleFunc("/", router)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
