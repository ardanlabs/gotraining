// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

/*
func New(out io.Writer, prefix string, flag int) *Logger

out:    The out variable sets the destination to which log data will be written.
prefix: The prefix appears at the beginning of each generated log line.
flags:  The flag argument defines the logging properties.
*/

// Sample program to show how to extend the log package
// from the standard library.
package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	// Trace is for full detailed messages.
	Trace *log.Logger

	// Info is for important messages.
	Info *log.Logger

	// Warning is for need to know issue messages.
	Warning *log.Logger

	// Error is for error messages.
	Error *log.Logger
)

// initLog sets the devices for each log type.
func initLog(traceHandle io.Writer, infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer) {
	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {

	// Open a file for warnings.
	warnings, err := os.OpenFile("warnings.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open warning log file")
	}
	defer warnings.Close()

	// Open a file for errors.
	errors, err := os.OpenFile("errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open errors log file")
	}
	defer errors.Close()

	// Create a multi writer for errors.
	multi := io.MultiWriter(errors, os.Stderr)

	// Init the log package for each message type.
	initLog(ioutil.Discard, os.Stdout, warnings, multi)

	// Test each log type.
	Trace.Println("I have something standard to say.")
	Info.Println("Important Information.")
	Warning.Println("There is something you need to know about.")
	Error.Println("Something has failed.")
}
