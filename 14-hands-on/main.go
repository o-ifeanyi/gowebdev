package main

import (
	"gowebdev/14-hands-on/controllers"
	"gowebdev/14-hands-on/models"

	"github.com/gin-gonic/gin"
)

func main() {
	session := getSession()
	uc := controllers.NewUserController(session)

	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/user/:id", uc.GetUser)
	router.POST("/user", uc.CreatUser)
	router.DELETE("/user/:id", uc.DeleteUser)

	router.Run(":8080")
}

func getSession() map[string]models.User {
	return models.LoadUsers()
}
