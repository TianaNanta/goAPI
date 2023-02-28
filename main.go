package main

import (
	"github.com/TianaNanta/goAPI/controllers"
	"github.com/TianaNanta/goAPI/initializers"
	"github.com/TianaNanta/goAPI/middleware"
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
	r.GET("/users/me", middleware.RequireAuth, controllers.GetMe)
	r.GET("/users/me/email", middleware.RequireAuth, controllers.GetMyEmail)

	// POST method
	r.POST("/users", controllers.SignUp)
	r.POST("/users/login", controllers.Login)

	// PUT method
	r.PUT("/users/me/password", middleware.RequireAuth, controllers.UpdateUserPassword)
	r.PUT("/users/me", middleware.RequireAuth, controllers.UpdateUser)

	// DELETE method
	r.DELETE("/users/:id", controllers.DeleteUser)

	r.Run() // listen and serve on 0.0.0.0:8080
}
