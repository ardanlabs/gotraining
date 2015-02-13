## HTTP - API's

HTTP support in the standard library gives you everything you need to build web api's and web sites.

## Notes

* The HTTP package is a fully compliant with version 1.1 of the protocol.
* There is support for SSL/TLS.
* Adding routing and middleware to your applications requires a few simple patterns.

## Links

http://golang.org/pkg/net/http/

https://golang.org/doc/articles/wiki/

https://github.com/bradfitz/http2

https://github.com/interagent/http-api-design/blob/master/README.md

http://www.restapitutorial.com/httpstatuscodes.html

## Code Review

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
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).