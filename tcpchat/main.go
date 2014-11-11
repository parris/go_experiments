package main

import (
    "fmt"
    "os"
)

var PORT = os.Getenv("PORT")

func main() {
    userList := make(chan *User)
    go processNewUsers(userList)
    tcp(PORT, userList)
}
