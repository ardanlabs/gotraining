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
//go:noinline
func createUserV1() user {
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	return u
}

// createUserV2 creates a user value and shares
// the value with the caller.
//go:noinline
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
./example4.go:28: cannot inline createUserV1: marked go:noinline
./example4.go:40: cannot inline createUserV2: marked go:noinline
./example4.go:20: cannot inline main: non-leaf function
./example4.go:46: &u escapes to heap
./example4.go:46: 	from ~r0 (return) at ./example4.go:46
./example4.go:44: moved to heap: u

// See the intermediate assembly phase before
// generating the actual arch-specific assembly.

$ go build -gcflags -S
"".createUserV1 t=1 size=43 args=0x20 locals=0x0
	0x0000 00000 (/go/src/.../example4.go:28)	TEXT	"".createUserV1(SB), $0-32
	0x0000 00000 (/go/src/.../example4.go:28)	FUNCDATA	$0, gclocals·ff19ed39bdde8a01a800918ac3ef0ec7(SB)
	0x0000 00000 (/go/src/.../example4.go:28)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (/go/src/.../example4.go:30)	LEAQ	go.string."Bill"(SB), AX
	0x0007 00007 (/go/src/.../example4.go:34)	MOVQ	AX, "".~r0+8(FP)
	0x000c 00012 (/go/src/.../example4.go:34)	MOVQ	$4, "".~r0+16(FP)
	0x0015 00021 (/go/src/.../example4.go:31)	LEAQ	go.string."bill@ardanlabs.com"(SB), AX
	0x001c 00028 (/go/src/.../example4.go:34)	MOVQ	AX, "".~r0+24(FP)
	0x0021 00033 (/go/src/.../example4.go:34)	MOVQ	$18, "".~r0+32(FP)
	0x002a 00042 (/go/src/.../example4.go:34)	RET

// See the actual machine representation by using
// the disasembler.

$ go tool objdump -s main.main example4
TEXT main.main(SB) /go/src/.../example4.go
	example4.go:21	0x104bf31	e8ba000000		CALL main.createUserV1(SB)
	example4.go:21	0x104bf36	488b442418		MOVQ 0x18(SP), AX
	example4.go:21	0x104bf3b	488b4c2410		MOVQ 0x10(SP), CX
	example4.go:21	0x104bf40	48894c2420		MOVQ CX, 0x20(SP)
	example4.go:21	0x104bf45	488b1424		MOVQ 0(SP), DX
	example4.go:21	0x104bf49	488b5c2408		MOVQ 0x8(SP), BX
	example4.go:21	0x104bf4e	48891d73c50400		MOVQ BX, 0x4c573(IP)
	example4.go:21	0x104bf55	4889057cc50400		MOVQ AX, 0x4c57c(IP)
	example4.go:21	0x104bf5c	8b05fe6b0600		MOVL 0x66bfe(IP), AX
	example4.go:21	0x104bf62	85c0			TESTL AX, AX
	example4.go:21	0x104bf64	7549			JNE 0x104bfaf
	example4.go:21	0x104bf66	48891553c50400		MOVQ DX, 0x4c553(IP)
	example4.go:21	0x104bf6d	48890d5cc50400		MOVQ CX, 0x4c55c(IP)

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
