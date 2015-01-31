## HTTP

HTTP support in the standard library gives you everything you need to build web api's and web sites.

## Notes

* The HTTP package is a fully compliant with version 1.1 of the protocol.
* There is support for SSL/TLS.
* Adding routing and middleware to your applications requires a few simple patterns.
* Building views is support with the template package.

## Links

http://golang.org/pkg/net/http/

https://golang.org/doc/articles/wiki/

https://github.com/bradfitz/http2

## Code Review

[Web App](app)  
The code is built within two packages. The service package handles the processing of HTTP requests and responses. HTML templates are used to render the views. The search package handles the processing of searches agains the different search engines. An interface called Searcher is declared to support the implementation of new Searchers.

## Exercises

### Exercise 1

TBD

___
[![Ardan Labs](../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).