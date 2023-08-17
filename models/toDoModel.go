package models

import "fmt"

type ToDo struct {
	Id	int
	Title	string
	Description	string
	Completed bool
}

var todos []ToDo
var todoiIdCounter int = 1

func CreateTodo(title, description string) ToDo {
	fmt.Println("title, description")
	fmt.Println(title, description)
	todo := ToDo{
		Id: todoiIdCounter,
		Title: title,
		Description: description,
		Completed: false,
	}
	todos = append(todos, todo)
	todoiIdCounter++
	fmt.Println(todo.Title)
	return todo
}

func GetTodos() []ToDo {
	return todos
}

func GetTodoById(id int) *ToDo {
	for i, todo := range todos {
		if todo.Id == id {
			fmt.Println(todos[i])
			return &todos[i]
		}
	}
	return nil
}

func UpdateTodoStatus(id int, completed bool) {
	for i := range todos {
		if todos[i].Id == id{
			todos[i].Completed = completed
			fmt.Println(todos[i].Description)
			return
		}
	}
}

func DeleteTodoById(id int) {
	for i, todo := range todos {
		if todo.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return
		}
	}
}
