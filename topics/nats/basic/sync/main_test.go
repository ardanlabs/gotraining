/*
	go test -run none -bench BenchmarkNextMsg -benchmem -memprofile mem.out -memprofilerate 1

	BenchmarkNextMsg-8	    1000	   1390932 ns/op	     192 B/op	       3 allocs/op

	go tool pprof --inuse_space sync.test mem.out
	ROUTINE ======================== github.com/nats-io/nats.(*Subscription).NextMsg in /Users/bill/code/go/src/github.com/nats-io/nats/nats.go
       64B        80B (flat, cum) 0.089% of Total
         .          .   1988:	mch := s.mch
         .          .   1989:	max := s.max
         .          .   1990:	s.mu.Unlock()
         .          .   1991:
         .          .   1992:	var ok bool
         .        16B   1993:	t := time.NewTimer(timeout)
         .          .   1994:	defer t.Stop()
         .          .   1995:
       64B        64B   1996:	select {
         .          .   1997:	case msg, ok = <-mch:
         .          .   1998:		if !ok {
         .          .   1999:			return nil, ErrConnectionClosed
         .          .   2000:		}
         .          .   2001:		// Update some stats.

	go tool pprof --inuse_objects sync.test mem.out
	ROUTINE ======================== github.com/nats-io/nats.(*Subscription).NextMsg in /Users/bill/code/go/src/github.com/nats-io/nats/nats.go
         1          2 (flat, cum)  3.85% of Total
         .          .   1988:	mch := s.mch
         .          .   1989:	max := s.max
         .          .   1990:	s.mu.Unlock()
         .          .   1991:
         .          .   1992:	var ok bool
         .          1   1993:	t := time.NewTimer(timeout)
         .          .   1994:	defer t.Stop()
         .          .   1995:
         1          1   1996:	select {
         .          .   1997:	case msg, ok = <-mch:
         .          .   1998:		if !ok {
         .          .   1999:			return nil, ErrConnectionClosed
         .          .   2000:		}
         .          .   2001:		// Update some stats.

    go tool pprof --alloc_space sync.test mem.out
	ROUTINE ======================== github.com/nats-io/nats.(*Subscription).NextMsg in /Users/bill/code/go/src/github.com/nats-io/nats/nats.go
       64B   206.52kB (flat, cum) 69.06% of Total
         .          .   1988:	mch := s.mch
         .          .   1989:	max := s.max
         .          .   1990:	s.mu.Unlock()
         .          .   1991:
         .          .   1992:	var ok bool
         .   206.45kB   1993:	t := time.NewTimer(timeout)
         .          .   1994:	defer t.Stop()
         .          .   1995:
       64B        64B   1996:	select {
         .          .   1997:	case msg, ok = <-mch:
         .          .   1998:		if !ok {
         .          .   1999:			return nil, ErrConnectionClosed
         .          .   2000:		}
         .          .   2001:		// Update some stats.

	go tool pprof --alloc_objects sync.test mem.out
	ROUTINE ======================== github.com/nats-io/nats.(*Subscription).NextMsg in /Users/bill/code/go/src/github.com/nats-io/nats/nats.go
         1       3305 (flat, cum) 97.21% of Total
         .          .   1988:	mch := s.mch
         .          .   1989:	max := s.max
         .          .   1990:	s.mu.Unlock()
         .          .   1991:
         .          .   1992:	var ok bool
         .       3304   1993:	t := time.NewTimer(timeout)
         .          .   1994:	defer t.Stop()
         .          .   1995:
         1          1   1996:	select {
         .          .   1997:	case msg, ok = <-mch:
         .          .   1998:		if !ok {
         .          .   1999:			return nil, ErrConnectionClosed
         .          .   2000:		}
         .          .   2001:		// Update some stats.
*/
package main

import (
	"log"
	"testing"
	"time"

	"github.com/nats-io/nats"
)

var sub *nats.Subscription

func init() {
	// Declare the key to use for publishing/subscribing.
	const key = "test"

	// Connect to the local nats server.
	conn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Println("Unable to connect to NATS")
		return
	}

	// Subscribe to receive messages for the specified key.
	// Passing nil for the handler so everything is a manual pull.
	sub, err = conn.SubscribeSync(key)
	if err != nil {
		log.Println("Subscribing for specified key:", err)
		return
	}
}

var msg *nats.Msg

func BenchmarkNextMsg(b *testing.B) {
	var m *nats.Msg

	for i := 0; i < b.N; i++ {
		m, _ = sub.NextMsg(time.Millisecond)
	}

	msg = m
}
