package main

import (
    "net"
    "os"
    "fmt"
)

// Handles incoming requests.
func handleRequest(conn net.Conn, messages chan string) {
    // Make a buffer to hold incoming data.
    buf := make([]byte, 1024)

    // Read the incoming connection into the buffer.
    _, err := conn.Read(buf)

    if err != nil {
        fmt.Println("Error reading:", err.Error())
    }

    message := string(buf)

    select {
    case messages <- message:
        conn.Write([]byte("Message accepted:" + message))
    default:
        conn.Write([]byte("Message not accepted"))
    }

    // Close the connection when you're done with it.
    conn.Close()
}

func tcp(messages chan string) {
    // Start up a server
    l, err := net.Listen("tcp", ":" + *serverPort)

    // Error if port is already in use
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }

    // At the end of the program stop listening
    defer l.Close()

    // Log that everything is going alright
    fmt.Println("Listening on port " + *serverPort)

    // forever...
    for {
        // Wait for an incoming connection
        // l.Accept acts as a generator and we wait until a new connection occurs
        conn, err := l.Accept()

        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            // yea... error accepting, but carry on
            continue
        }

        // Handle connections in a new goroutine.
        go handleRequest(conn, messages)
    }
}
