package main

import (
	"github.com/TianaNanta/goAPI/controllers"
	"github.com/TianaNanta/goAPI/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	// GET method
	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUser)

	// POST method
	r.POST("/users", controllers.CreateUser)

	// PUT method
	r.PUT("/users/:id/password", controllers.UpdateUserPassword)
	r.PUT("/users/:id", controllers.UpdateUser)

	// DELETE method
	r.DELETE("/users/:id", controllers.DeleteUser)

	r.Run() // listen and serve on 0.0.0.0:8080
}
