package main

import (
	"net/http"
	"github.com/raykotab/todoMVC/controllers"
)

func main() {
	http.HandleFunc("/", controllers.IndexHandler)
	http.HandleFunc("/create", controllers.CreateHandler)
	http.HandleFunc("/toggle-complete", controllers.ToggleCompleteHandler)
	http.HandleFunc("/delete", controllers.DeleteHandler)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
