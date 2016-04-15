// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Download any document from the web and display the content in
// the terminal and write it to a file at the same time.
package main

// Add imports.

func main() {

	// Download the RSS feed for "http://www.goinggo.net/feeds/posts/default".
	// Check for errors.

	// Arrange for the response Body to be Closed using defer.

	// Declare a slice of io.Writer interface values.

	// Append stdout to the slice of writers.

	// Open a file named "goinggo.rss" and check for errors.

	// Close the file when the function returns.

	// Append the file to the slice of writers.

	// Create a MultiWriter interface value from the writers
	// inside the slice of io.Writer values.

	// Write the response to both the stdout and file using the
	// MultiWriter. Check for errors.
}
