package main

import (
	"TestResume/WebCreaterTests/controllers"
	"TestResume/WebCreaterTests/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/users", controllers.FindUser)
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.FindUserId)
	r.PATCH("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)
	r.POST("/books", controllers.CreateBook)

	r.Run()
}
