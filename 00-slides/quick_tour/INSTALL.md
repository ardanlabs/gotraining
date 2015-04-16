# Installing Go

Installing Go is a very simple process, but if you want the entire set of tools, it can be more complicated.

We wil cover just the basic installation in this quick start guide.

## What Operating System?

Go can run natively on most operating systems, including Windows, Mac, and Linux.

### Mac

The quickest way to install on the mac is to use `HomeBrew`.

If you don't have `HomeBrew` you can install it easily with this line:

```sh
ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
```

Once it is insatlled, simply run these commands:

```sh
brew update
brew install go
```

### Windows and Linux

You can download installers for all operating systems (Including Mac), including Windows, here:

[http://golang.org/dl/](http://golang.org/dl/)


## Setting up your Environment

Go requires you to set the  `GOPATH` for the compiler to work properly.

### GOPATH on Mac or Linux

```sh
mkdir $HOME/go
export GOPATH=$HOME/go
```

In depth details on `GOPATH can be found [https://code.google.com/p/go-wiki/wiki/GOPATH](here).

## GOPATH on  Windows

From a command prompt:

```sh
mkdir "%USERPROFILE%\go"
```

Go to the Control Panel > System > Advanced Tab > Environment Variables.

Add a new *User* Variable (not a system variable)

Variable name: GOPATH
Variable value: %USERPROFILE%\go

*NOTE:* You may need to reboot for this variable to take effect.

## Your first Go program

Ok, now it's time to create our first go program.  To do so, create a file called `hello.go` with your
preferred text editor, and add the following content:

```go
package main

import "fmt"

func main() {
    fmt.Printf("hello, world\n")
}
```

To run it, use the following command:

```sh
$ go run hello.go
hello, world
```
If you see the "hello, world" message then your Go installation is working.

*NOTE:* This was taken directly from [golang.org/doc/install](http://golang.org/doc/install)


## Summary

Congratulations, you can now write more Go code!








