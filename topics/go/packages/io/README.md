## io - Standard Library

The ability to stream and pass data around is incredibility important. Data is constantly coming at our programs whether over a socket, file, device, etc. Many times this data needs to just be moved from one stream. Sometimes it needs to be encrypted, hashed or stored for safe keeping. The Writer and Reader interfaces may be the most heavily used and supported interfaces in both the standard library and the community.

## Notes

* The standard library provides all the infrastructure we need to stream and work with data.
* If we implement the Reader and Writer interfaces in our own types, we get this functionality for free.
* Implementing interfaces to existing functionality saves us from both development time and bugs.

## Documentation

[Interface Declarations](documentation/interfaces.md)

## Links

[io package](https://golang.org/pkg/io/)     
[Go Walkthrough: io package](https://medium.com/@benbjohnson/go-walkthrough-io-package-8ac5e95a9fbd#.d2ebstv0q) - Ben Johnson    

## Code Review

[Standard Library Working Together](example1/example1.go) ([Go Playground](https://play.golang.org/p/n-Pz_ZEW8CJ))  
[Simple curl with io.Reader and io.Writer](example2/example2.go) ([Go Playground](https://play.golang.org/p/O28tQtijcCQ))  
[MultiWriters with curl example](example3/example3.go) ([Go Playground](https://play.golang.org/p/XAZf-VYl9I3))  
[Stream processing](example4/example4.go) ([Go Playground](https://play.golang.org/p/jc4mBb-A1wZ))  

## Advanced Code Review

[TeeReader and io composition](advanced/example1/example1.go) ([Go Playground](https://play.golang.org/p/oR6fRusVHl_m))  
[Gzip and Md5 support with curl example](advanced/example2/example2.go) ([Go Playground](https://play.golang.org/p/VPpLpE_ccll))

## Exercises

### Exercise 1

Download any document from the web and display the content in the terminal and write it to a file at the same time.

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/lORHKHse--q)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/wPjVlm7QinK))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
