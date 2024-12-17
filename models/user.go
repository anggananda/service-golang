package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `form:"name" json:"name"`
	Email string `form:"email" json:"email"`
	Role  string `form:"role" json:"role" gorm:"default:'staff'"`
}
