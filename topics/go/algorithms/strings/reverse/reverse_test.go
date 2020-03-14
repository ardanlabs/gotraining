package strings_test

import (
	"github.com/ardanlabs/gotraining/topics/go/algorithms/strings/reverse"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestReverseString(t *testing.T) {

	// Create string for testing.
	str := "Hello World"

	// Expected value of reverse string.
	exp := "dlroW olleH"

	res := strings.ReverseString(str)

	// Test if result equals expected value.
	if res != exp {
		t.Logf("\t%s\tShould be able to reverse string.", failed)
		t.Fatalf("\t\tGot %s, Expected %s.", res, exp)
	}
	t.Logf("\t%s\tShould be able to reverse string.", succeed)
}
