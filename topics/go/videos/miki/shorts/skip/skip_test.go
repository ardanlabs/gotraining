package main

import (
	"os"
	"testing"
)

func inCI() bool {
	return os.Getenv("CI") != "" || os.Getenv("BUILD_NUMBER") != ""
}

func TestMigration(t *testing.T) {
	if !inCI() {
		t.Skip("not in CI")
	}

	// TODO
}
