package mvc

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"code.google.com/p/go-uuid/uuid"
	"github.com/ArdanStudios/gotraining/11-http/api/mongodb"
	"github.com/dimfeld/httptreemux"
	"gopkg.in/mgo.v2"
)

// Context contains data in context with all requests.
type Context struct {
	Session   *mgo.Session
	Writer    http.ResponseWriter
	Request   *http.Request
	Params    map[string]string
	SessionID string
}

// AddRoute allows routes to be injected into the middleware with the context.
func AddRoute(router *httptreemux.TreeMux, path string, userHandler func(c *Context), verb string) {
	f := func(w http.ResponseWriter, r *http.Request, p map[string]string) {
		uid := uuid.New()
		log.Printf("%s : mvc : handler : Started : Path[%s] URL[%s]\n", uid, path, r.URL.RequestURI())

		c := Context{
			Session:   mongodb.GetSession(),
			Writer:    w,
			Request:   r,
			Params:    p,
			SessionID: uid,
		}

		defer func() {
			if r := recover(); r != nil {
				log.Println(uid, ": mvc : handler : PANIC :", r)
			}

			c.Session.Close()
			log.Println(uid, ": mvc : handler : Completed")
		}()

		if err := c.authenticate(); err != nil {
			log.Println(uid, ": mvc : handler : ERROR :", err)
			return
		}

		userHandler(&c)
	}

	router.Handle(verb, path, f)

	log.Printf("main : mvc : AddRoute : Added : Path[%s]\n", path)
}

// authenticate handles the authentication of each request.
func (c *Context) authenticate() error {
	log.Println(c.SessionID, ": mvc : authenticate : Started")

	// ServeError(w, errors.New("Auth Error"), http.StatusUnauthorized)

	log.Println(c.SessionID, ": mvc : authenticate : Completed")
	return nil
}

// ServeError handles application errors
func (c *Context) ServeError(err error, statusCode int) {
	log.Printf("%s : mvc : ServeError : Started : Status[%d]\n", c.SessionID, statusCode)

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

	s := string(data)
	log.Printf("%s : mvc : ServeError : Response\n%s\n", c.SessionID, s)

	http.Error(c.Writer, s, statusCode)
	log.Println(c.SessionID, ": mvc : ServeError : Completed")
}

// ServeJSON handles serving values as JSON.
func (c *Context) ServeJSON(v interface{}) {
	log.Printf("%s : mvc : ServeJSON : Started\n", c.SessionID)

	data, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		http.Error(c.Writer, "", http.StatusInternalServerError)
		return
	}

	s := string(data)
	log.Printf("%s : mvc : ServeError : Response\n%s\n", c.SessionID, s)

	fmt.Fprintf(c.Writer, s)
	log.Println(c.SessionID, ": mvc : ServeJSON : Completed")
}
