## Generics

This is inital code to showcase the current implementation of the Go spec for generics that can be found [here](https://go.googlesource.com/proposal/+/refs/heads/master/design/go2draft-type-parameters.md). This code is subject to break and change as the proposal and its implementation is flushed out.

There is a [go2go](https://go2goplay.golang.org/) playground that will allow you to experiment with the current proposal. `go2go` is a [transpiler](https://en.wikipedia.org/wiki/Source-to-source_compiler) that converts generics syntax into regular Go code. This is the tooling you need to experiment with the draft.

## Extra Reading

Here are blog posts to help you get started learning more about the current design draft.

[Current Design Draft](https://go.googlesource.com/proposal/+/refs/heads/master/design/go2draft-type-parameters) - Go Team  
[The Next Step for Generics](https://blog.golang.org/generics-next-step) - Go Team  
[Early notes on the generics proposal](https://rakyll.org/generics-proposal/) - JBD  

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