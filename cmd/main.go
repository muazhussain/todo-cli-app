package main

import (
	"todo/internal/command"
	"todo/internal/storage"
	"todo/pkg/todo"
)

func main() {
	todoList := todo.TodoList{}

	storage := storage.Storage[todo.TodoList]{FileName: "pkg/models/todos.json"}
	storage.Load(&todoList)

	commandFlags := command.NewCommandFlags()
	commandFlags.Execute(&todoList)

	storage.Save(todoList)
}
