// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Current Status Codes:
//		200 OK           : StatusOK                  : Call is success and returning data.
//		204 No Content   : StatusNoContent           : Call is success and returns no data.
//		400 Bad Request  : StatusBadRequest          : Invalid post data (syntax or semantics).
//		401 Unauthorized : StatusUnauthorized        : Authentication failure.
//		404 Not Found    : StatusNotFound            : Invalid URL or identifier.
//		500 Internal     : StatusInternalServerError : Application specific beyond scope of user.

package app

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

// Invalid describes a validation error belonging to a specific field.
type Invalid struct {
	Fld string `json:"field_name"`
	Err string `json:"error"`
}

// InvalidError is a custom error type for invalid fields.
type InvalidError []Invalid

// Error implements the error interface for InvalidError.
func (ei InvalidError) Error() string {
	// TODO
	return "TODO" // fmt.Sprintf("%+v", ei)
}

// jsonError is the response for errors that occur within the API.
type jsonError struct {
	Error  string       `json:"error"`
	Fields InvalidError `json:"fields,omitempty"`
}

//==============================================================================

var (
	// ErrNotAuthorized occurs when the call is not authorized.
	ErrNotAuthorized = errors.New("Not authorized")

	// ErrDBNotConfigured occurs when the DB is not initialized.
	ErrDBNotConfigured = errors.New("DB not initialized")

	// ErrNotFound is abstracting the mgo not found error.
	ErrNotFound = errors.New("Entity not found")

	// ErrInvalidID occurs when an ID is not in a valid form.
	ErrInvalidID = errors.New("ID is not in it's proper form")

	// ErrValidation occurs when there are validation errors.
	ErrValidation = errors.New("Validation errors occurred")
)

//==============================================================================

// Error handles all error responses for the API.
func Error(w http.ResponseWriter, traceID string, err error) {

	switch err {
	case ErrNotFound:
		RespondError(w, traceID, err, http.StatusNotFound)
		return

	case ErrInvalidID:
		RespondError(w, traceID, err, http.StatusBadRequest)
		return

	case ErrValidation:
		RespondError(w, traceID, err, http.StatusBadRequest)
		return

	case ErrNotAuthorized:
		RespondError(w, traceID, err, http.StatusUnauthorized)
		return
	}

	switch e := err.(type) {
	case InvalidError:
		v := jsonError{
			Error:  "field validation failure",
			Fields: e,
		}

		Respond(w, traceID, v, http.StatusBadRequest)
		return
	}

	RespondError(w, traceID, err, http.StatusInternalServerError)
}

// RespondError sends JSON describing the error
func RespondError(w http.ResponseWriter, traceID string, err error, code int) {
	Respond(w, traceID, jsonError{Error: err.Error()}, code)
}

// Respond sends JSON to the client.
// If code is StatusNoContent, v is expected to be nil.
func Respond(w http.ResponseWriter, traceID string, data interface{}, code int) {
	log.Printf("%s : api : Respond : Started : Code[%d]\n", traceID, code)

	// Load any user defined header values.
	if app.userHeaders != nil {
		for key, value := range app.userHeaders {
			log.Printf("%s : api : Respond : Setting user headers : %s:%s\n", traceID, key, value)
			w.Header().Set(key, value)
		}
	}

	// Just set the status code and we are done.
	if code == http.StatusNoContent {
		w.WriteHeader(code)
		return
	}

	// Set the content type.
	w.Header().Set("Content-Type", "application/json")

	// Write the status code.
	w.WriteHeader(code)

	// Marshal the data into a JSON string.
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Printf("%s : api : Respond %v Marshalling JSON response\n", traceID, err)
		jsonData = []byte("{}")
	}

	// We can send the result straight through.
	io.WriteString(w, string(jsonData))

	log.Printf("%s : api : Respond : Completed\n", traceID)
}
