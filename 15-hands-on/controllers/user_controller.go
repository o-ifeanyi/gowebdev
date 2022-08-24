package controllers

import (
	"encoding/json"
	"fmt"
	"gowebdev/15-hands-on/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(m *mgo.Session) *UserController {
	return &UserController{m}
}

func (uc UserController) GetUser(c *gin.Context) {
	id := c.Param("id")

	if !bson.IsObjectIdHex(id) {
		c.String(http.StatusNotFound, "User not found: %s\n", id)
		return
	}

	oid := bson.ObjectIdHex(id)
	user := models.User{}

	if err := uc.session.DB("database").C("users").FindId(oid).One(&user); err != nil {
		c.String(http.StatusNotFound, "User not found: %s\n", id)
		return
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	c.String(http.StatusOK, "Got user: %s\n", jsonUser)
}

func (uc UserController) CreatUser(c *gin.Context) {
	user := models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		return
	}
	user.Id = bson.NewObjectId()

	if err := uc.session.DB("database").C("users").Insert(user); err != nil {
		c.String(http.StatusNotFound, "User not created: %s\n", user.Id)
		return
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.String(http.StatusOK, "Created new user: %s\n", jsonUser)
}

func (uc UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if !bson.IsObjectIdHex(id) {
		c.String(http.StatusNotFound, "User not found: %s\n", id)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("database").C("users").RemoveId(oid); err != nil {
		c.String(http.StatusNotFound, "User not found: %s\n", id)
		return
	}

	c.String(http.StatusOK, "Deleted user: %s\n", id)
}
