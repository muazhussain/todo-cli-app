package main

import (
	"todo-cli-app/internal/command"
	"todo-cli-app/internal/storage"
	"todo-cli-app/pkg/todo"
)

func main() {
	todoList := todo.TodoList{}

	storage := storage.Storage[todo.TodoList]{FileName: "pkg/models/todos.json"}
	storage.Load(&todoList)

	commandFlags := command.NewCommandFlags()
	commandFlags.Execute(&todoList)

	storage.Save(todoList)
}
