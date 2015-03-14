package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type user struct {
	Name  string
	Email string
	Phone string
}

var users []user

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", usersHandler)
	mux.HandleFunc("/search", searchUsers)
	// This will act as a catch all for the rest of the routes
	mux.Handle("/", http.FileServer(http.Dir("public")))

	http.ListenAndServe(":4000", mux)
}

func usersHandler(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// List the users
		respondJSON(rw, http.StatusOK, users)

	case "POST":
		u := user{
			Name:  r.PostFormValue("name"),
			Email: r.PostFormValue("email"),
			Phone: r.PostFormValue("phone"),
		}
		users = append(users, u)
		http.Redirect(rw, r, "/users", http.StatusSeeOther)

	default:
		http.Error(rw, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func searchUsers(rw http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(rw, "Query is required", http.StatusBadRequest)
		return
	}

	for _, u := range users {
		if strings.Contains(u.Name, query) {
			respondJSON(rw, http.StatusOK, u)
			return
		}
	}

	http.Error(rw, "Not Found", http.StatusNotFound)
}

func respondJSON(w http.ResponseWriter, code int, val interface{}) error {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(val)
}
