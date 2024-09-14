package command

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"todo-cli-app/pkg/todo"
)

type CommandFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCommandFlags() *CommandFlags {
	commandFlags := &CommandFlags{}
	flag.StringVar(&commandFlags.Add, "add", "", "Add a new todo specified by title")
	flag.IntVar(&commandFlags.Del, "del", -1, "Delete a todo specified by index")
	flag.StringVar(&commandFlags.Edit, "edit", "", "Edit a todo specified by index and title. Format: id:title")
	flag.IntVar(&commandFlags.Toggle, "toggle", -1, "Toggle a todo specified by index")
	flag.BoolVar(&commandFlags.List, "list", false, "List all todos")
	flag.Parse()
	return commandFlags
}

func (commandFlags *CommandFlags) Execute(todoList *todo.TodoList) {
	switch {
	case commandFlags.Add != "":
		todoList.Add(commandFlags.Add)
	case commandFlags.Del != -1:
		todoList.Delete(commandFlags.Del)
	case commandFlags.Edit != "":
		parts := strings.SplitN(commandFlags.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Invalid edit command. Format: id:title")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Invalid edit command. Format: id:title")
			os.Exit(1)
		}
		todoList.Edit(index, parts[1])
	case commandFlags.Toggle != -1:
		todoList.Toggle(commandFlags.Toggle)
	case commandFlags.List:
		todoList.Print()
	default:
		fmt.Println("Invalid command")
		os.Exit(1)
	}
}
