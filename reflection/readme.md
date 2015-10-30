## Reflection

Reflection is the ability to inspect a value to derive type or other meta-data. Reflection can give our program incredible flexibility to work with data of different types or create values on the fly. Reflection is critical for the encoding and decoding of data.

## Notes

* The reflection package allows us to inspect our types.
* We can add "tags" to our struct fields to store and use meta-data.
* Encoding package leverages reflection and we can as well.

## Links

http://blog.golang.org/laws-of-reflection

## Code Review

[Empty Interface](example1/example1.go) ([Go Playground](https://play.golang.org/p/OSeD9F_P46))

[Reflect struct types with tags](example2/example2.go) ([Go Playground](https://play.golang.org/p/fSMITKsv3p))

[Decoding function for integers](example3/example3.go) ([Go Playground](https://play.golang.org/p/bWQ6hiVECQ))

## Exercises

### Exercise 1
Declare a struct type that represents a request for a customer invoice. Include a CustomerID and InvoiceID field. Define tags that can be used to validate the request. Define tags that specify both the length and range for the ID to be valid. Declare a function named validate that accepts values of any type and processes the tags. Display the results of the validation.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/lgMQHWpZul)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/xy-wyPrsjz))

___
[![Ardan Labs](../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
