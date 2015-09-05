// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/Qmj_C5PEs1

// Sample program to show how you can use the FileServer handle with
// this service as a single call to ListenAndServe.
package main

import "net/http"

func main() {
	// Since the http.Handler interface is so universal and ubiquitous, you can
	// encapsulate some really neat pieces of functionality in a single
	// http.Handler. You can mount http.Handlers to routers, middleware, filters,
	// and other server constructs.
	//
	// In this example we will create a file server in one line of code
	http.ListenAndServe(":4000", http.FileServer(http.Dir(".")))
}
