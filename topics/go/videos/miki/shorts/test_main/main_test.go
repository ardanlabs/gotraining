// Use TestMain for package level setup/teardown (fixtures).

package main

import (
	"flag"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()
	setup()

	v := m.Run()

	teardown()
	os.Exit(v)
}

func TestNothing(t *testing.T) {
	t.Log("setup")
	defer t.Log("teardown")

	// TODO
}

func setup() {
	log.Printf("SETUP")
}

func teardown() {
	log.Printf("TEARDOWN")
}
