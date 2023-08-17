package controllers

import (
	"html/template"
	"net/http"
	"strconv"
	"fmt"

	"github.com/raykotab/todoMVC/models"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	description := r.FormValue("description")
	models.CreateTodo(title, description)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	todos := models.GetTodos()
	fmt.Println("index todos")
	fmt.Println(todos)

	tmpl := template.Must(template.ParseFiles("views/templates/index.html"))
	tmpl.Execute(w, todos)
}

func ToggleCompleteHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	completed, _ := strconv.ParseBool(r.FormValue("completed"))
	fmt.Println(id, completed)

	models.UpdateTodoStatus(id, completed)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	fmt.Println(id)

	models.DeleteTodoById(id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
