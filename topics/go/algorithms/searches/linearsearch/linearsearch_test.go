package linearsearch

import (
	"math/rand"
	"testing"
	"time"
)

func TestLinearSearch(t *testing.T) {
	l, p := generateRandomList(99)
	data := struct {
		list []int
		pick int
	}{
		list: l,
		pick: p,
	}

	t.Run("Linear Search Iterative", func(t *testing.T) {
		t.Log("Search to find the target value")
		{
			result := linearSearchIterative(data.list, data.pick)

			if result == -1 {
				t.Errorf("excepted %d, but got -1", result)
			}

			result2 := linearSearchIterative(data.list, -10)
			if result2 != -1 {
				t.Errorf("excepted -1 , but got %d", result2)
			}

		}
	})

	t.Run("Linear Search Recursive", func(t *testing.T) {
		t.Log("Search to find the target value")
		{
			result := linearSearchRecursive(data.list, data.pick, 0)

			if result == -1 {
				t.Errorf("excepted %d, but got -1", result)
			}

			result2 := linearSearchRecursive(data.list, -10, 0)
			if result2 != -1 {
				t.Errorf("excepted -1 , but got %d", result2)
			}
		}
	})

	t.Run("Double Linear Search Iterative", func(t *testing.T) {
		t.Log("Search to find the target value in both side")
		{
			result := doubleLinearSearchIterative(data.list, data.pick)

			if result == -1 {
				t.Errorf("excepted %d, but got -1", result)
			}

			result2 := doubleLinearSearchIterative(data.list, -10)
			if result2 != -1 {
				t.Errorf("excepted -1 , but got %d", result2)
			}

		}
	})

	t.Run("Double Linear Search Recursive", func(t *testing.T) {
		t.Log("Search to find the target value in both side")
		{
			result := doubleLinearSearchRecursive(data.list, data.pick, 0, len(data.list)-1)

			if result == -1 {
				t.Errorf("excepted %d, but got -1", result)
			}

			result2 := doubleLinearSearchRecursive(data.list, -10, 0, len(data.list)-1)
			if result2 != -1 {
				t.Errorf("excepted -1 , but got %d", result2)
			}
		}
	})

}

// generateRandomList it generates a random list of numbers
func generateRandomList(value int) (list []int, pick int) {

	rand.Seed(time.Now().Unix())

	// generate the array numbers
	list = rand.Perm(value)

	//  random index
	random := rand.Intn(value)

	return list, list[random]
}
