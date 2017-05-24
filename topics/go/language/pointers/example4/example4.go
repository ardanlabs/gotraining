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
	u1 := createUserV1()
	u2 := createUserV2()

	println("u1", &u1, "u2", &u2)
}

// createUserV1 creates a user value and passed
// a copy back to the caller.
//go:noinline
func createUserV1() user {
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	println("V1", &u)

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

	println("V2", &u)

	return &u
}

/*
// See escape analysis and inling decisions.

$ go build -gcflags "-m -m"
# github.com/ardanlabs/gotraining/topics/go/language/pointers/example4
./example4.go:27: cannot inline createUserV1: marked go:noinline
./example4.go:41: cannot inline createUserV2: marked go:noinline
./example4.go:17: cannot inline main: non-leaf function
./example4.go:33: createUserV1 &u does not escape
./example4.go:49: &u escapes to heap
./example4.go:49: 	from ~r0 (return) at ./example4.go:49
./example4.go:45: moved to heap: u
./example4.go:47: createUserV2 &u does not escape
./example4.go:21: main &u1 does not escape
./example4.go:21: main &u2 does not escape

// See the intermediate assembly phase before
// generating the actual arch-specific assembly.

$ go build -gcflags -S
"".createUserV1 t=1 size=221 args=0x20 locals=0x38
	0x0000 00000 (github.com/ardanlabs/gotraining/.../example4.go:27)	TEXT	"".createUserV1(SB), $56-32
	0x0000 00000 (github.com/ardanlabs/gotraining/.../example4.go:27)	MOVQ	(TLS), CX
	0x0009 00009 (github.com/ardanlabs/gotraining/.../example4.go:27)	CMPQ	SP, 16(CX)
	0x000d 00013 (github.com/ardanlabs/gotraining/.../example4.go:27)	JLS	211
	0x0013 00019 (github.com/ardanlabs/gotraining/.../example4.go:27)	SUBQ	$56, SP
	0x0017 00023 (github.com/ardanlabs/gotraining/.../example4.go:27)	MOVQ	BP, 48(SP)
	0x001c 00028 (github.com/ardanlabs/gotraining/.../example4.go:27)	LEAQ	48(SP), BP
// See the actual machine representation by using
// the disasembler.

$ go tool objdump -s main.main example4
TEXT main.main(SB) github.com/ardanlabs/gotraining/topics/go/language/pointers/example4/example4.go
	example4.go:18	0x104bf31	e8ba000000		CALL main.createUserV1(SB)
	example4.go:18	0x104bf36	488b0424		MOVQ 0(SP), AX
	example4.go:18	0x104bf3a	488b4c2408		MOVQ 0x8(SP), CX
	example4.go:18	0x104bf3f	488b542410		MOVQ 0x10(SP), DX
	example4.go:18	0x104bf44	488b5c2418		MOVQ 0x18(SP), BX
	example4.go:18	0x104bf49	4889442428		MOVQ AX, 0x28(SP)
	example4.go:18	0x104bf4e	48894c2430		MOVQ CX, 0x30(SP)
	example4.go:18	0x104bf53	4889542438		MOVQ DX, 0x38(SP)
	example4.go:18	0x104bf58	48895c2440		MOVQ BX, 0x40(SP)

// See a list of the symbols in an artifact with
// annotations and size.

$ go tool nm example4
 104bff0 T main.createUserV1
 104c0d0 T main.createUserV2
 104c1e0 T main.init
 10b2940 B main.initdone.
 104bf10 T main.main
 106cf80 R main.statictmp_4
*/
