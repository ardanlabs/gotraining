# Single Day Go Overview

In this class, we will cover the key architectural and design aspects that distinguish Go from other imperative languages, including Go's approach to object oriented programming and concurrency. We will also learn how these characteristics make Go a strong choice for applications that demand fast development, reliability, and scalability.

Full lesson material, taught over multiple two-day courses, can be found in the [Ultimate Go section](../ultimate/README.md). Please see the [gotraining README](../../README.md) for contact information.

Examples marked with an asterisk are from [golang.org](http://golang.org/), provided under a [BSD license](https://golang.org/LICENSE).

For a great talk/video that reinforces many of these ideas, see Andrew Gerrand's [Code that Grows with Grace](https://talks.golang.org/2012/chat.slide).

## Syntax and Type Semantics

Easy to learn and reason about, trivial to machine-process.

[Variables, Builtins, and Type Safety](http://play.golang.org/p/6w6hBNE75a) ([Source](../../topics/variables/example1/example1.go))

[Constants](http://play.golang.org/p/ZHfzj2_Rse) ([Source](misc/consts/main.go))

[Structs](http://play.golang.org/p/TAX6NpPaEu) ([Source](../../topics/struct_types/example1/example1.go))

## Pointers and Reference Types

Safe and simple. Strings and slices solve many traditional pointer use-cases.

[Arrays](http://play.golang.org/p/2D24t6fbW_) ([Source](../../topics/arrays/example1/example1.go))

[Slices](http://play.golang.org/p/AFb1SZ_1WZ) ([Source](../../topics/slices/example3/example3.go))

[Appending](http://play.golang.org/p/BSNAUj2pd-) ([Source](../../topics/slices/exercises/exercise1/exercise1.go))

[Strings](http://play.golang.org/p/W3c_iWsvqj) ([Source](../../topics/slices/example5/example5.go))

[Variadic Functions](http://play.golang.org/p/5uDVuormwB) ([Source](../../topics/slices/example6/example6.go))

[Maps](http://play.golang.org/p/FcY_0ckwOZ) ([Source](../../topics/maps/example3/example3.go))

[Pointers](http://play.golang.org/p/FWmGnVUDoA) ([Source](../../topics/pointers/example2/example2.go))

[Peano Pointers](http://play.golang.org/p/7XdrgbTfZn) ([Source](https://golang.org/doc/play/peano.go)) *

## Methods and Interfaces

Allow your code to grow gracefully without you.

[Methods](http://play.golang.org/p/ovMH0wrl4B) ([Source](../../topics/methods/example1/example1.go))

[Fibonacci Closures](http://play.golang.org/p/A0nH96VB4S) ([Source](https://golang.org/doc/play/fib.go)) *

[First-Class Methods](http://play.golang.org/p/MNI1jR8Ets) ([Source](../../topics/methods/advanced/example1/example1.go))

[Interfaces](http://play.golang.org/p/CaBE4Z8-VR) ([Source](../../topics/interfaces/exercises/exercise1/exercise1.go))

## Embedding and Composition

An extension/wrapping mechanism based on composition, not inheritance.

[Embedding](http://play.golang.org/p/wAV3xnKj60) ([Source](../../topics/embedding/example2/example2.go))

[Embedding and Interfaces](http://play.golang.org/p/_MiwwXZbVI) ([Source](../../topics/embedding/example3/example3.go))

[Composition](http://play.golang.org/p/QnkL-UIVJN) ([Source](../../topics/composition/example2/example2.go))

## Concurrency

Goroutines and channels for convenient and safe concurrent algorithms.

[Concurrent Pi](http://play.golang.org/p/RdbPXQcZHi) ([Source](https://golang.org/doc/play/pi.go)) *

[Concurrent Prime Sieve](https://golang.org/s/prime-sieve) ([Source](https://golang.org/doc/play/sieve.go)) *

[Select](http://play.golang.org/p/Sv_eWCWqiJ) ([Source](../../topics/channels/example4/example4.go))

## Standard Library

Excellent value for its volume. I/O is especially powerful.

[Packages](../../topics/exporting/example5/example5.go) and [Exporting](../../topics/exporting/example5/users/users.go)

[Errors](http://play.golang.org/p/PiSDQj1UCB) ([Source](../../topics/error_handling/example1/example1.go))

[Error Variables](https://play.golang.org/p/FRnwmQx_ZI) ([Source](../../topics/error_handling/example2/example2.go))

[Behavior As Error Context](http://play.golang.org/p/6GYqwSxHjI) ([Source](../../topics/error_handling/example4/example4.go))

[Basic cURL](misc/curl1/gocurl.go)

[Enhanced cURL](misc/curl2/gocurl.go)

[HTTP Server](misc/serv/serv.go)

## Tooling

Universally adopted formatting, testing, and build tools. Many analysis tools are available (and custom tools are easy to write).

[Tooling Overview](http://go-talks.appspot.com/github.com/xtblog/gotalks/tools.slide)

___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
