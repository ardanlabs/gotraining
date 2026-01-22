package main

import (
	"fmt"
	"regexp"
)

var (
	// 2023-11-23T10:42:42Z WARNING read of /etc/passwd
	re    = `(?P<time>[^ ]+) (?P<level>[A-Z]+) (?P<message>.*)`
	logRe = regexp.MustCompile(re)
)

func main() {
	logLine := "2023-11-23T10:42:42Z WARNING read of /etc/passwd"
	matches := logRe.FindStringSubmatch(logLine)
	if len(matches) == 0 {
		fmt.Println("no matches")
		return
	}

	fmt.Println("time:", matches[logRe.SubexpIndex("time")])
}
