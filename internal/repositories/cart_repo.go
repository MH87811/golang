package repositories

import (
	"gorm.io/gorm"
	"shop/internal/models"
)

type CartRepo struct {
	db *gorm.DB
}

func NewCartRepo(db *gorm.DB) *CartRepo {
	return &CartRepo{db: db}
}

func (r *CartRepo) GetOrCreateCart(userID uint) (models.Cart, error) {
	var cart models.Cart

	err := r.db.Preload("Items.Product").Where("user_id = ?", userID).First(&cart).Error
	if err == nil {
		return cart, nil
	}

	if err == gorm.ErrRecordNotFound {
		cart = models.Cart{UserID: userID}
		err = r.db.Create(&cart).Error
		return cart, err
	}

	return models.Cart{}, err
}

func (r *CartRepo) FindItem(cartID, productID uint) (models.CartItems, error) {
	var item models.CartItems
	err := r.db.Where("cart_id = ? AND product_id = ?", cartID, productID).First(&item).Error
	return item, err
}

func (r *CartRepo) CreateItem(item models.CartItems) error {
	return r.db.Create(&item).Error
}

func (r *CartRepo) UpdateCart(item models.CartItems) error {
	return r.db.Save(&item).Error
}

func (r *CartRepo) DeleteItem(itemID uint) error {
	return r.db.Delete(&models.CartItems{}, itemID).Error
}

func (r *CartRepo) GetCartWithItems(userID uint) (models.Cart, error) {
	var cart models.Cart
	err := r.db.Preload("Items.Product").Where("user_id = ?", userID).First(&cart).Error
	return cart, err
}
