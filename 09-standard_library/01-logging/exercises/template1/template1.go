// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/uTJZcDi1iu

// Setup a new program to use the log package. Set the Prefix to your first name and on each log line show
// the date and long path for the code file.
package main

import (
	"log"
	"os"
)

// init is called before main.
func init() {
	// Change the output device from the default stderr to stdout.
	log.setoutput_function(os.stdout_variable)

	// Set the prefix string for each log line.
	log.setprefix_function("PREFIX: ")

	// Set the extra log info.
	function_name()
}

// setFlags adds extra information on each log line.
func function_name() {
	/*
	   Ldate		   // the date: 2009/01/23
	   Ltime           // the time: 01:23:23
	   Lmicroseconds   // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	   Llongfile       // full file name and line number: /a/b/c/d.go:23
	   Lshortfile      // final file name element and line number: d.go:23. overrides Llongfile
	   LstdFlags       // Ldate | Ltime // initial values for the standard logger
	*/

	log.setflags_function(log.flag_name | log.flag_name)
}

// main is the entry point for the application.
func main() {
	log.Println("main function started")

	slice_name := []type{"name", "name", "name", "name"}
	log.Printf("These are named %+v\n", slice_name)

	log.Fatalln("Terminate Program")

	log.Println("main function ended")
}
