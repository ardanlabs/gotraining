## Reflection - Standard Library

Reflection is the ability to inspect a value to derive type or other meta-data. Reflection can give our program incredible flexibility to work with data of different types or create values on the fly. Reflection is critical for the encoding and decoding of data.

## Notes

* The reflection package allows us to inspect our types.
* We can add "tags" to our struct fields to store and use meta-data.
* Encoding package leverages reflection and we can as well.

## Links

http://blog.golang.org/laws-of-reflection

## Code Review

### Interfaces

Example shows how to reflect over a struct type value that is stored inside an interface value.  
[Struct Types](interface/struct/struct.go) ([Go Playground](https://play.golang.org/p/kHC6nuHYty))  

Example shows how to reflect over a slice of struct type values that are stored inside an interface value.  
[Slices](interface/slice/slice.go) ([Go Playground](https://play.golang.org/p/UyRIlkjVjW))  

Example shows how to reflect over a map of struct type values that are stored inside an interface value.  
[Maps](interface/map/map.go) ([Go Playground](https://play.golang.org/p/-_niEdmavG))  

Example shows how to reflect over a struct type pointer that is stored inside an interface value.  
[Pointers](interface/pointer/pointer.go) ([Go Playground](https://play.golang.org/p/itFSg3BL0o))  

### Tags

Example shows how to reflect on a struct type with tags.  
[Tags](tag/tag.go) ([Go Playground](https://play.golang.org/p/s6FE6J58Es))  

### Inspection / Decoding

Example shows how to inspect a structs fields and display the field name, type and value.  
[Struct Types](code/inspect/struct/struct.go) ([Go Playground](https://play.golang.org/p/ahHLMtun9y))  

Example shows how to use reflection to decode an integer.  
[Integers](code/interface/integer/integer.go) ([Go Playground](https://play.golang.org/p/LmVkzpm57a))  

## Exercises

### Exercise 1
Declare a struct type that represents a request for a customer invoice. Include a CustomerID and InvoiceID field. Define tags that can be used to validate the request. Define tags that specify both the length and range for the ID to be valid. Declare a function named validate that accepts values of any type and processes the tags. Display the results of the validation.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/LKWPS9cN_n)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/pDTvc6jEjt))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
