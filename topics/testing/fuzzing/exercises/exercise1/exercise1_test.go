package fuzzprot_test

import (
	"reflect"
	"testing"

	fuzzprot "github.com/ardanlabs/gotraining/topics/fuzzing/exercises/exercise1"
)

const succeed = "\u2713"
const failed = "\u2717"

// TestUnpackUsers tests the unpacking routine.
func TestUnpackUsers(t *testing.T) {
	tests := []struct {
		input []byte
		users []fuzzprot.User
	}{
		{[]byte("\x01\x0346\x01\x03ADM\x02\x04Bill\x00"), []fuzzprot.User{{Age: 46, Type: "ADM", Name: "Bill"}}},
	}

	t.Log("Given the need to test the UnpackUsers function.")
	{
		for i, tt := range tests {
			t.Logf("\tTest %d:\tWhen checking input: %v", i, tt.input)
			{
				if u, _ := fuzzprot.UnpackUsers(tt.input); !reflect.DeepEqual(tt.users, u) {
					t.Log("GOT:", u)
					t.Log("EXP:", tt.users)
					t.Fatalf("\t%s\tShould get the user.", failed)
				}
				t.Logf("\t%s\tShould get the user.", succeed)
			}
		}
	}
}
