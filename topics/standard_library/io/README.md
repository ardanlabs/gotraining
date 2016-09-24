## io - Standard Library

The ability to stream and pass data around is incredibility important. Data is constantly coming at our programs whether over a socket, file, device, etc. Many times this data needs to just be moved from one stream. Sometimes it needs to be encrypted, hashed or stored for safe keeping. The Writer and Reader interfaces may be the most heavily used and supported interfaces in both the standard library and the community.

## Notes

* The standard library provides all the infrastructure we need to stream and work with data.
* If we implement the Reader and Writer interfaces in our own types, we get this functionality for free.
* Implementing interfaces to existing functionality saves us from both development time and bugs.

## Documentation

[Interface Declarations](documentation/interfaces.md)

## Links

http://golang.org/pkg/io/  
[io package](https://medium.com/@benbjohnson/go-walkthrough-io-package-8ac5e95a9fbd#.d2ebstv0q) - Ben Johnson  

## Code Review

[Standard Library Working Together](example1/example1.go) ([Go Playground](https://play.golang.org/p/Ikm0s6vjoi))  
[Simple curl with io.Reader and io.Writer](example2/example2.go) ([Go Playground](https://play.golang.org/p/b_BxHFATti))  
[MultiWriters with curl example](example3/example3.go) ([Go Playground](https://play.golang.org/p/3UeN6iAE-k))  
[Stream processing](example4/example4.go) ([Go Playground](https://play.golang.org/p/9h53-8jZUW))  

## Advanced Code Review

[TeeReader and io composition](advanced/example1/example1.go) ([Go Playground](https://play.golang.org/p/9QSXbjtPxe))  
[Gzip and Md5 support with curl example](advanced/example2/example2.go) ([Go Playground](https://play.golang.org/p/kN97kdqRGy))

## Exercises

### Exercise 1

Download any document from the web and display the content in the terminal and write it to a file at the same time.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/ZCqK8ek58U)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/bogTavYBEx))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
