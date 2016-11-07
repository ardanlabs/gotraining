## NATS

NATS acts as a central nervous system for distributed systems such as mobile devices, IoT networks, enterprise microservices and cloud native infrastructure. Unlike traditional enterprise messaging systems, NATS provides an always on ‘dial-tone’.

## Installation

		Install the NATS client
		go get github.com/nats-io/nats

		Install the NATS service
		go get github.com/nats-io/gnatsd

		Run the NATS service
		gnatsd -D -V

## Links

### Documentation

https://nats.io  
https://github.com/nats-io/nats  
https://godoc.org/github.com/nats-io/nats  

### Posts and Articles

http://www.slideshare.net/Apcera/simple-solutions-for-complex-problems  
http://bravenewgeek.com/dissecting-message-queues  
http://danielwertheim.se/nats-what-a-beautiful-protocol  

## Code Review

[Basic Examples](basic)  
[Chat Client](chat)  
[Services](services)
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
