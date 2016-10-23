## Web - Templates

Learn the basics of using templates to parse and generate markup.

## Notes

* The standard library has much of what you need to build services and apps.
* The http package provides the building blocks.
* There are other great packages in the Go ecosystem to help.

## Links

https://golang.org/pkg/net/http/  
https://golang.org/doc/articles/wiki/  
https://github.com/GeertJohan/go.rice

## Code Review

[Basic Template - Code](example1/main.go) ([Go Playground](https://play.golang.org/p/WPOIJ0B2Lt)) | 
[Basic Template - Test](example1/main_test.go) ([Go Playground](https://play.golang.org/p/n1wqyeyioL))    
[Data Parsing - Code](example2/main.go) ([Go Playground](https://play.golang.org/p/H7ftbGtIFO) | 
[Data Parsing - Test](example2/main_test.go) ([Go Playground](https://play.golang.org/p/kPDlG_hkU4))    
[Struct Parsing - Code](example3/main.go) ([Go Playground](https://play.golang.org/p/nWxoAtMNoT)) | 
[Struct Parsing - Test](example3/main_test.go) ([Go Playground](https://play.golang.org/p/aO7mKDf6sZ))      
[Generating Markup - Code](example4/main.go) ([Go Playground](https://play.golang.org/p/7DK_ISr8Cb)) | 
[Generating Markup - Test](example4/main_test.go) ([Go Playground](https://play.golang.org/p/PNfkzwyety))    
[Complex Markup - Code](example5/main.go) ([Go Playground](https://play.golang.org/p/ljKK-0OxsX)) | 
[Complex Markup - Test](example5/main_test.go) ([Go Playground](https://play.golang.org/p/OKHfh9y1rn))    
[Serving Assets - Code](example6/main.go) ([Go Playground](https://play.golang.org/p/BNSNaDcDvW)) | 
[Serving Assets - Test](example6/main_test.go) ([Go Playground](https://play.golang.org/p/VjHHW8n1M_))    
[Bundling Assets - Code](example7/main.go) ([Go Playground](https://play.golang.org/p/Z17Aopw3QP)) | 
[Bundling Assets - Test](example7/main_test.go) ([Go Playground](https://play.golang.org/p/Sj3QrmIAzB)) | 
[Bundling Assets - Test](example7/rice-box.go) ([Go Playground](https://play.golang.org/p/dCuaTv27FM))   

## Exercises

### Exercise 1

Take the code from example 5 (Complex Markup) and add a map of key/value pairs that represents the user's roles. Then add a section to the template to render the set of roles a user has.
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
