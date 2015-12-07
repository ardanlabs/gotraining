## Channels
Channels are a reference type that provide a safe mechanism to share data between goroutines. Unbuffered channel give a guarantee of delivery that data has passed from one goroutine to the other. Buffered channels allow for data to pass through the channel without such guarantees. Unbuffered channels require both a sending and receiving goroutine to be ready at the same instant before any send or receive operation can complete. Buffered channels don't force goroutines to be ready at the same instant to perform sends and receives.

## Notes

* Unbuffered channels provide a guarentee that data has been exchanged at some instant.
* Buffered channels provide great support for managing goroutines and resources.
* Closed channels can provide a system wide mechanism for notifications.
* A send on an unbuffered channel happens before the corresponding receive from that channel completes.
* A receive from an unbuffered channel happens before the send on that channel completes.
* The closing of a channel happens before a receive that returns a zero value because the channel is closed.

## Diagrams

### How an unbuffered channel works.

![Ardan Labs](unbuffered.png)

### How a buffered channel works.

![Ardan Labs](buffered.png)

## Links

http://blog.golang.org/share-memory-by-communicating

http://www.goinggo.net/2014/02/the-nature-of-channels-in-go.html

## Code Review

[Unbuffered channels - Tennis game](example1/example1.go) ([Go Playground](https://play.golang.org/p/kuxUFMqy-9))

[Unbuffered channels - Relay race](example2/example2.go) ([Go Playground](http://play.golang.org/p/r1-v3Pf0wz))

[Buffered channels - Retrieving results](example3/example3.go) ([Go Playground](http://play.golang.org/p/_nAP4DKQpL))

[Timer channels and Select](example4/example4.go) ([Go Playground](http://play.golang.org/p/OwIUceuKrM))

## Exercises

### Exercise 1
Write a program where two goroutines pass an integer back and forth ten times. Display when each goroutine receives the integer. Increment the integer with each pass. Once the integer equals ten, terminate the program cleanly.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/pkWv2bBYRf)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/IUnk_NNF5N))

### Exercise 2
Write a program that uses a buffered channel to maintain a buffer of four strings. In main, send the strings 'A', 'B', 'C' and 'D' into the channel. Then create 20 goroutines that receive a string from the channel, display the value and then send the string back into the channel. Once each goroutine is done performing that task, allow the goroutine to terminate.

[Template](exercises/template2/template2.go) ([Go Playground](http://play.golang.org/p/JRDwcCdBDi)) | 
[Answer](exercises/exercise2/exercise2.go) ([Go Playground](http://play.golang.org/p/AUViIO0quk))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
