package main

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
