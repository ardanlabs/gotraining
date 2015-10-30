## HTTP - API's

The Go standard library provides all the critical building blocks for producing web sites and APIs.

## Notes

* net/http provides an HTTP/1.1 compliant protocol implementation.
* There is support for SSL/TLS.
* Adding routing and middleware to your applications requires only a few simple patterns.

## Links

https://golang.org/pkg/net/http/

https://golang.org/doc/articles/wiki/

https://github.com/bradfitz/http2

https://github.com/interagent/http-api-design/blob/master/README.md

http://www.restapitutorial.com/httpstatuscodes.html

## Code Review

[Hello World Server](example1/main.go) ([Go Playground](https://play.golang.org/p/vB_ZytmqC1))

[1 Line File Server](example2/main.go) ([Go Playground](https://play.golang.org/p/Qmj_C5PEs1))

[Request and Response Basics](example3/main.go) ([Go Playground](https://play.golang.org/p/SIk8XWmwWa))

## Advanced Code Review

[Web API](api)  
Sample code that provides best practices for building a RESTful API in Go. It leverages the standard library except for the router where a package named [httptreemux](https://github.com/dimfeld/httptreemux) is used. This router provides some nice conveniences such as handling verbs and access to parameters.

## Exercises

### Exercise 1

TBD

___
[![Ardan Labs](../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
