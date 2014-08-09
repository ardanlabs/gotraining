## Goroutines - Concurrency and Channels

Goroutines are functions that are created and scheduled to be run indenpently. Goroutines are multiplexed against a shared thread that is own by context. The scheduler is responsible for the management and execution of goroutines.

## Notes

* Goroutines are functions that are scheduled to run independent.
* The scheduler uses a context that owns an OS thread and goroutine run queue.
* Concurrency is not parallelism.
* We must always maintain an accounting of running goroutines and shutdown cleanly.

## Code Review

[Goroutines and concurrency](example1/example1.go) ([Go Playground](UPDATE))

[Goroutine time slicing](example2/example2.go) ([Go Playground](http://play.golang.org/p/viYA-f4zBI))

[Goroutines and parallelism](example3/example3.go) ([Go Playground](UPDATE))

## Exercises

### Exercise 1

**Part A** Create a program that declares two anonymous functions. Once that counts up to 100 from 0 and one that counts down to 0 from 100. Display each number with an unique identifier for each goroutine. Then create goroutines from these functions and don't let main return until the goroutines complete.

**Part B** Run the program in parallel.

[Answer](exercises/exercise1/exercise1.go) ([Go Playground](NEED PLAYGROUND))

___
[![GoingGo Training](../../00-slides/images/ggt_logo.png)](http://www.goinggotraining.net)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)