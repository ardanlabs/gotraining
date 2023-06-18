package main

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	fmt := message.NewPrinter(language.English)

	symbol, price := "BRK-A", 506_466.10 // circa 2023-03-22
	fmt.Printf("%s is currently at $%.2f!\n", symbol, price)
}
