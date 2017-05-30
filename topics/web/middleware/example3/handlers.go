package main

import "net/http"

// helloWorld is one of our application handlers. It greets the world.
func helloWorld(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello World"))
}

// secret is a handler which reveals a secret phrase. It should be
// behind a middleware that ensures we don't just give it out freely.
func secret(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Be sure to drink your Ovaltine!"))
}
