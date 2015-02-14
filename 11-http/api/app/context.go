// Package app provides application support for context and MongoDB access.
// Current Status Codes:
// 		200 Successful   : StatusOK                  : Call is success and returning data.
// 		204 No Content   : StatusNoContent           : Call is success and no data being returned.
// 		400 Bad Request  : StatusBadRequest          : Invalid post data.
// 		401 Unauthorized : StatusUnauthorized        : Authentication failure.
// 		404 Not Found    : StatusNotFound            : Invalid URL or identifier.
// 		409 Validation   : StatusConflict            : Validation error on parameters / post data.
// 		500 Internal     : StatusInternalServerError : Application specific beyond scope of user.
package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
)

// Context contains data associated with a single request.
type Context struct {
	Session   *mgo.Session
	Writer    http.ResponseWriter
	Request   *http.Request
	Params    map[string]string
	SessionID string
}

// Invalid is the response for validation errors.
type Invalid struct {
	Fld string `json:"field_name"`
	Err string `json:"error"`
}

// Authenticate handles the authentication of each request.
func (c *Context) Authenticate() error {
	log.Println(c.SessionID, ": api : Authenticate : Started")

	// ServeError(w, errors.New("Auth Error"), http.StatusUnauthorized)

	log.Println(c.SessionID, ": api : Authenticate : Completed")
	return nil
}

// RespondSuccess200 means the call is success and returning data.
func (c *Context) RespondSuccess200(v interface{}) {
	log.Println(c.SessionID, ": api : RespondSuccess200 : Started")

	data, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	s := string(data)
	log.Printf("%s : api : RespondSuccess200 : Response\n%s\n", c.SessionID, s)

	fmt.Fprintf(c.Writer, s)
	log.Println(c.SessionID, ": api : RespondSuccess200 : Completed")
}

// RespondNoContent204 means the call succeeded but no data.
func (c *Context) RespondNoContent204() {
	log.Println(c.SessionID, ": api : RespondNoContent204 : Started")

	http.Error(c.Writer, "", http.StatusNoContent)

	log.Println(c.SessionID, ": api : RespondNoContent204 : Completed")
}

// RespondBadRequest400 means the call contained invalid post data.
func (c *Context) RespondBadRequest400(err error) {
	log.Println(c.SessionID, ": api : RespondBadRequest400 : Started")

	c.respondError(err, http.StatusBadRequest)

	log.Println(c.SessionID, ": api : RespondBadRequest400 : Completed")
}

// RespondUnauthorized401 means the call failed authentication.
func (c *Context) RespondUnauthorized401(err error) {
	log.Println(c.SessionID, ": api : RespondUnauthorized401 : Started")

	c.respondError(err, http.StatusUnauthorized)

	log.Println(c.SessionID, ": api : RespondUnauthorized401 : Completed")
}

// RespondNotFound404 means the call contained an URL or identifier.
func (c *Context) RespondNotFound404() {
	log.Println(c.SessionID, ": api : RespondNotFound404 : Started")

	http.NotFound(c.Writer, c.Request)

	log.Println(c.SessionID, ": api : RespondNotFound404 : Completed")
}

// RespondValidation409 means the call failed validation.
func (c *Context) RespondValidation409(v []Invalid) {
	log.Println(c.SessionID, ": api : RespondValidation409 : Started")

	data, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	s := string(data)
	log.Printf("%s : api : RespondValidation409 : Response\n%s\n", c.SessionID, s)

	http.Error(c.Writer, s, http.StatusConflict)
	log.Println(c.SessionID, ": api : RespondValidation409 : Completed")
}

// RespondInternal500 means the call resulted in an application error.
func (c *Context) RespondInternal500(err error) {
	log.Println(c.SessionID, ": api : RespondInternal500 : Started")

	c.respondError(err, http.StatusInternalServerError)

	log.Println(c.SessionID, ": api : RespondInternal500 : Completed")
}

// respondError handles application errors
func (c *Context) respondError(err error, status int) {
	e := struct {
		Err string
	}{
		Err: err.Error(),
	}

	data, err := json.MarshalIndent(&e, "", "    ")
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	s := string(data)
	log.Printf("%s : api : RespondError%d : Response\n%s\n", c.SessionID, status, s)

	http.Error(c.Writer, s, status)
}
