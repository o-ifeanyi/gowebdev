package main

import (
	"gowebdev/15-hands-on/controllers"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
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

func getSession() *mgo.Session {
	m, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return m
}
