package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID uint        `json:"user_id"`
	User   User        `gorm:"foreignKey:UserID"`
	Items  []CartItems `gorm:"foreignKey:CartID"`
}

type CartItems struct {
	gorm.Model
	CartID    uint `json:"cart_id" gorm:"idx_cart_product"`
	ProductID uint `json:"product_id" gorm:"idx_cart_product"`
	Product   Product
	Quantity  uint `json:"quantity"`
}
