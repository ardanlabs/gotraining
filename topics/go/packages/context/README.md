## Context - Standard Library

The package context defines the Context type, which carries deadlines, cancelation signals, and other request-scoped values across API boundaries and between processes.

## Notes

* Incoming requests to a server should create a Context.
* Outgoing calls to servers should accept a Context. 
* The chain of function calls between them must propagate the Context.
* Replace a Context using WithCancel, WithDeadline, WithTimeout, or WithValue.
* When a Context is canceled, all Contexts derived from it are also canceled.
* Do not store Contexts inside a struct type; instead, pass a Context explicitly to each function that needs it.
* Do not pass a nil Context, even if a function permits it. Pass context.TODO if you are unsure about which Context to use.
* Use context Values only for request-scoped data that transits processes and APIs, not for passing optional parameters to functions.
* The same Context may be passed to functions running in different goroutines; Contexts are safe for simultaneous use by multiple goroutines.

## Links

[Package context](https://golang.org/pkg/context)  
[Go Concurrency Patterns: Context](https://blog.golang.org/context) - Sameer Ajmani  
[Cancellation, Context, and Plumbing](https://vimeo.com/115309491) - Sameer Ajmani  
[Using contexts to avoid leaking goroutines](http://golang.rakyll.org/leakingctx/) -- Jaana Burcu Dogan  

## Code Review

[Store / Retrieve context values](example1/example1.go) ([Go Playground](https://play.golang.org/p/VkLs3x-Vbd))  
[WithCancel API](example2/example2.go) ([Go Playground](https://play.golang.org/p/ZLcBz25Q4r))  
[WithDeadline API](example3/example3.go) ([Go Playground](https://play.golang.org/p/2lq42MzJAo))  
[WithTimeout API](example4/example4.go) ([Go Playground](https://play.golang.org/p/hft7tgGNxS))  
[Request/Response](example5/example5.go) ([Go Playground](https://play.golang.org/p/urtOUiAyCF))  
[Avoid leaking goroutines](example6/example6.go) ([Go Playground](https://play.golang.org/p/36YjxYZKC3))

## Exercises

### Exercise 1

Use the template and follow the directions. You will be writing a web handler that performs a mock database call but will timeout based on a context if the call takes too long. You will also save state into the context.

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/T05C1L8Mu6)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/2L_DF8-pH7))  
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
