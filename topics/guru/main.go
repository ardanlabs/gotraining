// use guru referrers and channel-peers analysis on `ch`

package main

import "log"

func main() {
	ch := Make()

	go Send(ch)
	Recv(ch)

	go Recv(ch)
	Send(ch)

	NotApplicable(ch)

	go Close(ch)

	Recv(ch)
	Recv(ch)
}

func Make() chan bool {
	return make(chan bool)
}

func Send(ch chan bool) {
	ch <- true
}

func Recv(ch chan bool) {
	<-ch
	log.Println("received")
}

func Close(ch chan bool) {
	close(ch)
}

func NotApplicable(ch chan bool) {
	ch = make(chan bool, 1)
	ch <- false
}
