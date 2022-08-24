package controllers

import (
	"encoding/json"
	"fmt"
	"gowebdev/14-hands-on/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	session map[string]models.User
}

func NewUserController(m map[string]models.User) *UserController {
	return &UserController{m}
}

func (uc UserController) GetUser(c *gin.Context) {
	id := c.Param("id")

	if user, ok := uc.session[id]; ok {
		jsonUser, err := json.Marshal(user)
		if err != nil {
			fmt.Println(err)
		}
		c.String(http.StatusOK, "Got user: %s\n", jsonUser)
	} else {
		c.String(http.StatusNotFound, "User not found: %s\n", id)
	}
}

func (uc UserController) CreatUser(c *gin.Context) {
	user := models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		return
	}
	user.Id = uuid.New().String()

	uc.session[user.Id] = user
	models.StoreUsers(uc.session)

	jsonUser, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.String(http.StatusOK, "Created new user: %s\n", jsonUser)
}

func (uc UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if _, ok := uc.session[id]; ok {
		delete(uc.session, id)
		models.StoreUsers(uc.session)
		c.String(http.StatusOK, "Deleted user: %s\n", id)
	} else {
		c.String(http.StatusNotFound, "User not found: %s\n", id)
	}
}
