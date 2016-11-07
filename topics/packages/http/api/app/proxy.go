// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package app

import "net/http"

// ProxyResponseWriter records the status code written by a call to the
// WriteHeader function on a http.ResponseWriter interface. This type also
// implements the http.ResponseWriter interface.
type ProxyResponseWriter struct {
	Status int
	http.ResponseWriter
}

// Header implements the http.ResponseWriter interface and simply relays the
// request.
func (prw *ProxyResponseWriter) Header() http.Header {
	return prw.ResponseWriter.Header()
}

// Write implements the http.ResponseWriter interface and simply relays the
// request.
func (prw *ProxyResponseWriter) Write(data []byte) (int, error) {
	return prw.ResponseWriter.Write(data)
}

// WriteHeader implements the http.ResponseWriter interface and simply relays
// the request and records the status code written.
func (prw *ProxyResponseWriter) WriteHeader(status int) {
	prw.ResponseWriter.WriteHeader(status)
	prw.Status = status
}
