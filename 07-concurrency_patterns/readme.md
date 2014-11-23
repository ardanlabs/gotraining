## Concurrency Patterns
There are lots of different patterns we can create with goroutines and channels. Two interesting patterns are resource pooling and concurrent searching.

## Notes

* Resource pooling is best implemented with a buffer channel.
* Concurrenct searching is best implemented with an unbuffered channel.

## Links

http://blog.golang.org/pipelines

https://talks.golang.org/2012/concurrency.slide#1

https://blog.golang.org/context

http://blog.golang.org/advanced-go-concurrency-patterns

## Code Review

[Pooling](patterns/pool)

[Search](patterns/search)

___
[![GoingGo Training](../../00-slides/images/ggt_logo.png)](http://www.goinggotraining.net)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).