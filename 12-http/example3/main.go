// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

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
	mu     sync.RWMutex
	users  []user
	idxTpl *template.Template
	fs     http.Handler
}{
	users:  make([]user, 0),                                           // Initialized to avoid JSON null.
	idxTpl: template.Must(template.ParseFiles("template/index.html")), // idxTpl maintains access to the page template.
	fs:     http.FileServer(http.Dir("public")),                       // fs is a file serving handler for static files.
}

func main() {
	// This will handle all paths without specific routes.
	http.HandleFunc("/", baseHandler)

	// Create a ServeMux and add some routes.
	api := http.NewServeMux()
	api.HandleFunc("/users", usersHandler)
	api.HandleFunc("/search", searchUsers)

	// TODO: Describe what this is doing.
	http.Handle("/api/v1/", http.StripPrefix("/api/v1", api))

	// Start the service.
	bind := ":4000"
	log.Println("Serving HTTP on", bind)
	log.Fatalln(http.ListenAndServe(bind, nil))
}

// baseHandler handles serving the index template and static assets.
func baseHandler(w http.ResponseWriter, r *http.Request) {
	// If an static asset is being requested, use the file server.
	if r.URL.Path != "/" {
		app.fs.ServeHTTP(w, r)
		return
	}

	// Execute the template and respond with the results.
	app.mu.RLock()
	{
		data := struct{ Users []user }{app.users}
		app.idxTpl.Execute(w, data)
	}
	app.mu.RUnlock()
}

// usersHandler handles the /api/v1/users path.
func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET", "HEAD":
		// Respond the current slice of users.
		app.mu.RLock()
		{
			respondJSON(w, http.StatusOK, app.users)
		}
		app.mu.RUnlock()

	case "POST":
		// Add a new user to the slice.
		u := user{
			Name:  r.PostFormValue("name"),
			Email: r.PostFormValue("email"),
			Phone: r.PostFormValue("phone"),
		}

		app.mu.Lock()
		{
			app.users = append(app.users, u)
		}
		app.mu.Unlock()

		// TODO: Please explain this redirect and status?
		http.Redirect(w, r, "/", http.StatusSeeOther)

	default:
		// TODO: Please explain this code.
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

	// Interate over the slice looking for the user.
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

// respondJSON sends status and writes JSON to the client.
func respondJSON(w http.ResponseWriter, status int, val interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(val)
}
