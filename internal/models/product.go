package models

import "time"

type Product struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Price     uint      `json:"price" gorm:"not null"`
	Stock     int       `json:"stock" gorm:"not null"`
	UserID    uint      `json:"userId"`
	User      User      `gorm:"foreignKey:UserID"`
	CreatedAt time.Time `json:"createdAt"`
}
