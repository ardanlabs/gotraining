## NATS : Chat Client

This example shows how to create a basic chat client using NATS.

## Notes

* NATS provides everything needed to publish and subcribe messages into the bus.
* Sync, Async and Channels are supports.
* You can create request/response constructs as well.

## Code Review

[Chat Editbox](editbox.go) ([Go Playground](https://play.golang.org/p/kNTmzSV3Leo))  
[Chat App](main.go) ([Go Playground](https://play.golang.org/p/3WW-L0zXFkZ))  

## Exercises

### Exercise 1

Add support for chaning channels. Add a new `bot` command called `chan` where you can specify the new channel name.
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
