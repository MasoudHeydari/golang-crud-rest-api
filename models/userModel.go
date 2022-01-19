package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       uint8  `gorm:"age"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
