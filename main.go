package main

import (
	"net/http"
	"os"
	"path/filepath"
	"github.com/raykotab/todoMVC/controllers"
)

func main() {
	baseDir, _ := os.Getwd()
	tmplPath := filepath.Join(baseDir, "views/templates/index.html")
	indexHandler := func(w http.ResponseWriter, r *http.Request) {
		controllers.IndexHandler(w, r, tmplPath)
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/create", controllers.CreateHandler)
	http.HandleFunc("/toggle-complete", controllers.ToggleCompleteHandler)
	http.HandleFunc("/delete", controllers.DeleteHandler)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
