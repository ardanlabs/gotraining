package strings_test

import (
	strings "github.com/ardanlabs/gotraining/topics/go/algorithms/strings/reverse"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestReverseString(t *testing.T) {

	revTests := []struct {
		name   string
		input  string
		output string
	}{
		{"basic", "Hello World", "dlroW olleH"},
	}

	for _, tt := range revTests {
		actual := strings.ReverseString(tt.input)
		if actual != tt.output {
			t.Logf("\t%s\tShould be able to reverse string: %s\n.", failed, tt.input)
			t.Fatalf("\t\tGot %s, Expected %s.", actual, tt.output)
		}
		t.Logf("\t%s\tShould be able to reverse string.", succeed)
	}
}
