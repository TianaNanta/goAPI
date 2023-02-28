package models

import "gorm.io/gorm"

type Trosa struct {
	gorm.Model
	Amount   int  `json:"amount"`
	OwnerID  uint `json:"owner_id"`
	Owner    User `json:"owner"`
	InDeptID uint `json:"in_dept_id"`
	InDept   User `json:"in_dept"`
}
