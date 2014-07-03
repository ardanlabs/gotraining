package main

import (
	"fmt"
)

func main() {
	s := []int{10}

	p := &s[0]

	for i := 0; i < 1e6; i++ {
		s = append(s, 20)
		if p != &s[0] {
			s[0] = 30
			fmt.Println(p, &s[0], *p, s[0])
			return
		}
	}
}
