// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a program that is given a list of file names as arguments then prints
// the sha256 sum for the contents of each file. Print the hashes as a hex string.
package main

import (
	"crypto/sha256"
	"fmt"
	"hash"
	"io"
	"log"
	"os"

	"github.com/pkg/errors"
)

func main() {

	// Make a hash value from crypto/sha256.
	h := sha256.New()

	// Make a variable to count failures
	var failures int

	// Loop through all of os.Args skipping the first value.
	for _, arg := range os.Args[1:] {

		// Reset the hash value before each use.
		h.Reset()

		if err := process(arg, h); err != nil {
			failures++
			log.Print(err)
		}
	}

	// If at least one failure was encountered then exit with status code 1.
	if failures > 0 {
		os.Exit(1)
	}
}

func process(arg string, h hash.Hash) error {

	// Skip this argument if it is a directory.
	info, err := os.Stat(arg)
	if err != nil {
		return errors.Wrap(err, "stat file")
	}
	if info.IsDir() {
		return nil
	}

	// Attempt to open the file in question using os.Open.
	f, err := os.Open(arg)
	if err != nil {
		return errors.Wrap(err, "open file")
	}

	// Ensure the file is closed when we're done processing it
	defer f.Close()

	// Write the file to the hash so we can calculate it.
	// Tip: Your hash value is an io.Writer and the file value is an io.Reader.
	// The io.Copy function works with both.
	if _, err := io.Copy(h, f); err != nil {
		return errors.Wrap(err, "copying contents")
	}

	// Print the sha256 sum in hex format followed by the name of the file.
	// You can use the %x directive of fmt.Printf or use encoding/hex.
	sum := h.Sum(nil)

	fmt.Printf("%x %s\n", sum, arg)

	return nil
}
