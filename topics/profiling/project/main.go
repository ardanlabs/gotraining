// This program provides a sample web service that uses concurrency
// and channels to perform a coordinated set of asynchronous searches.
package main

import (
	"log"
	_ "net/http/pprof"
	"os"

	"github.com/ardanlabs/gotraining/topics/profiling/project/service"
)

// init is called before main. We are using init to
// set the logging package.
func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	log.SetOutput(os.Stdout)
}

// main is the entry point for the application.
func main() {
	service.Run()
}
