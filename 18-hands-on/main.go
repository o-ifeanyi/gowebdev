package main

import (
	"gowebdev/18-hands-on/config"
	"gowebdev/18-hands-on/handlers"
	"gowebdev/18-hands-on/repos"
	"net/http"
)

func main() {
	db := config.DB
	bookrepo := repos.NewBookRepo(db)
	bookhandler := handlers.NewBookHandler(bookrepo)

	http.HandleFunc("/", index)
	http.HandleFunc("/books", bookhandler.Index)
	http.HandleFunc("/books/show", bookhandler.Show)
	http.HandleFunc("/books/create", bookhandler.Create)
	http.HandleFunc("/books/create/process", bookhandler.CreateProcess)
	http.HandleFunc("/books/update", bookhandler.Update)
	http.HandleFunc("/books/update/process", bookhandler.UpdateProcess)
	http.HandleFunc("/books/delete/process", bookhandler.DeleteProcess)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
