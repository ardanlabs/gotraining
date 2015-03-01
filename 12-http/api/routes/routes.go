// Package routes declares paths and bind them to controller methods.
package routes

import (
	"log"
	"net/http"

	"code.google.com/p/go-uuid/uuid"
	"github.com/ArdanStudios/gotraining/12-http/api/app"
	"github.com/ArdanStudios/gotraining/12-http/api/ctrls"
	"github.com/dimfeld/httptreemux"
)

// TM is the treemux router for handling requests.
var TM *httptreemux.TreeMux

func init() {
	TM = httptreemux.New()
	users()
}

// users binds the routes to the handlers for the users service.
func users() {
	addRoute("GET", "/v1/users", ctrls.Users.List)
	addRoute("POST", "/v1/users", ctrls.Users.Create)
	addRoute("GET", "/v1/users/:id", ctrls.Users.Retrieve)
	addRoute("PUT", "/v1/users/:id", ctrls.Users.Update)
	addRoute("DELETE", "/v1/users/:id", ctrls.Users.Delete)
}

// addRoute allows routes to be injected into the middleware with the context.
func addRoute(verb string, path string, userHandler func(c *app.Context)) {
	f := func(w http.ResponseWriter, r *http.Request, p map[string]string) {
		uid := uuid.New()
		log.Printf("%s : routes : handler : Started : Verb[%s] Path[%s] URL[%s]\n", uid, verb, path, r.URL.RequestURI())

		c := app.Context{
			Session:        app.GetSession(),
			ResponseWriter: w,
			Request:        r,
			Params:         p,
			SessionID:      uid,
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

	TM.Handle(verb, path, f)

	log.Printf("routes : addRoute : Added : Path[%s]\n", path)
}
