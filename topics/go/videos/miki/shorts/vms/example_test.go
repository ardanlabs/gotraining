package vms_test

import (
	"fmt"
	"shorts/vms"
)

func ExampleClient_Start() {
	req := vms.StartRequest{
		Image: "debian:bookworm-slim",
		Count: 1,
	}
	fmt.Println(req)
	// ...

	// Output:
	// {debian:bookworm-slim 1}
}
