// This program provides a sample web service that uses concurrency
// and channels to perform a coordinated set of asynchronous searches.
package main

import (
	"expvar"
	"log"
	_ "net/http/pprof"
	"os"
	"runtime"
	"time"

	"github.com/ardanlabs/gotraining/topics/profiling/project/service"
)

// init is called before main. We are using init to
// set the logging package.
func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	log.SetOutput(os.Stdout)
}

// expvars is adding the goroutine counts to the variable set.
func expvars() {

	// Add goroutine counts to the variable set.
	gr := expvar.NewInt("goroutines")
	go func() {
		for _ = range time.Tick(time.Millisecond * 250) {
			gr.Set(int64(runtime.NumGoroutine()))
		}
	}()
}

// main is the entry point for the application.
func main() {
	expvars()
	service.Run()
}
