package todo

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type TodoList []Todo

func (todoList *TodoList) Add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	*todoList = append(*todoList, todo)
}

func (todoList *TodoList) ValidateIndex(index int) error {
	if index < 0 || index >= len(*todoList) {
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todoList *TodoList) Delete(index int) error {
	err := todoList.ValidateIndex(index)
	if err != nil {
		return err
	}
	*todoList = append((*todoList)[:index], (*todoList)[index+1:]...)
	return nil
}

func (todoList *TodoList) Toggle(index int) error {
	err := todoList.ValidateIndex(index)
	if err != nil {
		return err
	}
	isCompleted := (*todoList)[index].Completed
	if isCompleted {
		(*todoList)[index].CompletedAt = nil
	} else {
		completionTime := time.Now()
		(*todoList)[index].CompletedAt = &completionTime
	}
	(*todoList)[index].Completed = !isCompleted
	return nil
}

func (todoList *TodoList) Edit(index int, title string) error {
	err := todoList.ValidateIndex(index)
	if err != nil {
		return err
	}
	(*todoList)[index].Title = title
	return nil
}

func (todoList *TodoList) Print() {
	table := table.New(os.Stdout)
	table.SetRowLines(true)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")
	for index, todo := range *todoList {
		completed := "❌"
		completedAt := ""
		if todo.Completed {
			completed = "✅"
			if todo.CompletedAt != nil {
				completedAt = todo.CompletedAt.Format(time.RFC1123)
			}
		}
		table.AddRow(strconv.Itoa(index), todo.Title, completed, todo.CreatedAt.Format(time.RFC1123), completedAt)
	}
	table.Render()
}
