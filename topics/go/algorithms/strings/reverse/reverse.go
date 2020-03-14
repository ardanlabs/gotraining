package main

import (
	"fmt"
)

func main() {

	fmt.Println(ReverseString("Hello World"))
}

func ReverseString(str string) string {

	// Here we create a slice of codePoints.
	codePoints := []rune(str)

	// Here we create int that will be a pointer to the front of the codePoints.
	var fontCodePoint int

	// Here  we create int that will be a pointer to the end of the codePoints.
	backCodePoint := len(codePoints) - 1

	// While fontCodePoint is less than backCodePoint.
	for fontCodePoint < backCodePoint {

		// Here we swap the values of the slice at the codePoints[fontCodePoint] and codePoints[backCodePoint].
		codePoints[fontCodePoint], codePoints[backCodePoint] = codePoints[backCodePoint], codePoints[fontCodePoint]

		// Here we increase the value of fontCodePoint by 1.
		fontCodePoint = fontCodePoint + 1

		// Here we decrease the value of backCodePoint by 1.
		backCodePoint = backCodePoint - 1
	}

	// Here we return a new string in reverse order.
	return string(codePoints)
}

