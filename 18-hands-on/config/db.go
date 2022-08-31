package config

import (
	"fmt"

	"github.com/globalsign/mgo"
)

var DB *mgo.Database

var Books *mgo.Collection

func init() {
	s, err := mgo.Dial("mongodb://testuser:dbpassword@localhost:27017/bookstore")
	if err != nil {
		panic(err)
	}

	if err = s.Ping(); err != nil {
		panic(err)
	}

	DB = s.DB("bookstore")
	Books = DB.C("books")
	fmt.Println("You connected to your database.")
}
