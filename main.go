package main

import "flag"

var serverPort = flag.String("port", "3333", "The port to start this server on")

func main() {
    // Read any command line flags
    flag.Parse()

    globalListChannel := make(chan string)

    go processChanges(globalListChannel)
    tcp(globalListChannel)
}
