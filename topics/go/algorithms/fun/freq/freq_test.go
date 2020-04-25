/*
	// This is the API you need to build for these tests. You will need to
	// change the import path in this test to point to your code.

	package freq

	// Sequential uses a sequential algorithm.
	func Sequential(text []string) map[rune]int

	// ConcurrentUnlimited uses a concurrent algorithm based on an
	// unlimited fan out pattern.
	func ConcurrentUnlimited(text []string) map[rune]int

	// ConcurrentBounded uses a concurrent algorithm based on a bounded
	// fan out and no channels.
	func ConcurrentBounded(text []string) map[rune]int

	// ConcurrentBoundedChannel uses a concurrent algorithm based on a bounded
	// fan out using a channel.
	func ConcurrentBoundedChannel(text []string) map[rune]int
*/

package freq_test

import (
	"testing"

	"github.com/ardanlabs/gotraining/topics/go/algorithms/fun/freq"
)

// go test -run none -bench . -benchtime 3s

const succeed = "\u2713"
const failed = "\u2717"

func init() {
	inp = buildText()
}

func buildText() []string {
	const n = 100
	var s []string
	for i := 0; i < n; i++ {
		s = append(s, sentence)
	}
	for k, v := range out {
		out[k] = v * n
	}
	return s
}

var sentence = `The quick brown fox jumps over the lazy dog Stay hungry. Stay
foolish Keep going. Be all in Boldness be my friend Screw it, let’s do it My
life is my message Leave no stone unturned Dream big. Pray bigger`

var inp []string

var out = map[rune]int{
	'T': 1, 'q': 1, 'p': 2, '’': 1, 'i': 11, 'b': 4, 'w': 2, 'j': 1, 'B': 2,
	'L': 1, 'e': 20, 'v': 2, 'l': 7, ',': 1, 'h': 4, 'u': 5, 'f': 4, 's': 9,
	'g': 8, 'D': 1, 'P': 1, ' ': 37, 'z': 1, 'd': 5, '.': 3, 'c': 2, 'r': 9,
	'o': 11, 'm': 5, '\n': 2, 'x': 1, 'y': 8, 'S': 3, 'K': 1, 'k': 1, 'n': 10,
	't': 8, 'a': 8, 'M': 1,
}

var m map[rune]int

func BenchmarkSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m = freq.Sequential(inp)
	}
}

func BenchmarkConcurrentBounded(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m = freq.ConcurrentBounded(inp)
	}
}

func BenchmarkConcurrentBoundedChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m = freq.ConcurrentBoundedChannel(inp)
	}
}

func BenchmarkConcurrentUnlimited(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m = freq.ConcurrentUnlimited(inp)
	}
}

func TestCount(t *testing.T) {
	t.Log("Given the need to test Frequency Count.")
	{
		tests := []struct {
			name  string
			count func(text []string) map[rune]int
		}{
			{"Sequential", freq.Sequential},
			{"ConcurrentBounded", freq.ConcurrentBounded},
			{"ConcurrentBoundedChannel", freq.ConcurrentBoundedChannel},
			{"ConcurrentUnlimited", freq.ConcurrentUnlimited},
		}

		for i, tt := range tests {
			t.Logf("\tTest %d:\tWhen running %q", i, tt.name)
			{
				f := tt.count(inp)

				if len(f) != len(out) {
					t.Logf("\t%s\tShould get back the same number of runes.", failed)
				}
				t.Logf("\t%s\tShould get back the same number of runes.", succeed)

				for r, c := range f {
					if c2, ok := out[r]; !ok || c != c2 {
						t.Logf("\t%s\tShould see ranging over result matches the output.", failed)
						t.Fatalf("\t\tRune: %c  Got %d, Expected %d.", r, c, c2)
					}
				}
				t.Logf("\t%s\tShould see ranging over result matches the output.", succeed)

				for r, c := range out {
					if c2, ok := f[r]; !ok || c != c2 {
						t.Logf("\t%s\tShould see ranging over output matches the result.", failed)
						t.Fatalf("\t\tRune: %c  Got %d, Expected %d.", r, c, c2)
					}
				}
				t.Logf("\t%s\tShould see ranging over output matches the result.", succeed)
			}
		}
	}
}
