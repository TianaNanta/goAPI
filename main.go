package main

import (
	"net/http"

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
	r.Static("/avatar", "./avatar")
	r.LoadHTMLGlob("templates/*")
	r.MaxMultipartMemory = 14 << 20 // 14 MB

	// GET method
	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUser)
	r.GET("/users/me", middleware.RequireAuth, controllers.GetMe)
	r.GET("/users/me/email", middleware.RequireAuth, controllers.GetMyEmail)
	r.GET("/users/me/avatar", middleware.RequireAuth, func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", gin.H{})
	})

	// POST method
	r.POST("/users", controllers.SignUp)
	r.POST("/users/login", controllers.Login)
	r.POST("/users/me/avatar", middleware.RequireAuth, controllers.UploadAvatar)

	// PUT method
	r.PUT("/users/me/password", middleware.RequireAuth, controllers.UpdateUserPassword)
	r.PUT("/users/me/update", middleware.RequireAuth, controllers.UpdateUser)

	// DELETE method
	r.DELETE("/users/me/delete", controllers.DeleteUser)

	r.Run() // listen and serve on 0.0.0.0:8080
}
