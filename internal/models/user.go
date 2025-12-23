package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID uint `json:"id"`
	//Name     string `json:"name"`
	Email    string `json:"email" gorm:"uniqueIndex;not null"`
	Password string `json:"password,omitempty"`
	//Role     string `json:"role"`
}
