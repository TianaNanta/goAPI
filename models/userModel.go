package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" binding:"required" gorm:"unique"`
	Email    string `json:"email" null:"true"`
	Password string `json:"password" null:"true"`
	Avatar   string `json:"avatar" null:"true"`
}
