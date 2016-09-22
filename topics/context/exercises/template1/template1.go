// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Use the template and follow the directions. You will be writing a web handler
// that performs a mock database call but will timeout based on a context if the call
// takes too long. You will also save state into the context.
package main

// Add imports.

// Declare a new type named `key` that is based on an int.

// Declare a constant named `userIPKey` of type `key` set to
// the value of 0.

// Declare a struct type named `User` with two `string` based
// fields named `Name` and `Email`.

func main() {
	routes()

	log.Println("listener : Started : Listening on: http://localhost:4000")
	http.ListenAndServe(":4000", nil)
}

// routes sets the routes for the web service.
func routes() {
	http.HandleFunc("/user", findUser)
}

// Implement the findUser function to leverage the context for
// both timeouts and state.
func findUser(rw http.ResponseWriter, r *http.Request) {

	// Create a context that timesout in fifty milliseconds.

	// Defer the call to cancel.

	// Save the `r.RemoteAddr` value in the context using `userIPKey`
	// as the key. This call returns a new context so replace the
	// current `ctx` value with this new one. The original context is
	// the parent context for this new child context.

	// Create a channel with a buffer size of 1 that works with
	// pointers of type `User` 

	// Use this goroutine to make the database call. Use the channel
	// to get the user back.
	go func() {

		// Get the `r.RemoteAddr` value from the context and log
		// the value you get back.

		// Call the `readDatabase` function provided below and
		// send the returned `User` pointer on the channel.

		// Log that the goroutine is terminating.
	}()

	// Wait for the database call to finish or the timeout.
	select {

	// Add a case to wait on the channel for the `User` pointer.

		// Call the `sendResponse` function provided below to
		// send the `User` to the caller. Use `http.StatusOK`
		// as the status code.

		// Log we sent the response with a StatusOk
		
		return

	// Add a case to wait on the `ctx.Done()` channel.

		// Use this struct value for the error response.
		e := struct{ Error string }{ctx.Err().Error()}

		// Call the `sendResponse` function provided below to
		// send the error to the caller. Use `http.StatusRequestTimeout`
		// as the status code.

		// Log we sent the response with a StatusRequestTimeout

		return
	}
}

// readDatabase performs a pretend database call with
// a second of latency.
func readDatabase() *User {
	u := User{
		Name:  "Bill",
		Email: "bill@ardanlabs.com",
	}

	// Create 100 milliseconds of latency.
	time.Sleep(100 * time.Millisecond)

	return &u
}

// sendResponse marshals the provided value into json and returns
// that back to the caller.
func sendResponse(rw http.ResponseWriter, v interface{}, statusCode int) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)
	json.NewEncoder(rw).Encode(v)
}
