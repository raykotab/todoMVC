package controllers

import (
	"fmt"
	"github.com/raykotab/todoMVC/models"
	"html/template"
	"net/http"
	"strconv"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	title := r.FormValue("title")
	description := r.FormValue("description")
	models.CreateTodo(title, description)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func IndexHandler(w http.ResponseWriter, r *http.Request, tmplPath string) {
	todos := models.GetTodos()

	// baseDir, _ := os.Getwd()
	// tmplPath := filepath.Join(baseDir, "views/templates/index.html")
	tmpl := template.Must(template.ParseFiles(tmplPath))
	tmpl.Execute(w, todos)
}

func ToggleCompleteHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	completed, _ := strconv.ParseBool(r.FormValue("completed"))

	models.UpdateTodoStatus(id, completed)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))

	models.DeleteTodoById(id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
