package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "65 in favour, 43 against and 21 abstentions."

	re := regexp.MustCompile(`(\d+).*(\d+).*(\d+)`)
	matches := re.FindStringSubmatch(text)
	fmt.Println(matches[1:]) // 0 is whole match
	// [65 2 1]

	re = regexp.MustCompile(`(\d+).*?(\d+).*?(\d+)`)
	matches = re.FindStringSubmatch(text)
	fmt.Println(matches[1:])
	// [65 43 21]

	re = regexp.MustCompile(`(\d+)(?U:.*)(\d+)(?U:.*)(\d+)`)
	matches = re.FindStringSubmatch(text)
	fmt.Println(matches[1:])
	// [65 43 21]
}
