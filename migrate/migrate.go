package main

import (
	"github.com/TianaNanta/goAPI/initializers"
	"github.com/TianaNanta/goAPI/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
