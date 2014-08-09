## Go Training
This is a 3 day, 21 hour bootcamp style course for existing developers who are looking to gain a working understanding of Go.

*Note: This material has been designed to be taught in a classroom environment. The code is well commented but missing some of the contextual concepts and ideas that will be covered in class.*

## Day 1
On this day we take our initial tour of the language. We learn about variables, types, data structures, interfaces and concurrency. We also explore what is idiomatic and how the language is very orthogonal. This includes following the community standards for coding and style.

#### Slides

[Opening Slide Deck](day1/opening/slide1.md)

#### Language Syntax

[Variables](../01-language_syntax/01-variables/readme.md) | 
[Struct Types](../01-language_syntax/02-struct_types/readme.md) | 
[Constants](../01-language_syntax/03-constants/readme.md) | 
[Pointers](../01-language_syntax/04-pointers/readme.md) | 
[Named Types](../01-language_syntax/05-named_types/readme.md) | 
[Functions](../01-language_syntax/06-functions/readme.md)

#### Arrays, Slices and Maps
[Arrays](../02-array_slices_maps/01-arrays/readme.md) | 
[Slices](../02-array_slices_maps/02-slices/readme.md) | 
[Maps](../02-array_slices_maps/03-maps/readme.md)

#### Methods, Interfaces and Embedding
[Methods](../03-methods_interfaces_embedding/01-methods/readme.md) | 
[Interfaces](../03-methods_interfaces_embedding/02-interfaces/readme.md) | 
[Embedding](../03-methods_interfaces_embedding/03-embedding/readme.md)

#### Concurrency and Channels - Part I
[Goroutines](../04-concurrency_channels/01-goroutines/readme.md) | 
[Race Conditions](../04-concurrency_channels/02-race_conditions/readme.md)

## Day 2
On this day we take go deeper into Go. We learn about channels, testing, packaging, logging, encoding, io and reflection.

#### Concurrency and Channels - Part II
[Channels](../04-concurrency_channels/03-channels/readme.md)

#### Packaging
[Packaging](../05-packaging/readme.md)

#### Testing and Benchmarking
[Database Program](../06-testing/readme.md)

#### Standard Library
[Logging](../07-standard_library/01-logging/readme.md) | 
[Encoding](../07-standard_library/02-encoding/readme.md) | 
[Writers/Readers](../07-standard_library/03-writers_readers/readme.md)

#### Reflection
[Reflection](../08-reflection/readme.md)

## Day 3

On this day we talk about more advanced topics and learn how to use Go to build a full application. The program implements functionality that can be found in many Go programs being developed today. The program pulls different data feeds from the web and compares the content against a search term. The content that matches is then displayed to the terminal window. The program reads text files, makes web calls, decodes both XML and JSON into struct type values and finally does all of this using Go concurrency to make things fast.

[Program Documentation](../go_in_action/documentation/index.md)

___
[![GoingGo Training](images/ggt_logo.png)](http://www.goinggotraining.net)
[![Ardan Studios](images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](images/ggb_logo.png)](http://www.goinggo.net)