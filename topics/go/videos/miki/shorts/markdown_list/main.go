package main

import (
	"bytes"
	"fmt"
)

func MarkdownList(items []string) string {
	var buf bytes.Buffer
	for _, item := range items {
		fmt.Fprintf(&buf, "- %s\n", item)
	}
	return buf.String()
}

func main() {
	items := []string{
		"Wake up",
		"Feed cat",
		"Coffee",
	}
	out := MarkdownList(items)
	fmt.Println(out)
}
