package main

import (
	"fmt"
	"reflect"
)

// =============================================================================

// Defining two types that implement the fmt.Stringer interface. Each
// implementation creates a stringified version of the concrete type.

type user struct {
	name  string
	email string
}

func (u user) String() string {
	return fmt.Sprintf("{type: \"user\", name: %q, email: %q}", u.name, u.email)
}

type customer struct {
	name  string
	email string
}

func (u customer) String() string {
	return fmt.Sprintf("{type: \"customer\", name: %q, email: %q}", u.name, u.email)
}

// =============================================================================

// These functions implement a stringify function that is specific to each of
// the concrete types implemented above. In each case, the stringify function
// returns a slice of strings. These function use the String method against
// the individual user or customer value.

func stringifyUsers(users []user) []string {
	ret := make([]string, 0, len(users))
	for _, user := range users {
		ret = append(ret, user.String())
	}
	return ret
}

func stringifyCustomers(customers []customer) []string {
	ret := make([]string, 0, len(customers))
	for _, customer := range customers {
		ret = append(ret, customer.String())
	}
	return ret
}

// =============================================================================

// This function provides an empty interface solution which uses type assertions
// for the different concrete slices to be supported. We've basically moved the
// functions from above into case statements. This function uses the String
// method against the value.

func stringifyAssert(v interface{}) []string {
	switch list := v.(type) {
	case []user:
		ret := make([]string, 0, len(list))
		for _, value := range list {
			ret = append(ret, value.String())
		}
		return ret
	case []customer:
		ret := make([]string, 0, len(list))
		for _, value := range list {
			ret = append(ret, value.String())
		}
		return ret
	}
	return nil
}

// =============================================================================

// This function provides a reflection solution which allows a slice of any
// type to be provided and stringified. This is a generic function thanks to the
// reflect package. Notice the call to the String method via relfection.

func stringifyReflect(v interface{}) []string {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Slice {
		return nil
	}
	ret := make([]string, 0, val.Len())
	for i := 0; i < val.Len(); i++ {
		m := val.Index(i).MethodByName("String")
		if !m.IsValid() {
			return nil
		}
		data := m.Call(nil)
		ret = append(ret, data[0].String())
	}
	return ret
}

// =============================================================================

// This function provides a generics solution which allows a slice of some type
// T (to be determined later) to be passed and stringified. This code more
// closely resembles the concrete implementations that we started with and is
// easier to read than the reflect implementation. However, an interface
// constraint of type fmt.Stringer is applied to allow the compiler to know the
// value of type T passed requires a String method.

func stringify[T fmt.Stringer](slice []T) []string {
    ret := make([]string, 0, len(slice))
    for _, value := range slice {
        ret = append(ret, value.String())
    }
    return ret
}

// =============================================================================

func main() {
	users := []user{
		{name: "Bill", email: "bill@ardanlabs.com"},
		{name: "Ale", email: "ale@whatever.com"},
	}
	s1 := stringifyUsers(users)
	s2 := stringifyAssert(users)
	s3 := stringifyReflect(users)
	s4 := stringify(users)

	customers := []customer{
		{name: "Google", email: "you@google.com"},
		{name: "MSFT", email: "you@msft.com"},
	}
	s5 := stringifyCustomers(customers)
	s6 := stringifyAssert(customers)
	s7 := stringifyReflect(customers)
	s8 := stringify(customers)

	fmt.Println("users Con:", s1)
	fmt.Println("users Int:", s2)
	fmt.Println("users Ref:", s3)
	fmt.Println("users Gen:", s4)
	fmt.Println("cust Con:", s5)
	fmt.Println("cust Int:", s6)
	fmt.Println("cust Ref:", s7)
	fmt.Println("cust Gen:", s8)
}