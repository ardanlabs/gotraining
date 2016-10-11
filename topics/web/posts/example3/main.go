package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/schema"
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
<p>
	<input type="text" name="FirstName" placeholder="First Name" />
</p>
<p>
	<input type="text" name="LastName" placeholder="Last Name" />
</p>
<p>
	<input type="submit" value="CLICK ME!!" />
</p>
</form>
	`))
}

type User struct {
	FirstName string
	LastName  string
}

func PostHandler(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}

	u := User{}

	decoder := schema.NewDecoder()
	err = decoder.Decode(&u, req.PostForm)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}

	fmt.Fprintf(res, "First Name: %s\nLast Name: %s", u.FirstName, u.LastName)
}

func main() {
	log.Fatal(http.ListenAndServe(":3000", App()))
}
