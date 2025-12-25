package repositories

import (
	"shop/internal/models"

	"gorm.io/gorm"
)

type OrderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) *OrderRepo {
	return &OrderRepo{db: db}
}

func (r *OrderRepo) Create(tx *gorm.DB, order *models.Order) error {
	return tx.Create(order).Error
}

func (r *OrderRepo) ListByUser(userID uint) ([]models.Order, error) {
	var order []models.Order
	err := r.db.
		Preload("Items.Product").
		Where("user_id = ?", userID).
		Order("created_at desc").
		Find(&order).Error

	return order, err
}

func (r *OrderRepo) FindByID(id uint) (*models.Order, error) {
	var order *models.Order
	err := r.db.Preload("Items.Product").First(&order, id).Error
	return order, err
}
