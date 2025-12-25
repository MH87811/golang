package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model

	UserID     uint
	User       User
	Status     OrderStatus
	TotalPrice uint
	Items      []OrderItems `gorm:"foreignKey:OrderID"`
}

type OrderItems struct {
	gorm.Model

	OrderID   uint
	Order     Order
	ProductID uint
	Product   Product
	Quantity  uint
	Price     uint
}

type OrderStatus string

const (
	OrderPending   OrderStatus = "pending"
	OrderPaid      OrderStatus = "paid"
	OrderCancelled OrderStatus = "cancelled"
)
