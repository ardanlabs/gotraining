## Generics

This is inital code to showcase the current implementation of the Go spec for generics that can be found [here](https://go.googlesource.com/proposal/+/refs/heads/master/design/go2draft-type-parameters.md). This code is subject to break and change as the proposal and its implementation is flushed out.

There is a [go2go](https://go2goplay.golang.org/) playground that will allow you to experiment with the current proposal. `go2go` is a [transpiler](https://en.wikipedia.org/wiki/Source-to-source_compiler) that converts generics syntax into regular Go code. This is the tooling you need to experiment with the draft.

## Very High Level Overview

This comes from the draft document and provides a nice overview of what generics support is being worked on for the first potential release.

* Functions can have an additional type parameter list that uses square brackets but otherwise looks like an ordinary parameter list: func F[T any](p T) { ... }.
* These type parameters can be used by the regular parameters and in the function body.
* Types can also have a type parameter list: type M[T any] []T.
* Each type parameter has a type constraint, just as each ordinary parameter has a type: func F[T Constraint](p T) { ... }.
* Type constraints are interface types.
* The new predeclared name any is a type constraint that permits any type.
* Interface types used as type constraints can have a list of predeclared types; only type arguments that match one of those types satisfy the constraint.
* Generic functions may only use operations permitted by the type constraint.
* Using a generic function or type requires passing type arguments.
* Type inference permits omitting the type arguments of a function call in common cases.

## Omissions

This comes from the draft document and provides a nice overview of what generics support is NOT being worked on for the first potential release.

* No specialization. There is no way to write multiple versions of a generic function that are designed to work with specific type arguments.
* No metaprogramming. There is no way to write code that is executed at compile time to generate code to be executed at run time.
* No higher level abstraction. There is no way to use a function with type arguments other than to call it or instantiate it. There is no way to use a generic type other than to instantiate it.
* No general type description. In order to use operators in a generic function, constraints list specific types, rather than describing the characteristics that a type must have. This is easy to understand but may be limiting at times.
* No covariance or contravariance of function parameters.
* No operator methods. You can write a generic container that is compile-time type-safe, but you can only access it with ordinary methods, not with syntax like c[k].
* No currying. There is no way to partially instantiate a generic function or type, other than by using a helper function or a wrapper type. All type arguments must be either explicitly passed or inferred at instantiation time.
* No variadic type parameters. There is no support for variadic type parameters, which would permit writing a single generic function that takes different numbers of both type parameters and regular parameters.
* No adaptors. There is no way for a constraint to define adaptors that could be used to support type arguments that do not already implement the constraint, such as, for example, defining an == operator in terms of an Equal method, or vice-versa.
* No parameterization on non-type values such as constants. This arises most obviously for arrays, where it might sometimes be convenient to write type Matrix[n int] [n][n]float64. It might also sometimes be useful to specify significant values for a container type, such as a default value for elements.

## Posts About This Code

Here are blog posts about this code to help you understand the code better.

[Generics Part 01: Basic Syntax](https://www.ardanlabs.com/blog/2020/07/generics-01-basic-syntax.html) - William Kennedy  
[Generics Part 02: Underlying Types](https://www.ardanlabs.com/blog/2020/08/generics-02-underlying-types.html) - William Kennedy  
[Generics Part 03: Struct Types and Data Semantics](https://www.ardanlabs.com/blog/2020/09/generics-03-struct-types-and-data-semantics.html) - William Kennedy  

## Extra Reading

Here are blog posts to help you get started learning more about the current design draft.

[Current Design Draft](https://go.googlesource.com/proposal/+/refs/heads/master/design/go2draft-type-parameters.md) - Go Team  
[The Next Step for Generics](https://blog.golang.org/generics-next-step) - Go Team  
[Early notes on the generics proposal](https://rakyll.org/generics-proposal/) - JBD  
[Generics in Go](https://bitfieldconsulting.com/golang/generics) - John Arundel  

## Installing Locally

You can install and run the `go2go` tooling locally on your machine by following these steps.

```
// Clone the current source code for Go on disk.
$ cd $HOME
$ mkdir go2go
$ cd go2go
$ git clone https://go.googlesource.com/go goroot
$ cd goroot

// Fetch all the branches and checkout dev.go2go.
$ git fetch
$ git checkout dev.go2go

// Change into the source directory and build.
$ cd src
$ ./make.bash

// Navigate back to where you cloned the repo.
$ cd $GOPATH/github.com/ardanlabs/gotraining/topics/go/generics

// Then source the `.env` file located in the generics folder.
// export GO2GO_DEST=$HOME/go2go/goroot
// export PATH="$GO2GO_DEST/bin:$PATH"
// export GOROOT="$GO2GO_DEST"
// export GO2PATH="$GO2GO_DEST/src/cmd/go2go/testdata/go2path"
$ source .env

// With those settings in your current environment use this version of Go.
$ go version
go version devel +3a25e98917 Wed Jun 17 19:42:47 2020 +0000 darwin/amd64
```

## Running Examples

In most cases, I have tried to provide a concrete solution to a problem, then a corresponding interface solution, and then a generic solution. I think this is the best way to map the code changes and visualize the differences.

For any given `.go2` source code file, just run the following command.

```
$ go tool go2go run generic.go2
```

This will transpile and execute the code.

Have Fun!!!