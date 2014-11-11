package main

import (
    "net"
    "os"
    "fmt"
)

// Handles incoming requests.
func handleRequest(conn net.Conn, userList chan *User) {
    // Make a buffer to hold incoming data.
    buf := make([]byte, 1024)

    // Read the incoming connection into the buffer.
    _, err := conn.Read(buf)

    if err != nil {
        fmt.Println("Error reading:", err.Error())
    }

    handle := string(buf)

    select {
    case userList <- &User{handle, make(chan string), make(chan string), conn, make(chan bool)}:
        conn.Write([]byte("Welcome " + handle))
    default:
        conn.Write([]byte("Can't join :("))
        // Close the connection if we couldn't add the user
        conn.Close()
    }
}

func tcp(PORT string, userList chan *User) {
    // Start up a server
    tcpListener, err := net.Listen("tcp", ":" + PORT)

    // Error if port is already in use
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }

    // At the end of the program stop listening
    defer tcpListener.Close()

    // Log that everything is going alright
    fmt.Println("TCP Server started at :" + PORT)

    // forever...
    for {
        // Wait for an incoming connection
        // tcpListener.Accept acts as a generator and we wait until a new connection occurs
        conn, err := tcpListener.Accept()

        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            // yea... error accepting, but carry on
            continue
        }

        // Handle connections in a new goroutine.
        go handleRequest(conn)
    }
}
