package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID     uint   `json:"id" gorm:"primaryKey,autoIncrement"`
	Name   string `json:"name" gorm:"not null"`
	Price  uint   `json:"price" gorm:"not null"`
	Stock  uint   `json:"stock" gorm:"not null"`
	UserID uint   `json:"userId"`
	User   User   `gorm:"foreignKey:UserID"`
	//CreatedAt time.Time `json:"createdAt"`
}
