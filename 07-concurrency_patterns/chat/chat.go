// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// Package chat implements a basic chat room.
package chat

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
)

// message is the data received and sent to users in the chatroom.
type message struct {
	data string
	conn net.Conn
}

// client represents every connection in the room.
type client struct {
	name   string
	cr     *ChatRoom
	reader *bufio.Reader
	writer *bufio.Writer
	wg     sync.WaitGroup
	conn   net.Conn
}

func init() {
	log.SetFlags(log.Lshortfile)
}

// read waits for message and sends it to the chatroom for procesing.
func (c *client) read() {
	for {
		// Wait for a message to arrive.
		line, err := c.reader.ReadString('\n')
		if err != nil {
			// Assume this happens right now on shutdown.
			c.wg.Done()
			return
		}

		c.cr.outgoing <- message{
			data: line,
			conn: c.conn,
		}
	}
}

// write is a goroutine to handle processing outgoing
// messages to this client.
func (c *client) write(m message) {
	c.writer.WriteString(fmt.Sprintf("%s %s", c.name, m.data))
	c.writer.Flush()
}

// drop closes the client connection and read goroutine.
func (c *client) drop() {
	// Close the connection.
	c.conn.Close()
	c.wg.Wait()
}

// newClient create a new client for an incoming connection.
func newClient(cr *ChatRoom, conn net.Conn, name string) *client {
	c := client{
		name:   name,
		cr:     cr,
		reader: bufio.NewReader(conn),
		writer: bufio.NewWriter(conn),
		conn:   conn,
	}

	c.wg.Add(1)
	go c.read()

	return &c
}

// ChatRoom contains a set of networked client connections.
type ChatRoom struct {
	listener net.Listener
	clients  []*client
	joining  chan net.Conn
	outgoing chan message
	shutdown chan struct{}
	wg       sync.WaitGroup
}

// sendGroupMessage sends a message to all clients in the room.
func (cr *ChatRoom) sendGroupMessage(m message) {
	for _, c := range cr.clients {
		if c.conn != m.conn {
			c.write(m)
		}
	}
}

// join takes a new connection and adds it to the room.
func (cr *ChatRoom) join(conn net.Conn) {
	name := fmt.Sprintf("Conn: %d", len(cr.clients))
	c := newClient(cr, conn, name)
	cr.clients = append(cr.clients, c)
}

// start turns the chatroom on.
func (cr *ChatRoom) start() {
	cr.wg.Add(2)

	// Chatroom processing goroutne.
	go func() {
		for {
			select {
			case message := <-cr.outgoing:
				// Sent message to the group.
				cr.sendGroupMessage(message)
			case conn := <-cr.joining:
				// Join this connection to the room.
				cr.join(conn)
			case <-cr.shutdown:
				// Chatroom shutting down.
				cr.wg.Done()
				return
			}
		}
	}()

	// Chatroom connection accept goroutine.
	go func() {
		var err error
		if cr.listener, err = net.Listen("tcp", ":6666"); err != nil {
			log.Fatalln(err)
		}

		for {
			conn, err := cr.listener.Accept()
			if err != nil {
				// For now assume this error is caused because
				// we are being asked to shutdown.
				cr.wg.Done()
				return
			}

			// Add this new connection to the room.
			cr.joining <- conn
		}
	}()
}

// Close shutdown the chatroom and closes all connections.
func (cr *ChatRoom) Close() {
	// Don't accept anymore client connections.
	cr.listener.Close()

	// Signal the chatroom processing goroutine to stop.
	close(cr.shutdown)
	cr.wg.Wait()

	// Drop all existing connections.
	for _, c := range cr.clients {
		c.drop()
	}
}

// New creates a new chatroom.
func New() *ChatRoom {
	// Create a ChatRoom value.
	chatRoom := ChatRoom{
		joining:  make(chan net.Conn),
		outgoing: make(chan message),
		shutdown: make(chan struct{}),
	}

	// Start the chatroom.
	chatRoom.start()

	// Return a pointer back to the caller.
	return &chatRoom
}
