package main

import (
	"fmt"
	"net"
	"io"
)

var clients = make(map[net.Conn]bool)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	clients[conn] = true

	// Read data from the connection
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client disconnected:", conn.RemoteAddr().String())
			} else {
				fmt.Println("Error reading:", err.Error())
			}
			break
		}

		// Convert the received data to a string and display it
		data := string(buffer[:n])
		fmt.Printf("Received data from %s: %s\n", conn.RemoteAddr().String(), data)

		// Broadcast the message to all other connected clients
		for client := range clients {
			if client != conn {
				_, err := client.Write([]byte(data))
				if err != nil {
					fmt.Println("Error sending data:", err.Error())
				}
			}
		}
	}

	delete(clients, conn)
	fmt.Println("Connection closed for", conn.RemoteAddr().String())
}


func main() {
	// Set up a TCP listener on localhost:8080
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()

	fmt.Println("Chat TCP server started. Waiting for connections...")

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			return
		}

		fmt.Println("New connection established!")

		// Handle each connection concurrently
		go handleConnection(conn)
	}
}
