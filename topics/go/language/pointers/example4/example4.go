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

// main is the entry point for the application.
func main() {
	createUserV1()
	createUserV2()
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

go build -gcflags "-m -m"

./example4.go:24: can inline createUserV1 as: func() user { u := user literal; return u }
./example4.go:35: can inline createUserV2 as: func() *user { u := user literal; return &u }
./example4.go:17: can inline main as: func() { createUserV1(); createUserV2() }
./example4.go:18: inlining call to createUserV1 func() user { u := user literal; return u }
./example4.go:19: inlining call to createUserV2 func() *user { u := user literal; return &u }
./example4.go:19: main &u does not escape
./example4.go:41: &u escapes to heap
./example4.go:41: 	from ~r0 (return) at ./example4.go:41
./example4.go:39: moved to heap: u

// See the intermediate assembly phase before
// generating the actual arch-specific assembly.

go build -gcflags -S

"".createUserV1 t=1 size=43 args=0x20 locals=0x0
	0x0000 00000 example4.go:24		TEXT	"".createUserV1(SB), $0-32
	0x0000 00000 example4.go:24		FUNCDATA	$0, gclocals·ff19ed39bdde8a01a800918ac3ef0ec7(SB)
	0x0000 00000 example4.go:24		FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 example4.go:18		LEAQ	go.string."Bill"(SB), AX
	0x0007 00007 example4.go:30		MOVQ	AX, "".~r0+8(FP)
	0x000c 00012 example4.go:30		MOVQ	$4, "".~r0+16(FP)
	0x0015 00021 example4.go:18		LEAQ	go.string."bill@ardanlabs.com"(SB), AX
	0x001c 00028 example4.go:30		MOVQ	AX, "".~r0+24(FP)
	0x0021 00033 example4.go:30		MOVQ	$18, "".~r0+32(FP)
	0x002a 00042 example4.go:30		RET

// See the actual machine representation by using
// the disasembler.

go tool objdump -s main.main example4

TEXT main.main(SB) example4.go
	example4.go:17	0x104bf10	4883ec28		SUBQ $0x28, SP
	example4.go:17	0x104bf14	48896c2420		MOVQ BP, 0x20(SP)
	example4.go:17	0x104bf19	488d6c2420		LEAQ 0x20(SP), BP
	example4.go:19	0x104bf1e	48c7042400000000	MOVQ $0x0, 0(SP)
	example4.go:19	0x104bf26	48c744240800000000	MOVQ $0x0, 0x8(SP)
	example4.go:19	0x104bf2f	48c744241800000000	MOVQ $0x0, 0x18(SP)
	example4.go:19	0x104bf38	488d05c9a30100		LEAQ 0x1a3c9(IP), AX
	example4.go:19	0x104bf3f	48890424		MOVQ AX, 0(SP)
	example4.go:19	0x104bf43	48c744240804000000	MOVQ $0x4, 0x8(SP)
	example4.go:19	0x104bf4c	488d05acb30100		LEAQ 0x1b3ac(IP), AX
	example4.go:19	0x104bf53	4889442410		MOVQ AX, 0x10(SP)
	example4.go:19	0x104bf58	48c744241812000000	MOVQ $0x12, 0x18(SP)
	example4.go:20	0x104bf61	488b6c2420		MOVQ 0x20(SP), BP
	example4.go:20	0x104bf66	4883c428		ADDQ $0x28, SP
	example4.go:20	0x104bf6a	c3			RET

// See a list of the symbols in an artifact with
// annotations and size.

go tool nm example4

104bf70 T main.init
10b2940 B main.initdone.
104bf10 T main.main
10983d0 B os.executablePath
*/
