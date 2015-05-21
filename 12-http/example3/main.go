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

// users is a simulated persistent data store
var (
	mu    sync.RWMutex      // mu protects users
	users = make([]user, 0) // initialized to avoid JSON null
)

// index page template
var idxTpl = template.Must(template.ParseFiles("template/index.html"))

// fs is a file serving handler for static files
var fs = http.FileServer(http.Dir("public"))

func main() {
	// This will handle all paths without specific routes
	http.HandleFunc("/", baseHandler)

	// Create a ServeMux and add some routes.
	api := http.NewServeMux()
	api.HandleFunc("/users", usersHandler)
	api.HandleFunc("/search", searchUsers)

	http.Handle("/api/v1/", http.StripPrefix("/api/v1", api))

	// Start the service.
	bind := ":4000"
	log.Println("Serving HTTP on", bind)
	log.Fatalln(http.ListenAndServe(bind, nil))
}

// baseHandler handles serving the index template and static assets.
func baseHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		fs.ServeHTTP(w, r)
		return
	}

	mu.RLock()
	data := struct{ Users []user }{users}
	mu.RUnlock()

	idxTpl.Execute(w, data)
}

// usersHandler handles the /api/v1/users path.
func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET", "HEAD":
		// List the users
		mu.RLock()
		respondJSON(w, http.StatusOK, users)
		mu.RUnlock()

	case "POST":
		u := user{
			Name:  r.PostFormValue("name"),
			Email: r.PostFormValue("email"),
			Phone: r.PostFormValue("phone"),
		}

		mu.Lock()
		users = append(users, u)
		mu.Unlock()

		http.Redirect(w, r, "/", http.StatusSeeOther)

	default:
		status := http.StatusMethodNotAllowed
		w.Header().Set("Allow", "GET, HEAD, POST")
		http.Error(w, http.StatusText(status), status)
	}
}

// searchUsers handles the /api/v1/search path.
func searchUsers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Query is required", http.StatusBadRequest)
		return
	}

	mu.RLock()
	list := users
	mu.RUnlock()

	for _, u := range list {
		if strings.Contains(u.Name, query) {
			respondJSON(w, http.StatusOK, u)
			return
		}
	}

	http.NotFound(w, r)
}

// respondJSON sends status and writes JSON to the client.
func respondJSON(w http.ResponseWriter, status int, val interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(val)
}
