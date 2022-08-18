package main

import (
	"html/template"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Password           []byte
	Email, First, Last string
}

var userDB = make(map[string]User)
var sessionDB = make(map[string]string)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	user := getUser(w, req)

	tpl.ExecuteTemplate(w, "index.gohtml", user)
}

func bar(w http.ResponseWriter, req *http.Request) {
	user := getUser(w, req)
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "bar.gohtml", user)
}

func signup(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var user User

	if req.Method == http.MethodPost {
		e := req.FormValue("email")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")

		if _, ok := userDB[e]; ok {
			http.Error(w, "User already exist", http.StatusForbidden)
			return
		}

		sid := uuid.New()
		c := &http.Cookie{
			Name:  "session",
			Value: sid.String(),
		}
		http.SetCookie(w, c)
		sessionDB[c.Value] = e

		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		user = User{Email: e, First: f, Last: l, Password: bs}
		userDB[e] = user
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "signup.gohtml", user)
}

func login(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var user User
	if req.Method == http.MethodPost {
		e := req.FormValue("email")
		p := req.FormValue("password")

		u, ok := userDB[e]
		if !ok {
			http.Error(w, "Email and/or password do not match", http.StatusForbidden)
			return
		}

		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Email and/or password do not match", http.StatusForbidden)
			return
		}

		sid := uuid.New()
		c := &http.Cookie{
			Name:  "session",
			Value: sid.String(),
		}
		http.SetCookie(w, c)
		sessionDB[c.Value] = e
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "login.gohtml", user)
}

func logout(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	c, _ := req.Cookie("session")

	delete(sessionDB, c.Value)

	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	http.Redirect(w, req, "/login", http.StatusSeeOther)
}
