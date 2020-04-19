package hash_test

import (
	"testing"

	"github.com/ardanlabs/gotraining/topics/go/algorithms/data/hash"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestHash(t *testing.T) {
	t.Log("Given the need to test hash functionality.")
	{
		testID := 0
		t.Logf("\tTest %d:\tWhen checking basic hashing operations", testID)
		{
			h := hash.New()
			k1, v1 := "key1", 1
			k2, v2 := "key2", 2
			h.Store(k1, v1)
			h.Store(k2, v2)

			if h.Len() != 2 {
				t.Errorf("\t%s\tTest %d:\tShould have the correct number of entries.", failed, testID)
				t.Fatalf("\t\tTest %d:\tGot %q, Expected %q", testID, h.Len(), 2)
			}
			t.Logf("\t%s\tTest %d:\tShould have the correct number of entries.", succeed, testID)

			v, err := h.Retrieve(k1)
			if err != nil {
				t.Fatalf("\t%s\tTest %d:\tShould be able to retrieve a value.", failed, testID)
			}
			t.Logf("\t%s\tTest %d:\tShould be able to retrieve a value.", succeed, testID)

			if v != v1 {
				t.Errorf("\t%s\tTest %d:\tShould have the correct value after retrieve.", failed, testID)
				t.Fatalf("\t\tTest %d:\tGot %q, Expected %q", testID, v, v1)
			}
			t.Logf("\t%s\tTest %d:\tShould have the correct value after retrieve.", succeed, testID)

			v1b := 11
			h.Store(k1, v1b)

			v, err = h.Retrieve(k1)
			if err != nil {
				t.Fatalf("\t%s\tTest %d:\tShould be able to retrieve a value.", failed, testID)
			}
			t.Logf("\t%s\tTest %d:\tShould be able to retrieve a value.", succeed, testID)

			if v != v1b {
				t.Errorf("\t%s\tTest %d:\tShould have the correct value after retrieve.", failed, testID)
				t.Fatalf("\t\tTest %d:\tGot %q, Expected %q", testID, v, v1)
			}
			t.Logf("\t%s\tTest %d:\tShould have the correct value after retrieve.", succeed, testID)

			if err := h.Delete(k1); err != nil {
				t.Fatalf("\t%s\tTest %d:\tShould be able to delete a value.", failed, testID)
			}
			t.Logf("\t%s\tTest %d:\tShould be able to delete a value.", succeed, testID)

			if _, err := h.Retrieve(k1); err == nil {
				t.Fatalf("\t%s\tTest %d:\tShould be able to see the value has been deleted.", failed, testID)
			}
			t.Logf("\t%s\tTest %d:\tShould be able to see the value has been deleted.", succeed, testID)

			k3 := "key3"
			if _, err = h.Retrieve(k3); err == nil {
				t.Fatalf("\t%s\tTest %d:\tShould be able to see the key does not exist.", failed, testID)
			}
			t.Logf("\t%s\tTest %d:\tShould be able to see the key does not exist.", succeed, testID)

			count := 0
			fn := func(key string, value int) bool {
				count++
				return true
			}
			h.Do(fn)
			if count != h.Len() {
				t.Errorf("\t%s\tTest %d:\tShould be able to run Do %d times.", failed, testID, count)
				t.Fatalf("\t\tTest %d:\tGot %q, Expected %q", testID, v, v1)
				t.Fatalf("Do ran %d times, expected %d", count, h.Len())
			}
			t.Logf("\t%s\tTest %d:\tShould be able to run Do %d times.", succeed, testID, count)
		}
	}
}
