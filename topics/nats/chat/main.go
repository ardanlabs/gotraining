package main

import (
	"fmt"

	"github.com/nats-io/nats"
)

const key = "chat"

func main() {

	f := func(s string) {

	}

	// // Connect to the local nats server.
	// conn, err := nats.Connect(nats.DefaultURL)
	// if err != nil {
	// 	log.Println("Unable to connect to NATS")
	// 	return
	// }

	// // Subscribe to receive messages for the specified key.
	// sub, err := conn.Subscribe(key, receive)
	// if err != nil {
	// 	log.Println("Subscribing for specified key:", err)
	// 	return
	// }

	// go send(conn)

	// // Unsubscribe from receiving these messages.
	// if err := sub.Unsubscribe(); err != nil {
	// 	log.Println("Error unsubscribing from the bus:", err)
	// 	return
	// }

	// // Close the connection to the NATS server.
	// conn.Close()

	draw(f)
}

// receive shows a message that is received.
func receive(m *nats.Msg) {
	fmt.Println(string(m.Data))
}

func send(conn *nats.Conn) {
	for {
		fmt.Print("\n\n> ")

		var s string
		i, err := fmt.Scanln(&s)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(i, s)

		// if err := conn.Publish(key, []byte(s)); err != nil {
		// 	log.Println("Publishing a message for specified key:", err)
		// 	return
		// }
	}
}
