// Example shows how to iterate over a slice.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Dog struct {
	Name string
	Age  int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	dogs, _ := GetDogs()

	for _, dog := range dogs {
		fmt.Printf("Addr[%p] Name[%s]\tAge[%d]\n", &dog, dog.Name, dog.Age)
	}
}

func GetDogs() ([]Dog, error) {
	names := []string{"bill", "jack", "sammy", "jill", "choley", "harley", "jamie", "Ed", "Lisa", "Missy"}

	var dogs []Dog
	for dog := 0; dog < 25; dog++ {
		name := rand.Intn(9)
		age := rand.Intn(20)
		dogs = append(dogs, Dog{names[name], age})
	}

	return dogs, nil
}
