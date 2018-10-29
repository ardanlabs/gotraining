## Variables

Variables are at the heart of the language and provide the ability to read from and write to memory. In Go, access to memory is type safe. This means the compiler takes type seriously and will not allow us to use variables outside the scope of how they are declared.

## Notes

* The purpose of all programs and all parts of those programs is to transform data from one form to the other.
* Code primarily allocates, reads and writes to memory.
* Understanding type is crucial to writing good code and understanding code.
* If you don't understand the data, you don't understand the problem.
* You understand the problem better by understanding the data.
* When variables are being declared to their zero value, use the keyword var.
* When variables are being declared and initialized, use the short variable declaration operator.

## Links

[Built-In Types](http://golang.org/ref/spec#Boolean_types)    
[Variables](https://golang.org/doc/effective_go.html#variables)    
[Gustavo's IEEE-754 Brain Teaser](https://www.ardanlabs.com/blog/2013/08/gustavos-ieee-754-brain-teaser.html) - William Kennedy    
[What's in a name](https://www.youtube.com/watch?v=sFUSP8Au_PE)    
[A brief history of “type”](http://arcanesentiment.blogspot.com/2015/01/a-brief-history-of-type.html) - Arcane Sentiment    

## Code Review

[Declare and initialize variables](example1/example1.go) ([Go Playground](https://play.golang.org/p/xD_6ghgB7wm))

## Exercises

### Exercise 1 

**Part A:** Declare three variables that are initialized to their zero value and three declared with a literal value. Declare variables of type string, int and bool. Display the values of those variables.

**Part B:** Declare a new variable of type float32 and initialize the variable by converting the literal value of Pi (3.14).

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/mQiNGaMaiAa)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/Ygxt9kW_WAV))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
