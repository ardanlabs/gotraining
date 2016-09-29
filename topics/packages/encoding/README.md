## Encoding - Standard Library

Encoding is the process or marshaling or unmarshaling data into different forms. Taking JSON string documents and convert them to values of our user defined types is a very common practice in many go programs today. Go's support for encoding is amazing and improves and gets faster with every release.

## Notes

* Support for Decoding and Encoding JSON and XML are provided by the standard libary.
* This package gets better and better with every release.

## Links

http://www.goinggo.net/2014/01/decode-json-documents-in-go.html

## Code Review

[Unmarshal JSON documents](example1/example1.go) ([Go Playground](http://play.golang.org/p/nwUGs8q2S6))  
[Unmarshal JSON files](example2/example2.go) ([Go Playground](http://play.golang.org/p/WfDYLZ5KeH))  
[Marshal a user defined type](example3/example3.go) ([Go Playground](http://play.golang.org/p/miSv1mPnK5))  
[Custom Marshaler and Unmarshler](example4/example4.go) ([Go Playground](http://play.golang.org/p/tOriq1dE0N))

## Exercises

### Exercise 1

**Part A** Create a file with an array of JSON documents that contain a user name and email address. Declare a struct type that maps to the JSON document. Using the json package, read the file and create a slice of this struct type. Display the slice.

**Part B** Marshal the slice into pretty print strings and display each element.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/5NNPJhIQDT)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/IyNucru7hi))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
