package base

import (
	"io"
	"log"
	"os"
)

// Logger is the default logger for the entire golearn package. It writes
// to stdout and has no prefix and no flags.
var Logger *log.Logger = log.New(os.Stdout, "", 0)

// SetLogger sets the base logger for the entire golearn package.
func SetLogger(logger *log.Logger) {
	Logger = logger
}

// SetLoggerOut creates a new base logger for the entire golearn
// package using the given out instead of the default, os.Stdout.
// The other log options are set to the default, i.e. no prefix and no
// flags.
func SetLoggerOut(out io.Writer) {
	Logger = log.New(out, "", 0)
}

// Silent turns off logging throughout the golearn package by setting
// the logger to write to dev/null.
func Silent() {
	if out, err := os.Open(os.DevNull); err != nil {
		panic(err)
	} else {
		Logger = log.New(out, "", 0)
	}
}
