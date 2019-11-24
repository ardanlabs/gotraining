## Encoding - Standard Library

Encoding is the process or marshaling or unmarshaling data into different forms. Taking JSON string documents and convert them to values of our user defined types is a very common practice in many go programs today. Go's support for encoding is amazing and improves and gets faster with every release.

## Notes

* Support for Decoding and Encoding JSON and XML are provided by the standard library.
* This package gets better and better with every release.

## Links

[Decode JSON Documents In Go](https://www.ardanlabs.com/blog/2014/01/decode-json-documents-in-go.html) - William Kennedy    

## Code Review

[Unmarshal JSON documents](example1/example1.go) ([Go Playground](https://play.golang.org/p/hCWu2AbC9KP))  
[Unmarshal JSON files](example2/example2.go) ([Go Playground](https://play.golang.org/p/g5-AUzfbcUS))  
[Marshal a user defined type](example3/example3.go) ([Go Playground](https://play.golang.org/p/B01KAwC-rpX))  
[Custom Marshaler and Unmarshler](example4/example4.go) ([Go Playground](https://play.golang.org/p/SolgBvtnBUr))

## Exercises

### Exercise 1

**Part A** Create a file with an array of JSON documents that contain a user name and email address. Declare a struct type that maps to the JSON document. Using the json package, read the file and create a slice of this struct type. Display the slice.

**Part B** Marshal the slice into pretty print strings and display each element.

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/aprvkRJ50js)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/i15FjSc4F2T))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
