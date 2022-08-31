package handlers

import (
	"database/sql"
	"fmt"
	"gowebdev/18-hands-on/config"
	"gowebdev/18-hands-on/repos"
	"net/http"
)

type BookHandler struct {
	Repo repos.BookRepo
}

func NewBookHandler(r repos.BookRepo) BookHandler {
	return BookHandler{r}
}

func (h BookHandler) Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "NOT ALLOWED", http.StatusMethodNotAllowed)
		return
	}

	bks, err := h.Repo.AllBooks()
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "books.gohtml", bks)
}

func (h BookHandler) Show(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "NOT ALLOWED", http.StatusMethodNotAllowed)
		return
	}

	bk, err := h.Repo.OneBook(r)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "show.gohtml", bk)
}

func (h BookHandler) Create(w http.ResponseWriter, r *http.Request) {
	config.TPL.ExecuteTemplate(w, "create.gohtml", nil)
}

func (h BookHandler) CreateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "NOT ALLOWED", http.StatusMethodNotAllowed)
		return
	}

	bk, err := h.Repo.PutBook(r)
	if err != nil {
		http.Error(w, "NOT ACCEPTED", http.StatusNotAcceptable)
		return
	}

	config.TPL.ExecuteTemplate(w, "created.gohtml", bk)
}

func (h BookHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "NOT ALLOWED", http.StatusMethodNotAllowed)
		return
	}

	bk, err := h.Repo.OneBook(r)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "update.gohtml", bk)
}

func (h BookHandler) UpdateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "NOT ALLOWED", http.StatusMethodNotAllowed)
		return
	}

	bk, err := h.Repo.UpdateBook(r)
	if err != nil {
		http.Error(w, "BAD REQUEST", http.StatusBadRequest)
		return
	}

	config.TPL.ExecuteTemplate(w, "updated.gohtml", bk)
}

func (h BookHandler) DeleteProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "NOT ALLOWED", http.StatusMethodNotAllowed)
		return
	}

	err := h.Repo.DeleteBook(r)
	if err != nil {
		http.Error(w, "BAD REQUEST", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
