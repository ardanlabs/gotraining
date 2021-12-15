package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// =============================================================================

// doWork will execute a work function in a goroutine and return a channel of
// type Result (to be determined later) back to the caller.
type doworkFn[Result any] func(context.Context) Result

func doWork[Result any](ctx context.Context, work doworkFn[Result]) chan Result {
	ch := make(chan Result, 1)

	go func() {
		ch <- work(ctx)
		fmt.Println("doWork : work complete")
	}()

	return ch
}

// =============================================================================

// poolWork will execute a work function via a pool of goroutines and return a
// channel of type Input (to be determined later) back to the caller. Once input
// is received by any given goroutine, the work function is executed and the
// Result value is displayed.
type poolWorkFn[Input any, Result any] func(input Input) Result

func poolWork[Input any, Result any](size int, work poolWorkFn[Input, Result]) (chan Input, func()) {
	var wg sync.WaitGroup
	wg.Add(size)

	ch := make(chan Input)

	for i := 0; i < size; i++ {
		go func() {
			defer wg.Done()
			for input := range ch {
				result := work(input)
				fmt.Println("pollWork :", result)
			}
		}()
	}

	cancel := func() {
		close(ch)
		wg.Wait()
	}

	return ch, cancel
}

// =============================================================================

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100 * time.Millisecond)
	defer cancel()

	dwf := func(ctx context.Context) string {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		return "work complete"
	}

	select {
	case v := <-doWork(ctx, dwf):
		fmt.Println("main:", v)
	case <-ctx.Done():
		fmt.Println("main: timeout")
	}

	fmt.Println("-------------------------------------------------------------")

	size := runtime.GOMAXPROCS(0)
	pwf := func(input int) string {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		return fmt.Sprintf("%d : received", input)
	}
	
	ch, cancel := poolWork(size, pwf)
	defer cancel()
	
	for i := 0; i < 5; i++ {
		ch <- i
	}
}