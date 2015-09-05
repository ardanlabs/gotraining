// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/SIk8XWmwWa

// Sample program to show how to use a ServeMux from the standard
// library. How to handle verbs and more complex routing.
package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strings"
	"sync"
)

// user represents a user in the system.
type user struct {
	Name  string
	Email string
	Phone string
}

// app contains package level state for the web service.
var app = struct {
	// Mutex for safe access.
	mu sync.RWMutex

	// In-Memory data store. Initialized to avoid JSON null.
	users []user

	// idxTpl maintains access to the page template.
	idxTpl *template.Template

	// fs is a file serving handler for static files.
	fs http.Handler
}{
	users:  make([]user, 0),
	idxTpl: template.Must(template.ParseFiles("template/index.html")),
	fs:     http.FileServer(http.Dir("public")),
}

// baseHandler handles serving the index template and static assets.
func baseHandler(w http.ResponseWriter, r *http.Request) {
	// If an static asset is being requested, use the file server.
	if r.URL.Path != "/" {
		app.fs.ServeHTTP(w, r)
		return
	}

	// Capture a safe copy of the slice header.
	var data []user
	app.mu.RLock()
	{
		data = app.users
	}
	app.mu.RUnlock()

	// Create a value for the template with this
	// slice header.
	users := struct {
		Users []user
	}{
		data,
	}

	// Execute the template for the users and send the
	// result to the client.
	app.idxTpl.Execute(w, users)
}

// respondJSON sends status and writes JSON to the client.
func respondJSON(w http.ResponseWriter, status int, val interface{}) error {
	// Standard way to respond back to the client.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(val)
}

// usersHandler handles the /api/v1/users path.
func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET", "HEAD":
		// Capture a safe copy of the slice header.
		var users []user
		app.mu.RLock()
		{
			users = app.users
		}
		app.mu.RUnlock()

		// Response to the client with these users.
		respondJSON(w, http.StatusOK, users)

	case "POST":
		// Create a new user value.
		u := user{
			Name:  r.PostFormValue("name"),
			Email: r.PostFormValue("email"),
			Phone: r.PostFormValue("phone"),
		}

		// Add the new user to the in-memory datastore.
		app.mu.Lock()
		{
			app.users = append(app.users, u)
		}
		app.mu.Unlock()

		// Redirect the client to the home page.
		http.Redirect(w, r, "/", http.StatusSeeOther)

	default:
		// We don't support these verbs. Let the user know.
		status := http.StatusMethodNotAllowed
		w.Header().Set("Allow", "GET, HEAD, POST")
		http.Error(w, http.StatusText(status), status)
	}
}

// searchUsers handles the /api/v1/search path.
func searchUsers(w http.ResponseWriter, r *http.Request) {
	// Extract the search term from the query string.
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Query is required", http.StatusBadRequest)
		return
	}

	// Iterate over the slice looking for the user.
	var foundUser *user
	app.mu.RLock()
	{
		for _, u := range app.users {
			if strings.Contains(u.Name, query) {
				foundUser = &u
				break
			}
		}
	}
	app.mu.RUnlock()

	// Did we find the user.
	if foundUser != nil {
		respondJSON(w, http.StatusOK, *foundUser)
		return
	}

	// Respond the user was not found.
	http.NotFound(w, r)
}

func main() {
	// This will handle all paths without specific routes.
	http.HandleFunc("/", baseHandler)

	// Create a ServeMux and add some routes.
	api := http.NewServeMux()
	api.HandleFunc("/users", usersHandler)
	api.HandleFunc("/search", searchUsers)

	// We are stripping out the prefix (/api/v1/) before the
	// route will be handled. This allows the routes above
	// to be compliant with this prefix.
	http.Handle("/api/v1/", http.StripPrefix("/api/v1", api))

	// Start the service.
	bind := ":4000"
	log.Println("Serving HTTP on", bind)
	log.Fatalln(http.ListenAndServe(bind, nil))
}
