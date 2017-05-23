// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Escape Analysis Flaws:
// https://docs.google.com/document/d/1CxgUBPlx9iJzkz9JWkb6tIpTe5q32QDmz8l0BouG0Cw/view

// Sample program to teach the mechanics of escape analysis.
package main

// user represents a user in the system.
type user struct {
	name  string
	email string
}

var u1 user
var u2 *user

// main is the entry point for the application.
func main() {
	u1 = createUserV1()
	u2 = createUserV2()
}

// createUserV1 creates a user value and passed
// a copy back to the caller.
func createUserV1() user {
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	return u
}

// createUserV2 creates a user value and shares
// the value with the caller.
func createUserV2() *user {
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	return &u
}

/*
// See escape analysis and inling decisions.

$ go build -gcflags "-m -m"
./example4.go:22: &u escapes to heap
./example4.go:22: 	from ~r0 (assign-pair) at ./example4.go:22
./example4.go:22: 	from u2 (assigned to top level variable) at ./example4.go:22
./example4.go:22: moved to heap: u
./example4.go:44: &u escapes to heap
./example4.go:44: 	from ~r0 (return) at ./example4.go:44
./example4.go:42: moved to heap: u

// See the intermediate assembly phase before
// generating the actual arch-specific assembly.

$ go build -gcflags -S
"".createUserV1 t=1 size=43 args=0x20 locals=0x0
	0x0000 00000 (example4.go:27)	TEXT	"".createUserV1(SB), $0-32
	0x0000 00000 (example4.go:27)	FUNCDATA	$0, gclocals·ff19ed39bdde8a01a800918ac3ef0ec7(SB)
	0x0000 00000 (example4.go:27)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (example4.go:21)	LEAQ	go.string."Bill"(SB), AX
	0x0007 00007 (example4.go:33)	MOVQ	AX, "".~r0+8(FP)
	0x000c 00012 (example4.go:33)	MOVQ	$4, "".~r0+16(FP)
	0x0015 00021 (example4.go:21)	LEAQ	go.string."bill@ardanlabs.com"(SB), AX
	0x001c 00028 (example4.go:33)	MOVQ	AX, "".~r0+24(FP)
	0x0021 00033 (example4.go:33)	MOVQ	$18, "".~r0+32(FP)
	0x002a 00042 (example4.go:33)	RET

// See the actual machine representation by using
// the disasembler.

$ go tool objdump -s main.main example4
TEXT main.main(SB) /go/src/.../example4.go
	example4.go:20	0x104bf10	65488b0c25a0080000	GS MOVQ GS:0x8a0, CX
	example4.go:20	0x104bf19	483b6110		CMPQ 0x10(CX), SP
	example4.go:20	0x104bf1d	0f8647010000		JBE 0x104c06a
	example4.go:20	0x104bf23	4883ec30		SUBQ $0x30, SP
	example4.go:20	0x104bf27	48896c2428		MOVQ BP, 0x28(SP)
	example4.go:20	0x104bf2c	488d6c2428		LEAQ 0x28(SP), BP

// See a list of the symbols in an artifact with
// annotations and size.

$ go tool nm example4
 104c080 T main.init
 10b2980 B main.initdone.
 104bf10 T main.main
 106ce00 R main.statictmp_2
 10984c0 B main.u1
 1098338 B main.u2
*/
