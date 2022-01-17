## Context - Standard Library

The package context defines the Context type, which carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.

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

[Context Package Semantics In Go](https://www.ardanlabs.com/blog/2019/09/context-package-semantics-in-go.html) - William Kennedy  
[Package context](https://golang.org/pkg/context) - Go Team    
[Go Concurrency Patterns: Context](https://blog.golang.org/context) - Sameer Ajmani    
[Cancellation, Context, and Plumbing](https://vimeo.com/115309491) - Sameer Ajmani    
[Using contexts to avoid leaking goroutines](https://rakyll.org/leakingctx/) - JBD    

## Code Review

**_"Context values are for request-scoped data that passes through programs in a distributed system. Litmus test: Could it be an HTTP header?" - Sameer Ajmani_**

[Store / Retrieve context values](example1/example1.go) ([Go Playground](https://play.golang.org/p/xPyS_DsbKGL))  
[WithCancel](example2/example2.go) ([Go Playground](https://play.golang.org/p/ubUSuXtsldm))  
[WithDeadline](example3/example3.go) ([Go Playground](https://play.golang.org/p/o55vCa8cjIt))  
[WithTimeout](example4/example4.go) ([Go Playground](https://play.golang.org/p/8RdBXtfDv1w))  
[Request/Response](example5/example5.go) ([Go Playground](https://play.golang.org/p/9x4kBKO-Y6q)  
[Cancellation](example6/example6.go) ([Go Playground](https://play.golang.org/p/PmhTXiCZUP1)  

## Exercises

### Exercise 1

Use the template and follow the directions. You will be writing a web handler that performs a mock database call but will timeout based on a context if the call takes too long. You will also save state into the context.

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/jIkgYBhqMNy)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/J5j1ygl6LtN))  
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
