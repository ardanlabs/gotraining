# Basic Command Line Program

Go allows you to quickly and easily create a command line program.

In this guide, we will cover the basic concepts command line arguments using the
[http://golang.org/pkg/flag/](Flags Package).

## Simple CLI

In our first program, we will see what we get out of the box with very little effort.

Create the following program by opening a file called `cli.go` and adding the following contents:

```go
// https://play.golang.org/p/DPigLqZ5Co

package main

import (
	"flag"
	"fmt"
)

func main() {
	var cmd string

	flag.StringVar(&cmd, "cmd", cmd, `cmd can be either "hello" or "bye"`)
	flag.Parse()

	switch cmd {
	case "hello":
		fmt.Println("Hello!")
	case "bye":
		fmt.Println("Bye!")
	default:
		flag.Usage()
	}
}
```

To run the program, issue this command:

```sh
go run cli.go
```

You should get something like this:

```sh
Usage of /var/folders/l7/3s7z7s1s4n72lvj4w6g_fdmm0000gn/T/go-build844850686/command-line-arguments/_obj/exe/basic:
  -cmd="": cmd can be either "hello" or "bye"
```

## Breaking it down.

As you can see, if we don't provide any arguments, it prints out the `Usage` of the program.

Let's pass it an argument:

```sh
go run cli.go -cmd=hello
```

Now you should see that it prints 

```sh
Hello!
```


## flag.StringVar

This method allows us to tell the flag package to look for specific argument names, in this case, `cmd`.

For more information, see the definition for [http://golang.org/pkg/flag/#FlagSet.StringVar](StringVar).


## Compiling a binary

But wait, what if I want to actually compile the binary?  Easy enough, run this command:

```sh
go build cli.go
```

You will now have a file called `cli` that is an executable.  To run that, issue this command:

```sh
./cli -cmd=hello
```

## Summary

Congratulations, you just wrote your first command line program!  We only scratched the surface,
but I hope you enjoyed the quick start.





