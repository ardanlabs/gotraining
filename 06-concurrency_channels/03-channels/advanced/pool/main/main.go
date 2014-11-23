// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// This sample program demostrates how to use the pool package
// to share a simulated set of database connections.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ArdanStudios/gotraining/04-concurrency_channels/03-channels/advanced/pool"
)

const (
	// maxGoroutines is the number of routines to use.
	maxGoroutines = 25

	// pooledResources is the number of resources to pool.
	pooledResources = 2
)

var (
	// connectionID maintains a counter.
	connectionID int32

	// wg is used to wait for the program to finish.
	wg sync.WaitGroup
)

// dbConnection simulates a resource to share.
type dbConnection struct {
	ID int32
}

// Close implements the interface for the pool package.
// Close performs any resource release management.
func (dbConn *dbConnection) Close() {
	fmt.Println("Close: Connection", dbConn.ID)
	return
}

// createConnection is a factory method called by the pool
// framework when new connections are needed.
func createConnection() (pool.Resource, error) {
	id := atomic.AddInt32(&connectionID, 1)

	fmt.Println("Create: New Connection", id)
	return &dbConnection{id}, nil
}

// main is the entry point for all Go programs.
func main() {
	// Add a count for each goroutine.
	wg.Add(maxGoroutines)

	// Create the buffered channel to hold
	// and manage the connections.
	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		fmt.Println(err)
	}

	// Schedule the pool to be closed when main returns.
	defer p.Close()

	// Perform queries using a connection from the pool.
	for query := 0; query < maxGoroutines; query++ {
		go performQueries(query, p)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}

	// Wait for the goroutines to finish.
	wg.Wait()

	// Close the pool.
	fmt.Println("*****> Shutdown Program.")

}

// performQueries tests the resource pool of connections.
func performQueries(query int, p *pool.Pool) {
	// Schedule the call to Done to tell main we are done.
	defer wg.Done()

	// Acquire a connection from the pool.
	conn, err := p.Acquire()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Release the connection back to the pool.
	defer p.Release(conn)

	// Wait to simulate a query response.
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Printf("Query: QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}
