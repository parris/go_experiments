package main

import (
    "net"
    "os"
    "fmt"
    "flag"
)

var (
    serverPort = flag.String("port", "3333", "The port to start this server on")
    globalList = TodoList{}
    newItems = make(chan int)
)

// func output() {
//     fmt.Println(globalList.toString())
// }

// Handles incoming requests.
func handleRequest(conn net.Conn) {
    // Make a buffer to hold incoming data.
    buf := make([]byte, 1024)

    // Read the incoming connection into the buffer.
    _, err := conn.Read(buf)

    if err != nil {
        fmt.Println("Error reading:", err.Error())
    }

    // Send a response back to person contacting us.
    conn.Write([]byte("Message received."))

    // Close the connection when you're done with it.
    conn.Close()
}

func main() {
    // Read any command line flags
    flag.Parse()

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
        go handleRequest(conn)
    }

    // go outputChanges()
}
