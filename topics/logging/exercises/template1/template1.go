// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Setup a new program to use the log package. Set the Prefix to your first name and on each log line show
// the date and long path for the code file.
package main

// Add imports.

func init() {

	// Change the output device from the default stderr to stdout.

	// Set the prefix string for each log line.

	// Set the extra log info.
}

// setFlags adds extra information on each log line.
func setFlags() {
	/*
	   Ldate		   // the date: 2009/01/23
	   Ltime           // the time: 01:23:23
	   Lmicroseconds   // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	   Llongfile       // full file name and line number: /a/b/c/d.go:23
	   Lshortfile      // final file name element and line number: d.go:23. overrides Llongfile
	   LstdFlags       // Ldate | Ltime // initial values for the standard logger
	*/

	// Set the flags.
}

func main() {

	// Use the Println function.

	// Create a slice of strings and initialize with names.
	// Log the values of the slice.

	// Use the Fatalln function.
}
