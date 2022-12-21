//go:build profile

package main

// Will export /debug/pprof/
import _ "net/http/pprof"
