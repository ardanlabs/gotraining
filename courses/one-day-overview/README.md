# Single Day Go Overview

In this class, we will cover the key architectural and design aspects that distinguish Go from other imperative languages, including Go's approach to object oriented programming and concurrency. We will also learn how these characteristics make Go a strong choice for applications that demand fast development, reliability, and scalability.

Full lesson material, taught over multiple two-day courses, can be found in the [Ultimate Go section](../ultimate/README.md). Please see the [gotraining README](../../README.md) for contact information.

Examples marked with an asterisk are from [golang.org](http://golang.org/), provided under a [BSD license](https://golang.org/LICENSE).

For a great talk/video that reinforces many of these ideas, see Andrew Gerrand's [Code that Grows with Grace](https://talks.golang.org/2012/chat.slide).

## Syntax and Type Semantics

Easy to learn and reason about, trivial to machine-process.

[Variables, Builtins, and Type Safety](https://play.golang.org/p/B5mjJKPYLh) ([Source](../../topics/language/variables/example1/example1.go))

[Constants](https://play.golang.org/p/OqJLBLhO7_) ([Source](../../topics/language/constants/example3/example3.go))

[Structs](https://play.golang.org/p/TEmOrIxl_P) ([Source](../../topics/language/struct_types/example1/example1.go))

## Pointers and Reference Types

Safe and simple. Strings and slices solve many traditional pointer use-cases.

[Arrays](https://play.golang.org/p/wUzREuHhLY) ([Source](../../topics/language/arrays/example1/example1.go))

[Slices](https://play.golang.org/p/Okc2EZG5_M) ([Source](../../topics/language/slices/example3/example3.go))

[Appending](https://play.golang.org/p/3WKISOXA-L) ([Source](../../topics/language/slices/exercises/exercise1/exercise1.go))

[Strings](https://play.golang.org/p/x0Q5ByzxGS) ([Source](../../topics/language/slices/example6/example6.go))

[Variadic Functions](https://play.golang.org/p/aTGRT1rhoO) ([Source](../../topics/language/slices/example7/example7.go))

[Maps](https://play.golang.org/p/B2klwmqmPZ) ([Source](../../topics/language/maps/example2/example2.go))

[Pointers](https://play.golang.org/p/6GUcA7-x3j) ([Source](../../topics/language/pointers/example2/example2.go))

[Peano Pointers](https://play.golang.org/p/7XdrgbTfZn) ([Source](https://golang.org/doc/play/peano.go)) *

## Methods and Interfaces

Allow your code to grow gracefully without you.

[Methods](https://play.golang.org/p/nxAwTRWk4N) ([Source](../../topics/language/methods/example1/example1.go))

[Fibonacci Closures](https://play.golang.org/p/A0nH96VB4S) ([Source](https://golang.org/doc/play/fib.go)) *

[First-Class Methods](https://play.golang.org/p/UP7qzHN-Au) ([Source](../../topics/language/methods/example3/example3.go))

[Interfaces](https://play.golang.org/p/06fecJbfE4) ([Source](../../topics/language/interfaces/exercises/exercise1/exercise1.go))

## Embedding and Composition

An extension/wrapping mechanism based on composition, not inheritance.

[Embedding](https://play.golang.org/p/QncBd6A5A4) ([Source](../../topics/language/embedding/example2/example2.go))

[Embedding and Interfaces](https://play.golang.org/p/vMEEJ7rOb4) ([Source](../../topics/language/embedding/example3/example3.go))

[Composition](https://play.golang.org/p/ufFSFxCdEs) ([Source](../../topics/api/composition/example4/example4.go))

## Concurrency

Goroutines and channels for convenient and safe concurrent algorithms.

[Concurrent Pi](https://play.golang.org/p/RdbPXQcZHi) ([Source](https://golang.org/doc/play/pi.go)) *

[Concurrent Prime Sieve](https://golang.org/s/prime-sieve) ([Source](https://golang.org/doc/play/sieve.go)) *

[Select](https://play.golang.org/p/TsJSagQawy) ([Source](../../topics/concurrency/channels/example4/example4.go))

## Standard Library

Excellent value for its volume. I/O is especially powerful.

[Packages](../../topics/language/exporting/example5/example5.go) and [Exporting](../../topics/language/exporting/example5/users/users.go)

[Errors](https://play.golang.org/p/aSjTxzNfP2) ([Source](../../topics/api/error_handling/example1/example1.go))

[Error Variables](https://play.golang.org/p/-vBG0m1Scs) ([Source](../../topics/api/error_handling/example2/example2.go))

[Behavior As Error Context](https://play.golang.org/p/Aylgou6Gq0) ([Source](../../topics/api/error_handling/example4/example4.go))

[Basic cURL](http://play.golang.org/p/b_BxHFATti))  ([Source](../../topics/standard_library/io/example2/example2.go))

[Enhanced cURL](http://play.golang.org/p/3UeN6iAE-k))  ([Source](../../topics/standard_library/io/example3/example3.go))  

[HTTP Server](https://play.golang.org/p/S0yUXdOa-i)  ([Source](../../topics/standard_library/http/example1/example1.go))

## Tooling

Universally adopted formatting, testing, and build tools. Many analysis tools are available (and custom tools are easy to write).

[Tooling Overview](http://go-talks.appspot.com/github.com/xtblog/gotalks/tools.slide)

___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
