package main

import (
	"test/handlers"
	"test/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.Connect()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/users", handlers.FindAllUsers)
	r.GET("/user/:id", handlers.FindUser)
	r.POST("/user/create", handlers.CreateUser)
	r.PUT("/user/update/:id", handlers.UpdateUser)
	r.DELETE("/user/delete/:id", handlers.DeleteUser)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
