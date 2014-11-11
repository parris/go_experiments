package main

import "fmt"

type TodoList struct {
    list []*TodoItem
}

// Constructs a new TodoItem from a string and appends it to the list
func (self *TodoList) Add(itemName string) {
    self.list = append(self.list[:], &TodoItem{itemName})
}

func (self *TodoList) toString() string {
    output := "Todo List: \n"
    output += "----------"

    for _, element := range self.list {
        output += "\n" + element.name
    }

    return output
}

func processChanges(messages chan string) {
    globalList := TodoList{}

    fmt.Println("Start processing changes")
    // forever!
    for {
        globalList.Add(<-messages)
        fmt.Println("\n\nAdding new task...\n")
        fmt.Println(globalList.toString())
    }
}
