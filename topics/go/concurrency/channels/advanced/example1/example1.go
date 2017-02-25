// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show the order of channel communication for unbuffered,
// buffered and closing channels based on the specification.
// https://golang.org/ref/mem#tmp_7
package main

func main() {
	unBuffered()
	buffered()
	closed()
}

// With unbuffered channels, the receive happens before the corresponding send.
// The write to a happens before the receive on c, which happens before the
// corresponding send on c completes, which happens before the print.
func unBuffered() {
	c := make(chan int)
	var a string

	go func() {
		a = "hello, world"
		<-c
	}()

	c <- 0

	// We are guaranteed to print "hello, world".
	println(a)
}

// With buffered channels, the send happens before the corresponding receive.
// The write to a happens before the send on c, which happens before the
// corresponding receive on c completes, which happens before the print.
func buffered() {
	c := make(chan int, 10)
	var a string

	go func() {
		a = "hello, world"
		c <- 0
	}()

	<-c

	// We are guaranteed to print "hello, world".
	println(a)
}

// With both types of channels, a close happens before the corresponding receive.
// The write to a happens before the close on c, which happens before the
// corresponding receive on c completes, which happens before the print.
func closed() {
	c := make(chan int, 10)
	var a string

	go func() {
		a = "hello, world"
		close(c)
	}()

	<-c

	// We are guaranteed to print "hello, world".
	println(a)
}
