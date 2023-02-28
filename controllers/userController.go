package controllers

import (
	"net/http"
	"os"

	"github.com/TianaNanta/goAPI/initializers"
	"github.com/TianaNanta/goAPI/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

	// check if the user with the same username already exists
	eror := initializers.DB.Where("username = ?", userInput.Username).First(&user).Error
	if eror == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Username already exists",
		})
		return
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

// login user
func Login(c *gin.Context) {
	var userInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	c.BindJSON(&userInput)

	var user models.User

	initializers.DB.Where("username = ?", userInput.Username).First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong password",
		})
		return
	}

	// create a new token object
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": 15000,
	})

	// sign the token with our secret
	tokenString, eror := token.SignedString([]byte(os.Getenv("SECRET")))

	if eror != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to sign token",
		})
		return
	}

	// return the token as cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
	})
}
