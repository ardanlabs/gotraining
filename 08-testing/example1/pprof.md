// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// First we can `go get`
// go get github.com/davecheney/profile
// graphviz Graph Visualization Tools (http://www.graphviz.org/Download.php)
// http://ghostscript.com/download/gsdnld.html
//    ./configure
//    make
//    make install
// go tool pprof --pdf ./example1 cpu.pprof > callgraph.pdf

// main is the entry point for the application.
func main() {
	cfg := profile.Config{
		MemProfile:     true,
		CPUProfile:     true,
		ProfilePath:    ".",  // store profiles in current directory
		NoShutdownHook: true, // do not hook SIGINT
	}

	// p.Stop() must be called before the program exits to
	// ensure profiling information is written to disk.
	p := profile.Start(&cfg)
	defer p.Stop()

	. . .
}
