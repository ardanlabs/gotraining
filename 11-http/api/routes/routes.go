// Package routes declare paths and bind them to controller methods.
package routes

import (
	"log"
	"net/http"

	"code.google.com/p/go-uuid/uuid"
	"github.com/ArdanStudios/gotraining/11-http/api/app"
	"github.com/ArdanStudios/gotraining/11-http/api/ctrls"
	"github.com/dimfeld/httptreemux"
)

// Users binds the routes to the handlers for the users service.
func Users(r *httptreemux.TreeMux) {
	addRoute(r, "GET", "/v1/users", ctrls.UsersList)
	addRoute(r, "POST", "/v1/users", ctrls.UsersCreate)
	addRoute(r, "GET", "/v1/users/:id", ctrls.UsersRetrieve)
	addRoute(r, "PUT", "/v1/users/:id", ctrls.UsersUpdate)
	addRoute(r, "DELETE", "/v1/users/:id", ctrls.UsersDelete)
}

// addRoute allows routes to be injected into the middleware with the context.
func addRoute(router *httptreemux.TreeMux, verb string, path string, userHandler func(c *app.Context)) {
	f := func(w http.ResponseWriter, r *http.Request, p map[string]string) {
		uid := uuid.New()
		log.Printf("%s : routes : handler : Started : Verb[%s] Path[%s] URL[%s]\n", uid, verb, path, r.URL.RequestURI())

		c := app.Context{
			Session:   app.GetSession(),
			Writer:    w,
			Request:   r,
			Params:    p,
			SessionID: uid,
		}

		defer func() {
			if r := recover(); r != nil {
				log.Println(uid, ": routes : handler : PANIC :", r)
			}

			c.Session.Close()
			log.Println(uid, ": routes : handler : Completed")
		}()

		if err := c.Authenticate(); err != nil {
			log.Println(uid, ": routes : handler : ERROR :", err)
			return
		}

		userHandler(&c)
	}

	router.Handle(verb, path, f)

	log.Printf("routes : addRoute : Added : Path[%s]\n", path)
}
