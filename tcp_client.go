package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func receiveMessages(conn net.Conn) {
	defer conn.Close()

	// Read messages from the server and display them in the console
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading message:", err.Error())
			break
		}
		fmt.Print("Received message from server: " + message)
	}
}

func main() {
	// Connect to the TCP server running on localhost:8080
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err.Error())
		return
	}
	defer conn.Close()

	// Start a goroutine to receive messages from the server
	go receiveMessages(conn)

	// Read user input from the console and send it to the server
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter message: ")
		message, _ := reader.ReadString('\n')
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending data:", err.Error())
			break
		}
	}
}
