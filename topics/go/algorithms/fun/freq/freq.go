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

// Concurrent uses a concurrent algorithm.
func Concurrent(text []string) map[rune]int {
	m := make(map[rune]int)

	goroutines := runtime.NumCPU()
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

// ConcurrentChannel uses a concurrent algorithm with channels.
func ConcurrentChannel(text []string) map[rune]int {
	m := make(map[rune]int)

	g := runtime.NumCPU()
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
