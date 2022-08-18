package main

import (
	"net/http"

	"github.com/google/uuid"
)

func getUser(w http.ResponseWriter, req *http.Request) User {
	cookie, err := req.Cookie("session")
	if err != nil {
		sid := uuid.New()
		cookie = &http.Cookie{
			Name:  "session",
			Value: sid.String(),
		}
		http.SetCookie(w, cookie)
	}

	var user User
	if uid, ok := sessionDB[cookie.Value]; ok {
		user = userDB[uid]
	}
	return user
}

func alreadyLoggedIn(req *http.Request) bool {
	cookie, err := req.Cookie("session")
	if err != nil {
		return false
	}
	uid := sessionDB[cookie.Value]
	_, ok := userDB[uid]
	return ok
}
