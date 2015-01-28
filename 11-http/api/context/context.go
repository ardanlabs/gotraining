package context

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sqs/mux"
	"gopkg.in/mgo.v2"
)

// Context contains data in context with all requests.
type Context struct {
	Session *mgo.Session
	Writer  http.ResponseWriter
	Request *http.Request
}

// AddRoute allows routes to be injected into the middleware with the context.
func AddRoute(router *mux.Router, path string, userHandler func(c *Context)) {
	f := func(w http.ResponseWriter, r *http.Request) {
		c := Context{
			Writer:  w,
			Request: r,
		}

		if err := c.authenticate(); err != nil {
			return
		}

		c.before()
		userHandler(&c)
		c.after()
	}

	router.HandleFunc(path, f)
}

// authenticate handles the authentication of each request.
func (c *Context) authenticate() error {
	log.Printf("controllers : authenticate : Started : Route[%s]\n", c.Request.URL.RequestURI())

	// ServeError(w, errors.New("Auth Error"), http.StatusUnauthorized)

	log.Println("controllers : authenticate : Completed")
	return nil
}

// before handles the setup of processing the request.
func (c *Context) before() {
	log.Printf("controllers : before : Started")

	log.Println("controllers : before : Completed")
}

// after handles the setup of processing the request.
func (c *Context) after() {
	log.Printf("controllers : after : Started")

	log.Println("controllers : after : Completed")
}

// ServeError handles application errors
func (c *Context) ServeError(err error, statusCode int) {
	log.Printf("controllers : ServeError : Started : Error[%s]\n", err)

	e := struct {
		Err string
	}{
		Err: err.Error(),
	}

	data, err := json.MarshalIndent(&e, "", "    ")
	if err != nil {
		http.Error(c.Writer, "", http.StatusInternalServerError)
		return
	}

	http.Error(c.Writer, string(data), statusCode)
	log.Println("controllers : ServeError : Completed")
}

// ServeJSON handles serving values as JSON.
func (c *Context) ServeJSON(v interface{}) {
	log.Printf("controllers : ServeJSON : Started : Error[%+v]\n", v)

	data, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		http.Error(c.Writer, "", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(c.Writer, string(data))
	log.Println("controllers : ServeJSON : Completed")
}
