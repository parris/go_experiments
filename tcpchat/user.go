import (
    "fmt"
    "net"
)

var users []*User

type User struct {
    handle string
    incoming chan string
    outgoing chan string
    conn net.Conn
    quit chan bool
}

func (self *Client) Read(buffer []byte) bool {
    bytesRead, error := self.Conn.Read(buffer)
    if error != nil {
        self.Leave()
        Log(error)
        return false
    }

    return true
}

func (self *Client) Leave() {
    self.quit <- true
    self.conn.Close()
}

func processNewUsers(userList chan *User) {
    for {
        fmt.Println("1")
        newUser := <- userList
        fmt.Println(newUser.handle + " joined")
        users = append(users, newUser)
    }
}
