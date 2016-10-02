## HTTP - Standard Library

The Go standard library provides all the critical building blocks for producing web sites and APIs.

## Notes

* net/http provides an HTTP/1.1 compliant protocol implementation.
* There is support for SSL/TLS.
* Adding routing and middleware to your applications requires only a few simple patterns.

## Links

https://golang.org/pkg/net/http  
https://golang.org/doc/articles/wiki  
https://github.com/bradfitz/http2  
https://github.com/interagent/http-api-design/blob/master/README.md  
http://www.restapitutorial.com/httpstatuscodes.html  
http://racksburg.com/choosing-an-http-status-code  
https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts

## Code Review

[Web API](api)  
Sample code that provides best practices for building a RESTful API in Go. It leverages the standard library except for the router where a package named [httptreemux](https://github.com/dimfeld/httptreemux) is used. This router provides some nice conveniences such as handling verbs and access to parameters.

## Exercises

### Exercise 1

**Step 1**  
Add a new set of routes to CRUD User Roles.

**Step 2**  
Add a new piece of middleware to mock authentication. Then apply this middleware to the Group of routes that create, update for delete data.

**Step 3**  
Add tests for the new CRUD API.
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
