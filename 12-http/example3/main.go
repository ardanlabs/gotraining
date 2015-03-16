// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// Sample program to show how to use a ServeMux from the standard
// library. How to handle verbs and more complex routing.
package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

// user represents a user in the system.
type user struct {
	Name  string
	Email string
	Phone string
}

// users is a slice of users.
var users []user

func main() {
	// Create a ServeMux and add some routes.
	mux := http.NewServeMux()
	mux.HandleFunc("/users", usersHandler)
	mux.HandleFunc("/search", searchUsers)

	// This will act as a catch all for the rest of the routes
	mux.Handle("/", http.FileServer(http.Dir("public")))

	// Start the service.
	http.ListenAndServe(":4000", mux)
}

// usersHandler handles the /users api call.
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

// searchUsers handles the /search api call.
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

// respondJSON writes the reponse for the api back to the caller
// in JSON.
func respondJSON(w http.ResponseWriter, code int, val interface{}) error {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(val)
}
