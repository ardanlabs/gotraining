# Channels - Concurrency and Channels

### Code Review

[Unbuffered channels - Tennis game](example1/example1.go) ([Go Playground](http://play.golang.org/p/PMnyLciVcS))

[Unbuffered channels - Relay race](example2/example2.go) ([Go Playground](http://play.golang.org/p/5B1MxmDuZI))

[Buffered channels - Control concurrency](example3/example3.go) ([Go Playground](http://play.golang.org/p/G9Gfy1drox))

(Advanced) [Timers](advanced/timer/timer.go)

(Advanced) [Semaphores](advanced/semaphore/semaphore.go)

(Advanced) [Pooling](advanced/pool/pool.go)

### Exercise 1
Review the provided program that will be used for future exercises. Program uses a function type, closures and creates goroutines to calculate Fibonacci numbers. This program is goroutine safe thanks to the sync package.
[Starter Program](exercise.go) ([Go Playground](http://play.golang.org/p/0nAEgCR2F2))

### Exercise 2
From exercise 1, use channels instead of the sync package.
[Simple Solution](final/final.go) ([Go Playground](http://play.golang.org/p/W4_O9x-a1n))

___
[![GoingGo Training](../../00-slides/images/ggt_logo.png)](http://www.goinggotraining.net)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)