package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	usTz, err := time.LoadLocation("US/Pacific")
	if err != nil {
		log.Fatal(err)
	}

	euTz, err := time.LoadLocation("CET")
	if err != nil {
		log.Fatal(err)
	}

	usTime := time.Date(2023, time.March, 21, 14, 30, 0, 0, usTz)
	euTime := usTime.In(euTz)
	fmt.Println(usTime.Format(time.RFC3339), "->", euTime.Format(time.RFC3339))
}
