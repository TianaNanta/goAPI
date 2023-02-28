package controllers

import (
	"net/http"

	"github.com/TianaNanta/goAPI/initializers"
	"github.com/TianaNanta/goAPI/models"
	"github.com/gin-gonic/gin"
)

// CreateTrosa for a Owner to give Trosa to a InDept
func CreateTrosa(c *gin.Context) {
	// get user id from token
	me, _ := c.Get("user")

	// get user input
	var userInput struct {
		InDeptUsername string `json:"in_dept_username" binding:"required"`
		Amount         int    `json:"amount" binding:"required"`
	}
	c.BindJSON(&userInput)

	// get receiver id from receiver username
	var indept models.User
	initializers.DB.Where("username = ?", userInput.InDeptUsername).First(&indept)

	// create trosa
	trosa := models.Trosa{
		OwnerID:  me.(models.User).ID,
		InDeptID: indept.ID,
		Amount:   userInput.Amount,
	}

	result := initializers.DB.Create(&trosa)
	err := result.Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create trosa",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"trosa": trosa,
	})
}

// GetTrosaForOwner ...
func GetTrosaForOwner(c *gin.Context) {
	// get user id from token
	me, _ := c.Get("user")

	// get trosa
	var trosa []models.Trosa
	initializers.DB.Where("owner_id = ?", me.(models.User).ID).Find(&trosa)

	c.JSON(http.StatusOK, gin.H{
		"trosa": trosa,
	})
}

// GetTrosaForInDept ...
func GetTrosaForInDept(c *gin.Context) {
	// get user id from token
	me, _ := c.Get("user")

	// get trosa
	var trosa []models.Trosa
	initializers.DB.Where("indept_id = ?", me.(models.User).ID).Find(&trosa)

	c.JSON(http.StatusOK, gin.H{
		"trosa": trosa,
	})
}
