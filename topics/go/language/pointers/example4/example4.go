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
./example4.go:14:6: cannot inline main: function too complex: cost 132 exceeds budget 80
./example4.go:39:2: u escapes to heap:
./example4.go:39:2:   flow: ~r0 = &u:
./example4.go:39:2:     from &u (address-of) at ./example4.go:46:9
./example4.go:39:2:     from return &u (return) at ./example4.go:46:2
./example4.go:39:2: moved to heap: u

// See the intermediate representation phase before
// generating the actual arch-specific assembly.

$ go build -gcflags -S
CALL	"".createUserV1(SB)
	0x0026 00038 MOVQ	(SP), AX
	0x002a 00042 MOVQ	8(SP), CX
	0x002f 00047 MOVQ	16(SP), DX
	0x0034 00052 MOVQ	24(SP), BX
	0x0039 00057 MOVQ	AX, "".u1+40(SP)
	0x003e 00062 MOVQ	CX, "".u1+48(SP)
	0x0043 00067 MOVQ	DX, "".u1+56(SP)
	0x0048 00072 MOVQ	BX, "".u1+64(SP)
	0x004d 00077 PCDATA	$1,

// See bounds checking decisions.

go build -gcflags="-d=ssa/check_bce/debug=1"

// See the actual machine representation by using
// the disasembler.

$ go tool objdump -s main.main example4
TEXT main.main(SB) github.com/ardanlabs/gotraining/topics/go/language/pointers/example4/example4.go
  example4.go:15	0x105e281		e8ba000000		CALL main.createUserV1(SB)
  example4.go:15	0x105e286		488b0424		MOVQ 0(SP), AX
  example4.go:15	0x105e28a		488b4c2408		MOVQ 0x8(SP), CX
  example4.go:15	0x105e28f		488b542410		MOVQ 0x10(SP), DX
  example4.go:15	0x105e294		488b5c2418		MOVQ 0x18(SP), BX
  example4.go:15	0x105e299		4889442428		MOVQ AX, 0x28(SP)
  example4.go:15	0x105e29e		48894c2430		MOVQ CX, 0x30(SP)
  example4.go:15	0x105e2a3		4889542438		MOVQ DX, 0x38(SP)
  example4.go:15	0x105e2a8		48895c2440		MOVQ BX, 0x40(SP)

// See a list of the symbols in an artifact with
// annotations and size.

$ go tool nm example4
 105e340 T main.createUserV1
 105e420 T main.createUserV2
 105e260 T main.main
 10cb230 B os.executablePath
*/
