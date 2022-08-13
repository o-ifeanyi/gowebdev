package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/apply", apply)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	if err := tpl.ExecuteTemplate(w, "index.gohtml", nil); err != nil {
		log.Fatalln(err)
	}
}

func about(w http.ResponseWriter, req *http.Request) {
	if err := tpl.ExecuteTemplate(w, "about.gohtml", nil); err != nil {
		log.Fatalln(err)
	}
}

func contact(w http.ResponseWriter, req *http.Request) {
	if err := tpl.ExecuteTemplate(w, "contact.gohtml", nil); err != nil {
		log.Fatalln(err)
	}
}

func apply(w http.ResponseWriter, req *http.Request) {
	if err := tpl.ExecuteTemplate(w, "apply.gohtml", nil); err != nil {
		log.Fatalln(err)
	}
}
