// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

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

	println("u1", &u1, "u2", u2)
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
// See escape analysis and inlining decisions.

$ go build -gcflags -m=2
# github.com/ardanlabs/gotraining/topics/go/language/pointers/example4
./example4.go:24:6: cannot inline createUserV1: marked go:noinline
./example4.go:38:6: cannot inline createUserV2: marked go:noinline
./example4.go:14:6: cannot inline main: non-leaf function
./example4.go:30:16: createUserV1 &u does not escape
./example4.go:46:9: &u escapes to heap
./example4.go:46:9: 	from ~r0 (return) at ./example4.go:46:2
./example4.go:39:2: moved to heap: u
./example4.go:44:16: createUserV2 &u does not escape
./example4.go:18:16: main &u1 does not escape
./example4.go:18:27: main &u2 does not escape

// See the intermediate representation phase before
// generating the actual arch-specific assembly.

$ go build -gcflags -S
0x0021 00033 (/.../example4.go:15)	CALL	"".createUserV1(SB)
0x0026 00038 (/.../example4.go:15)	MOVQ	(SP), AX
0x002a 00042 (/.../example4.go:15)	MOVQ	8(SP), CX
0x002f 00047 (/.../example4.go:15)	MOVQ	16(SP), DX
0x0034 00052 (/.../example4.go:15)	MOVQ	24(SP), BX
0x0039 00057 (/.../example4.go:15)	MOVQ	AX, "".u1+40(SP)
0x003e 00062 (/.../example4.go:15)	MOVQ	CX, "".u1+48(SP)
0x0043 00067 (/.../example4.go:15)	MOVQ	DX, "".u1+56(SP)
0x0048 00072 (/.../example4.go:15)	MOVQ	BX, "".u1+64(SP)
0x004d 00077 (/.../example4.go:16)	PCDATA	$0, $1

// See bounds checking decisions.

go build -gcflags="-d=ssa/check_bce/debug=1"

// See the actual machine representation by using
// the disasembler.

$ go tool objdump -s main.main example4
TEXT main.main(SB) github.com/ardanlabs/gotraining/topics/go/language/pointers/example4/example4.go
example4.go:15	0x104aaf1		e8aa000000		CALL main.createUserV1(SB)
example4.go:15	0x104aaf6		488b0424		MOVQ 0(SP), AX
example4.go:15	0x104aafa		488b4c2408		MOVQ 0x8(SP), CX
example4.go:15	0x104aaff		488b542410		MOVQ 0x10(SP), DX
example4.go:15	0x104ab04		488b5c2418		MOVQ 0x18(SP), BX
example4.go:15	0x104ab09		4889442428		MOVQ AX, 0x28(SP)
example4.go:15	0x104ab0e		48894c2430		MOVQ CX, 0x30(SP)
example4.go:15	0x104ab13		4889542438		MOVQ DX, 0x38(SP)
example4.go:15	0x104ab18		48895c2440		MOVQ BX, 0x40(SP)

// See a list of the symbols in an artifact with
// annotations and size.

$ go tool nm example4
104aba0 T main.createUserV1
104ac70 T main.createUserV2
104ad50 T main.init
10be460 B main.initdone.
104aad0 T main.main
*/
