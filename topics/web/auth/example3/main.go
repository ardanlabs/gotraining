// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to use JWTs
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// secret is the key we'll use to sign / ensure tokens. You may choose to use a
// different key per consumer or a rolling set of keys.
const secret = "something-secret"

// loginHandler receives a login request and issues a JWT
func loginHandler(res http.ResponseWriter, req *http.Request) {

	// info is the information we expect to be sent
	var info struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(req.Body).Decode(&info); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	// Check login. In a real app we would look up the user by email, hash
	// their password, then securely compare the hash with the stored hash.
	if info.Email != "jacob@example.com" || info.Password != "rory" {
		http.Error(res, "Could not authenticate", http.StatusUnauthorized)
		return
	}

	// Define a claims. This is the metadata we'll store in the token. We could
	// make our own type but for now we'll just use a standard type.
	now := time.Now().UTC()
	claims := jwt.StandardClaims{
		Subject:   info.Email,
		NotBefore: now.Unix(),
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(30 * time.Hour).Unix(),
	}

	// Generate the token and sign it using SHA256 and our secret key
	tkn, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
	if err != nil {
		log.Printf("Unable to generate JWT for user: %v", err)
		http.Error(res, "Could not generate token", http.StatusInternalServerError)
		return
	}

	// Give the token to the user
	res.Write([]byte(tkn))
}

// secureHandler delivers a secret message if the request contains a valid JWT.
func secureHandler(res http.ResponseWriter, req *http.Request) {

	// Get the token from the authentication header. We are going to chop off
	// the first 7 bytes for "Bearer " so we handle the empty case and the
	// too-short case together.
	tkn := req.Header.Get("Authentication")
	if len(tkn) < 7 {
		http.Error(res, "Not authorized", http.StatusUnauthorized)
		return
	}
	tkn = tkn[7:]

	// keyfunc is responsible for identifying the key for decoding a jwt. It is
	// passed a *jwt.Token which it may use in the process of finding the key.
	// In our case we're using a single value so we just return it.
	keyfunc := func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	}

	// Create a parser and parse the token. Here I've restricted the parser to
	// only use the different SHA methods. We use ParseWithClaims so we can
	// control the type of value used for storing the metadata.
	p := jwt.Parser{
		ValidMethods: []string{"HS256", "HS384", "HS512"},
	}
	t, err := p.ParseWithClaims(tkn, &jwt.StandardClaims{}, keyfunc)

	// Token is invalid
	if err != nil {
		http.Error(res, "Not authorized", http.StatusUnauthorized)
		return
	}

	// If we wanted to know more about the authenticated user we can use a type
	// assertion to get a useful claims value out of the token.
	claims := t.Claims.(*jwt.StandardClaims)
	log.Println("Serving secret to", claims.Subject)

	// Serve up the secret!
	fmt.Fprintln(res, "Be sure to drink your Ovaltine!")
}

// App loads the API for use.
func App() http.Handler {

	// Create a mux for binding routes.
	m := http.NewServeMux()

	m.HandleFunc("/login", loginHandler)
	m.HandleFunc("/secure", secureHandler)

	return m
}

func main() {

	// Start the http server to handle the request for
	// both versions of the API.
	log.Print("Listening on localhost:3000")
	log.Fatal(http.ListenAndServe("localhost:3000", App()))
}
