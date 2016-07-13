## NATS : Basic Examples

These examples show the basics of using the NATS client.

## Notes

* NATS provides everything needed to publish and subcribe messages into the bus.
* Sync, Async and Channels are supports.
* You can create request/response constructs as well.

## Code Review

[Synchronous Messaging](sync/main.go) ([Go Playground](http://play.golang.org/p/ZWTXLFuLRH))  
[Asynchronous Messaging](async/main.go) ([Go Playground](http://play.golang.org/p/EW34xIuS9P))    
[Channel Messaging](channels/main.go) ([Go Playground](http://play.golang.org/p/eZofyzr96R))  
[Requests](request/main.go) ([Go Playground](http://play.golang.org/p/CjqFQrxYel))  
[Queuing](queue/main.go) ([Go Playground](http://play.golang.org/p/2c8pbezdtX))  

## Exercises

### Exercise 1

Write a program that has two goroutines playing a game of tennis. Pass the ball between the two goroutines using the NATS service. Pick a random number to determine if a goroutine missed the ball. Shut the program down cleanly once a goroutine loses.
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
