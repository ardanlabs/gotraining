package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "To be or not to be"
	freqs := make(map[string]int) // word -> count
	for _, word := range strings.Fields(text) {
		word := strings.ToLower(word)
		freqs[word]++
	}

	for word, freq := range freqs {
		fmt.Printf("%s → %d\n", word, freq)
	}
}
