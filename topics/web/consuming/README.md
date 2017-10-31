## Web - Consuming Web APIs

Learn more about consuming web APIs.

## Notes

* The standard library has much of what you need to build services and apps.
* The http package provides the building blocks.
* There are other great packages in the Go ecosystem to help.

## Links

https://golang.org/pkg/net/http/  
https://golang.org/doc/articles/wiki/  
github.com/dvsekhvalnov/jose2go  

## Code Review

Default HTTP support: [Test](example1/main_test.go)  
POST calls: [Test](example2/main_test.go)  
PUT calls: [Test](example3/main_test.go)  
Client timeouts: [Test](example4/main_test.go)  
Custom transporter: [Test](example5/main_test.go)  
Signing requests with JSON Web Tokens: [Code](example6/main.go) | [Test](example6/main_test.go)  
Canceling with Context: [Test](example7/main_test.go)  

## Exercises

### Exercise 1

Call the GitHub API to get a list of contributors for the `ardanlabs/gotraining` repository.

* The API url is https://api.github.com/repos/ardanlabs/gotraining/contributors
* Docs for the API in general are https://developer.github.com/v3/
* Docs for the contributors endpoint are https://developer.github.com/v3/repos/#list-contributors
* To get around rate limiting you must generate a personal access token at https://github.com/settings/tokens

[Exercise Template](exercises/template1/main.go)  
[Exercise Answer](exercises/exercise1/main.go)

___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
