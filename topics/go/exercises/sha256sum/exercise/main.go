// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a program that is given a list of file names as arguments then prints
// the sha256 sum for the contents of each file. Print the hashes as a hex string.
package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	// Make a hash value from crypto/sha256.
	hash := sha256.New()

	// Keep track of how many failures we encounter.
	var failures int

	// Loop through all of os.Args skipping the first value.
	for _, name := range os.Args[1:] {

		// Attempt to open the file in question using os.Open.
		f, err := os.Open(name)
		if err != nil {
			log.Printf("could not open file %s: %v", name, err)
			failures++
			continue
		}

		// Call the Stat method so we can see if the named argument is a directory.
		stat, err := f.Stat()
		if err != nil {
			log.Printf("could not stat file %s: %v", name, err)
			failures++
			continue
		}

		// Skip directories.
		if stat.IsDir() {
			continue
		}

		// Reset the hash value before each use.
		hash.Reset()

		// Write the file to the hash so we can calculate it.
		// Tip: Your hash value is an io.Writer and the file value is an io.Reader.
		// The io.Copy function works with both.
		if _, err := io.Copy(hash, f); err != nil {
			log.Printf("could not hash file %s: %v", name, err)
			failures++
			continue
		}

		// Print the sha256 sum in hex format followed by the name of the file.
		// You can use the %x directive of fmt.Printf or use encoding/hex.
		fmt.Printf("%x  %s\n", hash.Sum(nil), name)
	}

	// If at least one failure was encountered then exit with status code 1.
	if failures > 0 {
		os.Exit(1)
	}
}
