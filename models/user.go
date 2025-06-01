package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"type:varchar(255);uniqueIndex" json:"email"`
	Password string `json:"-"`
	Notes    []Note `json:"notes"`
} 