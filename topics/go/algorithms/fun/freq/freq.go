// Package freq provides support for find the frequency in which a rune
// is found in a collection of text documents.
package freq

import (
	"runtime"
	"sync"
)

// Sequential uses a sequential algorithm.
func Sequential(text []string) map[rune]int {
	m := make(map[rune]int)
	for _, words := range text {
		for _, r := range words {
			m[r]++
		}
	}
	return m
}

// ConcurrentUnlimited uses a concurrent algorithm based on an
// unlimited fan out pattern.
func ConcurrentUnlimited(text []string) map[rune]int {
	ch := make(chan map[rune]int, len(text))
	for _, words := range text {
		go func(words string) {
			lm := make(map[rune]int)
			for _, r := range words {
				lm[r]++
			}
			ch <- lm
		}(words)
	}

	all := make(map[rune]int)
	for range text {
		lm := <-ch
		for r, c := range lm {
			all[r] += c
		}
	}

	return all
}

// ConcurrentBounded uses a concurrent algorithm based on a bounded
// fan out and no channels.
func ConcurrentBounded(text []string) map[rune]int {
	m := make(map[rune]int)

	goroutines := runtime.GOMAXPROCS(0)
	totalNumbers := len(text)
	lastGoroutine := goroutines - 1
	stride := totalNumbers / goroutines

	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(goroutines)

	for g := 0; g < goroutines; g++ {
		go func(g int) {
			lm := make(map[rune]int)
			defer func() {
				mu.Lock()
				defer mu.Unlock()
				for k, v := range lm {
					m[k] = m[k] + v
				}
				wg.Done()
			}()

			start := g * stride
			end := start + stride
			if g == lastGoroutine {
				end = totalNumbers
			}

			for _, words := range text[start:end] {
				for _, r := range words {
					lm[r]++
				}
			}
		}(g)
	}

	wg.Wait()
	return m
}

// ConcurrentBoundedChannel uses a concurrent algorithm based on a bounded
// fan out using a channel.
func ConcurrentBoundedChannel(text []string) map[rune]int {
	m := make(map[rune]int)

	g := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	wg.Add(g)

	var mu sync.Mutex
	ch := make(chan string, len(text))

	for i := 0; i < g; i++ {
		go func() {
			lm := make(map[rune]int)
			defer func() {
				mu.Lock()
				defer mu.Unlock()
				for k, v := range lm {
					m[k] = m[k] + v
				}
				wg.Done()
			}()

			for words := range ch {
				for _, r := range words {
					lm[r]++
				}
			}
		}()
	}

	for _, words := range text {
		ch <- words
	}
	close(ch)

	wg.Wait()
	return m
}
