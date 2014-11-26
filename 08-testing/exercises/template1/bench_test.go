// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

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

import (
	"fmt"
	"strconv"
	"testing"
)

// BenchmarkSprintf provides performance numbers for the fmt.Sprintf function
func function_name(parameter_name [operator]testing.type) {
	variable_name := N

	for i := 0; i < parameter_name.N; i++ {
		fmt.Sprintf("%d", variable_name)
	}
}

// BenchmarkFormat provides performance numbers for the strconv.FormatInt function
func function_name(parameter_name [operator]testing.type) {
	variable_name := int64(N)

	for i := 0; i < parameter_name.N; i++ {
		strconv.FormatInt(variable_name, N)
	}
}

// BenchmarkAtoi provides performance numbers for the strconv.Atoi function
func function_name(parameter_name [operator]testing.type) {
	variable_name := N

	for i := 0; i < parameter_name.N; i++ {
		strconv.Itoa(variable_name)
	}
}
