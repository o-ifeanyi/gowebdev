package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id     bson.ObjectId `json:"id"`
	Name   string        `json:"name"`
	Gender string        `json:"gender"`
	Age    int           `json:"age"`
}
