// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// This example is provided with help by Gabriel Aszalos.

// This sample program demostrates how to use the pool package
// to share a simulated set of database connections.
package main

import (
	"fmt"
	"io"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ArdanStudios/gotraining/07-concurrency_patterns/pool"
)

const (
	maxGoroutines   = 25 // the number of routines to use.
	pooledResources = 2  // number of resources in the pool
)

// dbConnection simulates a resource to share.
type dbConnection struct {
	ID int32
}

// Close implements the io.Closer interface so dbConnection
// can be managed by the pool. Close performs any resource
// release management.
func (dbConn *dbConnection) Close() error {
	fmt.Println("Close: Connection", dbConn.ID)
	return nil
}

// idCounter provides support for giving each connection a unique id.
var idCounter int32

// DbError is a customer error type for factory issues.
type DbError struct {
	ID int32
}

// Error implements the error interface.
func (d *DbError) Error() string {
	return fmt.Sprintf("Error Creating db conneciton: %d", d.ID)
}

// createConnection is a factory method that will be called by
// the pool when a new connection is needed.
func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	if id == 13 {
		return nil, &DbError{ID: id}
	}

	fmt.Println("Create: New Connection", id)

	return &dbConnection{id}, nil
}

// main is the entry point for all Go programs.
func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	// Create the pool to manage our connections.
	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Perform queries using connections from the pool.
	for query := 0; query < maxGoroutines; query++ {
		// Each goroutine needs its own copy of the query
		// value else they will all be sharing the same query
		// variable.
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)

		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}

	// Wait for the goroutines to finish.
	wg.Wait()

	// Close the pool.
	fmt.Println("*****> Shutdown Program.")
	p.Close()
}

// performQueries tests the resource pool of connections.
func performQueries(query int, p *pool.Pool) {
	// Acquire a connection from the pool.
	conn, err := p.Acquire()
	if err != nil {
		switch e := err.(type) {
		case *DbError:
			fmt.Println("Customer DB Error Type", e)
		default:
			if err == pool.ErrPoolClosed {
				fmt.Println("Error Pool Closed", err)
			} else {
				fmt.Println("Default", err)
			}
		}
		return
	}

	// Release the connection back to the pool.
	defer p.Release(conn)

	// Wait to simulate a query response.
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Printf("Query: QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}
