// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Package chat implements a basic chat room.
package chat

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

// temporary is declared to test for the existance of the method coming
// from the net package.
type temporary interface {
	Temporary() bool
}

// message is the data received and sent to users in the chatroom.
type message struct {
	data string
	conn net.Conn
}

// client represents a single connection in the room.
type client struct {
	name   string
	room   *Room
	reader *bufio.Reader
	writer *bufio.Writer
	wg     sync.WaitGroup
	conn   net.Conn
}

// read waits for message and sends it to the chatroom for procesing.
func (c *client) read() {
	for {

		// Wait for a message to arrive.
		line, err := c.reader.ReadString('\n')

		if err == nil {
			c.room.outgoing <- message{
				data: line,
				conn: c.conn,
			}
			continue
		}

		if e, ok := err.(temporary); ok && !e.Temporary() {
			log.Println("Temporary: Client leaving chat")
			c.wg.Done()
			return
		}

		if err == io.EOF {
			log.Println("EOF: Client leaving chat")
			c.wg.Done()
			return
		}

		log.Println("read-routine", err)
	}
}

// write is a goroutine to handle processing outgoing
// messages to this client.
func (c *client) write(m message) {
	msg := fmt.Sprintf("%s %s", c.name, m.data)
	log.Printf(msg)

	c.writer.WriteString(msg)
	c.writer.Flush()
}

// drop closes the client connection and read goroutine.
func (c *client) drop() {

	// Close the connection.
	c.conn.Close()
	c.wg.Wait()
}

// newClient create a new client for an incoming connection.
func newClient(room *Room, conn net.Conn, name string) *client {
	c := client{
		name:   name,
		room:   room,
		reader: bufio.NewReader(conn),
		writer: bufio.NewWriter(conn),
		conn:   conn,
	}

	c.wg.Add(1)
	go c.read()

	return &c
}

// Room contains a set of networked client connections.
type Room struct {
	listener net.Listener
	clients  []*client
	joining  chan net.Conn
	outgoing chan message
	shutdown chan struct{}
	wg       sync.WaitGroup
}

// sendGroupMessage sends a message to all clients in the room.
func (r *Room) sendGroupMessage(m message) {
	for _, c := range r.clients {
		if c.conn != m.conn {
			c.write(m)
		}
	}
}

// join takes a new connection and adds it to the room.
func (r *Room) join(conn net.Conn) {
	name := fmt.Sprintf("Conn: %d", len(r.clients))
	log.Println("New client joining chat:", name)

	c := newClient(r, conn, name)
	r.clients = append(r.clients, c)
}

// start turns the chatroom on.
func (r *Room) start() {
	r.wg.Add(2)

	// Chatroom processing goroutne.
	go func() {
		for {
			select {
			case message := <-r.outgoing:
				// Sent message to the group.
				r.sendGroupMessage(message)
			case conn := <-r.joining:
				// Join this connection to the room.
				r.join(conn)
			case <-r.shutdown:
				// Chatroom shutting down.
				r.wg.Done()
				return
			}
		}
	}()

	// Chatroom connection accept goroutine.
	go func() {
		var err error
		if r.listener, err = net.Listen("tcp", ":6000"); err != nil {
			log.Fatalln(err)
		}

		log.Println("Chat room started: 6000")

		for {
			conn, err := r.listener.Accept()
			if err != nil {
				// Check if the error is temporary or not.
				if e, ok := err.(temporary); ok {
					if !e.Temporary() {
						log.Println("Temporary: Chat room shutting down")
						r.wg.Done()
						return
					}
				}

				log.Println("accept-routine", err)
				continue
			}

			// Add this new connection to the room.
			r.joining <- conn
		}
	}()
}

// Close shutdown the chatroom and closes all connections.
func (r *Room) Close() error {

	// Don't accept anymore client connections.
	r.listener.Close()

	// Signal the chatroom processing goroutine to stop.
	close(r.shutdown)
	r.wg.Wait()

	// Drop all existing connections.
	for _, c := range r.clients {
		c.drop()
	}
	return nil
}

// New creates a new chatroom.
func New() *Room {

	// Create a Room value.
	chatRoom := Room{
		joining:  make(chan net.Conn),
		outgoing: make(chan message),
		shutdown: make(chan struct{}),
	}

	// Start the chatroom.
	chatRoom.start()

	// Return a pointer back to the caller.
	return &chatRoom
}
