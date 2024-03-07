/*
// non-loop version
package main
import (
    "net"
    "os"
)

func main() {
    strEcho := "Hello Rust server, from Go tcp client."
    servAddr := "localhost:8080"
    tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
    if err != nil {
        println("ResolveTCPAddr failed:", err.Error())
        os.Exit(1)
    }

    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    if err != nil {
        println("Dial failed:", err.Error())
        os.Exit(1)
    }

    _, err = conn.Write([]byte(strEcho))
    if err != nil {
        println("Written to server failed:", err.Error())
        os.Exit(1)
    }

    println("written to server = ", strEcho)

    reply := make([]byte, 1024)

    _, err = conn.Read(reply)
    if err != nil {
        println("Written to server failed:", err.Error())
        os.Exit(1)
    }

    println("reply from server=", string(reply))

    conn.Close()
}
*/



/*
// this create new connection for each iteration of the loop
package main

import (
    "net"
    "os"
    "time"
)

func main() {
    servAddr := "localhost:8080"
    tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
    if err != nil {
        println("ResolveTCPAddr failed:", err.Error())
        os.Exit(1)
    }

    for {
        // Connect to the server
        conn, err := net.DialTCP("tcp", nil, tcpAddr)
        if err != nil {
            println("Dial failed:", err.Error())
            os.Exit(1)
        }

        strEcho := "Hello Rust server, from Go tcp client."

        _, err = conn.Write([]byte(strEcho))
        if err != nil {
            println("Written to server failed:", err.Error())
            os.Exit(1)
        }

        println("written to server = ", strEcho)

        reply := make([]byte, 1024)

        _, err = conn.Read(reply)
        if err != nil {
            println("Written to server failed:", err.Error())
            os.Exit(1)
        }

        println("reply from server=", string(reply))

        conn.Close()

        // Add delay
        time.Sleep(time.Millisecond * 100)
    }
}

*/



/*
package main

import (
    "net"
    "os"
    "time"
)

func main() {
    servAddr := "localhost:8080"
    tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
    if err != nil {
        println("ResolveTCPAddr failed:", err.Error())
        os.Exit(1)
    }

    // Connect to the server
    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    if err != nil {
        println("Dial failed:", err.Error())
        os.Exit(1)
    }
    defer conn.Close()

    // Send and receive messages in a loop
    for {
        strEcho := "Hello Rust server, from Go tcp client."

        _, err = conn.Write([]byte(strEcho))
        if err != nil {
            println("Written to server failed:", err.Error())
            os.Exit(1)
        }

        println("written to server = ", strEcho)

        reply := make([]byte, 1024)

        _, err = conn.Read(reply)
        if err != nil {
            println("Written to server failed:", err.Error())
            os.Exit(1)
        }

        println("reply from server=", string(reply))

        // Add delay
        time.Sleep(time.Millisecond * 100)
    }
}
*/



package main

import (
    "fmt"
    "net"
    "os"
    // "time"
)

func main() {
    servAddr := "localhost:8080"
    tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
    if err != nil {
        fmt.Println("ResolveTCPAddr failed:", err.Error())
        os.Exit(1)
    }

    // Connect to the server
    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    if err != nil {
        fmt.Println("Dial failed:", err.Error())
        os.Exit(1)
    }
    defer conn.Close()

    // Send and receive messages in a loop
    counter := 1
    for {
        strEcho := fmt.Sprintf("Hello Rust server, from Go tcp client. %d", counter)

        _, err = conn.Write([]byte(strEcho))
        if err != nil {
            fmt.Println("Written to server failed:", err.Error())
            os.Exit(1)
        }

        fmt.Println("Written to server:", strEcho)

        reply := make([]byte, 1024)

        _, err = conn.Read(reply)
        if err != nil {
            fmt.Println("Read from server failed:", err.Error())
            os.Exit(1)
        }

        fmt.Println("Reply from server:", string(reply))

        counter++

        // Add delay
        // time.Sleep(time.Millisecond * 100)
    }
}
