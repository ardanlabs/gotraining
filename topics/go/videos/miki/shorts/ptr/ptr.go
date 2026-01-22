package main

import (
	"bytes"
	"fmt"
)

type ListRequest struct {
	Bucket     string
	Path       string
	MaxResults *int
}

func (l ListRequest) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "bucket:%q, path:%q, max_results:", l.Bucket, l.Path)
	if l.MaxResults == nil {
		fmt.Fprint(&buf, "<nil>")
	} else {
		fmt.Fprintf(&buf, "%d", *l.MaxResults)
	}

	return buf.String()
}

func main() {
	r := ListRequest{
		Bucket: "ardanlabs",
		Path:   "/videos",
	}
	fmt.Println(r)

}
