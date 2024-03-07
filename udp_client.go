/*
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Please provide host:port to connect to")
		os.Exit(1)
	}

	// Resolve the string address to a UDP address
	udpAddr, err := net.ResolveUDPAddr("udp", os.Args[1])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Dial to the address with UDP
	conn, err := net.DialUDP("udp", nil, udpAddr)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	counter := 1
	for {
		// Send a message to the server
		_, err = conn.Write([]byte(fmt.Sprintf("Hello UDP Server %d\n", counter)))
		fmt.Println("send...")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Read from the connection untill a new line is send
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		// Print the data read from the connection to the terminal
		fmt.Print("> ", string(data))

		counter++
	}
}
*/

// go udp_client equivalent to existing C# udp_client & Rust udp_server
package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	var hostPort string
	if len(os.Args) > 1 {
		hostPort = os.Args[1]
	} else {
		hostPort = "127.0.0.1:8080" // Default host:port
	}

	udpAddr, err := net.ResolveUDPAddr("udp", hostPort)
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Println("Error connecting to UDP server:", err)
		return
	}
	defer conn.Close()

	count := 0
	for {
		message := []byte(fmt.Sprintf("Hi server, %d", count))
		_, err = conn.Write(message)
		if err != nil {
			fmt.Println("Error sending message to server:", err)
			return
		}
		fmt.Printf("Sent: %s to %s\n", message, udpAddr)

		err = conn.SetReadDeadline(time.Now().Add(2 * time.Second))
		if err != nil {
			fmt.Println("Error setting read deadline:", err)
			return
		}

		buffer := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error receiving response from server:", err)
			continue
		}
		responseMessage := string(buffer[:n])
		fmt.Printf("Received from server: %s, message: %s\n", udpAddr, responseMessage)

		count++

		// time.Sleep(100 * time.Millisecond) // Wait for 0.1 second before sending the next message
	}
}
