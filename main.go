package main

import (
	"fmt"

	"github.com/saimib/code-kata-be-go-solution/todoist"
	"github.com/saimib/code-kata-be-go-solution/utils"
)

func main() {
	evenTodos := todoist.NewTodoist(utils.NewUtils()).GetEvenTodos()

	for _, todo := range evenTodos {
		fmt.Printf("todo title - %s || is completed? %t \n", todo.Title, todo.Completed)
	}
}
