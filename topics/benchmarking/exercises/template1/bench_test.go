// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

/*
http://golang.org/cmd/go/#hdr-Description_of_testing_flags
go test -run=XXX -bench=.
go test -run=XXX -bench=BenchmarkSprintf
go test -run=XXX -bench=BenchmarkFormat
go test -run=XXX -bench=BenchmarkAtoi
*/

// Write three benchmark tests for converting an integer into a string. First using the
// fmt.Sprintf function, then the strconv.FormatInt function and then strconv.Itoa.
// Identify which function performs the best.
package main

// Add imports.

// BenchmarkSprintf provides performance numbers for the fmt.Sprintf function.
// fmt.Sprintf("%d", number)

// BenchmarkFormat provides performance numbers for the strconv.FormatInt function.
// strconv.FormatInt(number, 10)

// BenchmarkItoa provides performance numbers for the strconv.Itoa function.
// strconv.Itoa(number)
