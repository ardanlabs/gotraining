// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a program that is given a list of file names as arguments then prints
// the sha256 sum for the contents of each file. Print the hashes as a hex string.
package main

func main() {

	// Make a hash value from crypto/sha256.

	// Make a variable to count failures

	// Loop through all of os.Args skipping the first value.
	{

		// Reset the hash value before each use.

		// Skip this argument if it is a directory.

		// Attempt to open the file in question using os.Open.

		// Ensure the file is closed when we're done processing it

		// Write the file to the hash so we can calculate it.
		// Tip: Your hash value is an io.Writer and the file value is an io.Reader.
		// The io.Copy function works with both.

		// Print the sha256 sum in hex format followed by the name of the file.
		// You can use the %x directive of fmt.Printf or use encoding/hex.
	}

	// If at least one failure was encountered then exit with status code 1.
}
