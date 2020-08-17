// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Program creates customers for the simulation of the sleeping barber problem
// implemented in the shop package.
package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync/atomic"
	"time"

	"github.com/ardanlabs/gotraining/topics/go/algorithms/fun/barber/shop"
)

func main() {
	const maxChairs = 10
	s := shop.Open(maxChairs)

	// Create a goroutine that is constantly, but inconsistently, generating
	// customers who are entering the shop.
	go func() {
		var id int64
		for {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			name := fmt.Sprintf("cust-%d", atomic.AddInt64(&id, 1))
			if err := s.EnterCustomer(name); err != nil {
				fmt.Printf("Customer %q told %q\n", name, err)
				if err == shop.ErrShopClosed {
					break
				}
			}
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	fmt.Println("Shutting down shop")
	s.Close()
}
