## Reflection - Standard Library

Reflection is the ability to inspect a value to derive type or other meta-data. Reflection can give our program incredible flexibility to work with data of different types or create values on the fly. Reflection is critical for the encoding and decoding of data.

## Notes

* The reflection package allows us to inspect our types.
* We can add "tags" to our struct fields to store and use meta-data.
* Encoding package leverages reflection and we can as well.

## Links

[The Laws of Reflection](https://blog.golang.org/laws-of-reflection) - Rob Pike    

## Code Review

### Interfaces

Example shows how to reflect over a struct type value that is stored inside an interface value.  
[Struct Types](interface/struct/struct.go) ([Go Playground](https://play.golang.org/p/YNKTJ9tqnt5))  

Example shows how to reflect over a slice of struct type values that are stored inside an interface value.  
[Slices](interface/slice/slice.go) ([Go Playground](https://play.golang.org/p/V3E0QMi_0KI))  

Example shows how to reflect over a map of struct type values that are stored inside an interface value.  
[Maps](interface/map/map.go) ([Go Playground](https://play.golang.org/p/1Pc1-xD1SWR))  

Example shows how to reflect over a struct type pointer that is stored inside an interface value.  
[Pointers](interface/pointer/pointer.go) ([Go Playground](https://play.golang.org/p/vIFwz8Y3RlS))  

### Tags

Example shows how to reflect on a struct type with tags.  
[Tags](tag/tag.go) ([Go Playground](https://play.golang.org/p/riurY960A9r))  

### Inspection / Decoding

Example shows how to inspect a structs fields and display the field name, type and value.  
[Struct Types](inspect/struct/struct.go) ([Go Playground](https://play.golang.org/p/lx0lCbDdyzT))  

Example shows how to use reflection to decode an integer.  
[Integers](inspect/integer/integer.go) ([Go Playground](https://play.golang.org/p/s3taHOHwX4A))  

## Exercises

### Exercise 1
Declare a struct type that represents a request for a customer invoice. Include a CustomerID and InvoiceID field. Define tags that can be used to validate the request. Define tags that specify both the length and range for the ID to be valid. Declare a function named validate that accepts values of any type and processes the tags. Display the results of the validation.

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/yUhGiGNC23V)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/ZgxnHRIe2M8))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
