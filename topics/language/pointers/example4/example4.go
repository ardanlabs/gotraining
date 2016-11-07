// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Escape Analysis Flaws:
// https://docs.google.com/document/d/1CxgUBPlx9iJzkz9JWkb6tIpTe5q32QDmz8l0BouG0Cw/view

// Sample program to show variables stay on or escape from the stack.
package main

// user represents a user in the system.
type user struct {
	name  string
	email string
}

// main is the entry point for the application.
func main() {
	stayOnStack()
	escapeToHeap()
}

// stayOnStack shows how the variable does not escape.
func stayOnStack() user {
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	return u
}

// escapeToHeap shows how the variable does escape.
func escapeToHeap() *user {
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	return &u
}

/*
// See escape analysis and inling decisions.

go build -gcflags -m

./example4.go:23: can inline stayOnStack
./example4.go:33: can inline escapeToHeap
./example4.go:17: can inline main
./example4.go:18: inlining call to stayOnStack
./example4.go:19: inlining call to escapeToHeap
./example4.go:19: main &u does not escape
./example4.go:37: moved to heap: u
./example4.go:39: &u escapes to heap


// See the intermediate assembly phase before
// generating the actual arch-specific assembly.

go build -gcflags -S

0x000f 00015 (pointers/example4/example4.go:18)	MOVQ	$4, DX
0x0016 00022 (pointers/example4/example4.go:18)	LEAQ	go.string."bill@ardanlabs.com"(SB), CX
0x001d 00029 (pointers/example4/example4.go:18)	MOVQ	$18, AX
0x0024 00036 (pointers/example4/example4.go:18)	NOP
0x0024 00036 (pointers/example4/example4.go:19)	MOVQ	$0, AX


// See the actual machine representation by using
// the disasembler.

go tool objdump -s main.main example4

example4.go:17	0x2040	4883ec28		SUBQ $0x28, SP
example4.go:17	0x2044	48896c2420		MOVQ BP, 0x20(SP)
example4.go:17	0x2049	488d6c2420		LEAQ 0x20(SP), BP
example4.go:19	0x204e	48c7042400000000	MOVQ $0x0, 0(SP)
example4.go:19	0x2056	48c744240800000000	MOVQ $0x0, 0x8(SP)

// See a list of the symbols in an artifact with
// annotations and size.

go tool nm example4

6cea0 R $f64.bfe62e42fefa39ef
96d20 B __cgo_init
96d28 B __cgo_notify_runtime_init_done
96d30 B __cgo_thread_start
4c8a0 T __rt0_amd64_darwin
4af50 T _gosave
4c8c0 T _main
*/
