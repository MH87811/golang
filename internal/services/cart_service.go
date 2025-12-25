package services

import (
	"errors"
	"shop/internal/models"
	"shop/internal/repositories"

	"gorm.io/gorm"
)

type CartService struct {
	cartRepo    *repositories.CartRepo
	productRepo *repositories.ProductRepo
}

func NewCartService(cartRepo *repositories.CartRepo, productRepo *repositories.ProductRepo) *CartService {
	return &CartService{cartRepo, productRepo}
}

func (s *CartService) Add(userID, productId, qty uint) error {
	if qty == 0 {
		return errors.New("qty must be greater than zero")
	}

	_, err := s.productRepo.FindByID(productId)
	if err != nil {
		return errors.New("product not found")
	}
	cart, err := s.cartRepo.GetOrCreateCart(userID)
	if err != nil {
		return err
	}

	item, err := s.cartRepo.FindItem(cart.ID, productId)
	if err == nil {
		item.Quantity += qty
		return s.cartRepo.UpdateCart(item)
	}

	if err != gorm.ErrRecordNotFound {
		return err
	}

	return s.cartRepo.CreateItem(models.CartItems{
		CartID:    cart.ID,
		ProductID: productId,
		Quantity:  qty,
	})
}

func (s *CartService) UpdateItem(userID, itemId, qty uint) error {
	if qty == 0 {
		return errors.New("qty must be greater than zero")
	}

	cart, err := s.cartRepo.GetOrCreateCart(userID)
	if err != nil {
		return err
	}

	var item models.CartItems
	item, err = s.cartRepo.FindItem(cart.ID, itemId)
	if err != nil {
		return errors.New("item not found")
	}

	item.Quantity = qty
	return s.cartRepo.UpdateCart(item)
}

func (s *CartService) RemoveItem(userID, itemId uint) error {
	_, err := s.cartRepo.GetOrCreateCart(userID)
	if err != nil {
		return err
	}
	return s.cartRepo.DeleteItem(itemId)
}

func (s *CartService) GetCart(userID uint) (models.Cart, error) {
	return s.cartRepo.GetCartWithItems(userID)
}
