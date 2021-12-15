package main

import "fmt"

// =============================================================================

// This interface defines a contraint not based on behavior but based on the
// type of data that is acceptable. This type of constrain is important when
// functions (like Add) need to perform operations (like +) that are not 
// supported by all types.

type addOnly interface {
    string | int | int8 | int16 | int32 | int64 | float64
}

func Add[T addOnly](v1 T, v2 T) T {
    return v1 + v2
}

// =============================================================================

// The index function declares that the value of type T must be compliant with
// the new predeclared type constraint "comparable". A type parameter with the
// comparable constraint accepts as a type argument any comparable type. It
// permits the use of == and != with values of that type parameter.

func index[T comparable](list []T, find T) int {
    for i, v := range list {
        if v == find {
            return i
        }
    }
    return -1
}

// =============================================================================

type person struct{
    name  string
    email string
}

func (p person) match(v person) bool {
    return p.name == v.name
}

type food struct{
    name     string
    category string
}

func (f food) match(v food) bool {
    return f.name == v.name
}

// =============================================================================

// The matcher interface defines two constraints. First, it constrains the data
// to what type is acceptable. Second, it constrains the behavior of the data. 
// The match method requires that a value of type T (to be determined later)
// will be the input of the method.
//
// The match function declares that the value of type T must implement the
// matcher interface and is used for the slice and value arguments to the
// function.
//
// Note: The type list inside the interface is not needed for match to work.
//       I'm trying to show how the type list and behavior can be combined.

type matcher[T any] interface {
    person | food
    match(v T) bool
}

func match[T matcher[T]](list []T, find T) int {
    for i, v := range list {
        if v.match(find) {
            return i
        }
    }
    return -1
}

// =============================================================================

func main() {
    fmt.Println(Add(10, 20))
    fmt.Println(Add("A", "B"))
    fmt.Println(Add(3.14159, 2.96))

    durations := []int{5000, 10, 40}
    findDur := 10
    i := index(durations, findDur)
    fmt.Printf("Index: %d for %d\n", i, findDur)

    people := []person{
        {name:"bill", email:"bill@email.com"},
        {name:"jill", email:"jill@email.com"},
        {name:"tony", email:"tony@email.com"},
    }
    findPerson := person{name:"tony"}

    i = index(people, findPerson)
    fmt.Printf("Index: %d for %s\n", i, findPerson.name)

    i = match(people, findPerson)
    fmt.Printf("Match: %d for %s\n", i, findPerson.name)

    foods := []food{
        {name:"apple", category: "fruit"},
        {name:"carrot", category: "veg"},
        {name:"chicken", category: "meat"},
    }
    findFood := food{name:"apple"}

    i = match(foods, findFood)
    fmt.Printf("Match: %d for %s\n", i, findFood.name)
}
