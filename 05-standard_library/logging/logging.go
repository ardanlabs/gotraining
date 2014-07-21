/*
func New(out io.Writer, prefix string, flag int) *Logger

out:    The out variable sets the destination to which log data will be written.
prefix: The prefix appears at the beginning of each generated log line.
flags:  The flag argument defines the logging properties.

Flags:
const (
// Bits or'ed together to control what's printed. There is no control over the
// order they appear (the order listed here) or the format they present (as
// described in the comments). A colon appears after these items:
// 2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
Ldate = 1 << iota // the date: 2009/01/23
Ltime             // the time: 01:23:23
Lmicroseconds     // microsecond resolution: 01:23:23.123123. assumes Ltime.
Llongfile         // full file name and line number: /a/b/c/d.go:23
Lshortfile        // final file name element and line number: d.go:23. overrides Llongfile
LstdFlags = Ldate | Ltime // initial values for the standard logger
)
*/

// Sample program to show how to use the log package from
// the standard library.
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

// main is the entry point for the application.
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
