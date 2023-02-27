package controllers

import (
	"net/http"

	"github.com/TianaNanta/goAPI/initializers"
	"github.com/TianaNanta/goAPI/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {

	// get user input
	var userInput struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Avatar   string `json:"avatar"`
	}
	c.BindJSON(&userInput)

	// hash password
	hashed, er := bcrypt.GenerateFromPassword([]byte(userInput.Password), 10)

	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// create user
	user := models.User{
		Username: userInput.Username,
		Email:    userInput.Email,
		Password: string(hashed),
		Avatar:   userInput.Avatar,
	}

	result := initializers.DB.Create(&user)
	err := result.Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	// return user
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// retrieve all users
func GetUsers(c *gin.Context) {
	var users []models.User

	initializers.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

// retrieve a single user
func GetUser(c *gin.Context) {
	var user models.User

	initializers.DB.First(&user, c.Param("id"))

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// update user password
func UpdateUserPassword(c *gin.Context) {
	var user models.User

	initializers.DB.First(&user, c.Param("id"))

	var userInput struct {
		Password string `json:"password"`
	}

	c.BindJSON(&userInput)

	user.Password = userInput.Password

	initializers.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// update user
func UpdateUser(c *gin.Context) {
	var user models.User

	initializers.DB.First(&user, c.Param("id"))

	var userInput struct {
		Username string `json:"username"`
		Avatar   string `json:"avatar"`
	}

	c.BindJSON(&userInput)

	user.Username = userInput.Username
	user.Avatar = userInput.Avatar

	initializers.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// delete user
func DeleteUser(c *gin.Context) {
	var user models.User

	initializers.DB.First(&user, c.Param("id"))

	initializers.DB.Delete(&user)

	c.Status(http.StatusOK)
}
