package main

import (
	"flag"
	"fmt"
	"time"
)

const dateFmt = "2006-01-02"

type DateFlag struct {
	Date *time.Time
}

func (f DateFlag) String() string {
	if f.Date == nil {
		return ""
	}

	return f.Date.Format(dateFmt)
}

func (f DateFlag) Set(s string) error {
	t, err := time.Parse(dateFmt, s)
	if err != nil {
		return err
	}

	*f.Date = t
	return nil
}

func main() {
	reportDate := time.Now()
	flag.Var(DateFlag{&reportDate}, "date", "report date (YYYY-MM-DD)")
	flag.Parse()
	fmt.Println(reportDate.Date())
}
