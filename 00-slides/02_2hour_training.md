## Go Training
This is a 2 hour accelerated style course for existing developers who are looking to gain a working understanding of Go.

*Note: This material has been designed to be taught in a classroom environment. The code is well commented but missing some of the contextual concepts and ideas that will be covered in class.*

## Hour 1
In this hour we take our initial tour of the language. We learn about variables, types, data structures and interfaces. We also explore what is idiomatic and how the language is very orthogonal. This includes following the community standards for coding and style.

#### Installation Mac

http://www.goinggo.net/2013/06/installing-go-gocode-gdb-and-liteide.html

#### Slides

[Opening Slide Deck](day1/opening/slide1.md)

#### Language Syntax, Methods and Interfaces

[Variables](../01-language_syntax/01-variables/readme.md) | 
[Struct Types](../01-language_syntax/02-struct_types/readme.md) | 
[Pointers](../01-language_syntax/03-pointers/readme.md) | 
[Constants](../01-language_syntax/04-constants/readme.md) | 
[Named Types](../01-language_syntax/05-named_types/readme.md) | 
[Functions](../01-language_syntax/06-functions/readme.md)

#### Methods and Interfaces
[Methods](../03-methods_interfaces_embedding/01-methods/readme.md) | 
[Interfaces](../03-methods_interfaces_embedding/02-interfaces/readme.md)

## Hour 2

In this hour we look at a program that implements functionality that can be found in many Go programs being developed today. The program provides a sample to the html package to create a simple search engine. The engine supports Google, Bing and Blekko searches. You can request results for all three engines or ask for just the first result. Searches are performed concurrently. Use the GOMAXPROCS environment variables to run the searches in parallel.

[Program Documentation](../web_app/readme.md)

___
[![GoingGo Training](images/ggt_logo.png)](http://www.goinggotraining.net)
[![Ardan Studios](images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](images/ggb_logo.png)](http://www.goinggo.net)
