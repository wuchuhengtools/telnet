package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	// Listen on all interfaces on port 74110.
	listener, err := net.Listen("tcp", "0.0.0.0:7411")
	defer listener.Close()
	if err != nil {
		panic(err)
	}
	// Accept connections in a loop.
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		// Handle the connection in a goroutine.
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	// Create a buffer to store the incoming data.
	for {
		newMsg, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading from the connection: %v", err)
			return
		}
		// Print the incoming data.
		log.Printf("Received: %s", newMsg)
		// Send the data back to the client.
		_, err = conn.Write([]byte("Server says: I got your message:" + newMsg))
	}

}
