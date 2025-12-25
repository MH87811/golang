package services

import (
	"errors"
	"shop/internal/models"
	"shop/internal/repositories"

	"gorm.io/gorm"
)

type OrderService struct {
	db          *gorm.DB
	orderRepo   *repositories.OrderRepo
	cartRepo    *repositories.CartRepo
	productRepo *repositories.ProductRepo
}

func NewOrderService(db *gorm.DB, orderRepo *repositories.OrderRepo, cartRepo *repositories.CartRepo, productRepo *repositories.ProductRepo) *OrderService {
	return &OrderService{db, orderRepo, cartRepo, productRepo}
}

func (s *OrderService) CreateFromCart(userID uint) (models.Order, error) {
	cart, err := s.cartRepo.GetOrCreateCart(userID)
	if err != nil {
		return models.Order{}, err
	}
	if len(cart.Items) == 0 {
		return models.Order{}, errors.New("no items in cart")
	}

	var order models.Order

	err = s.db.Transaction(func(tx *gorm.DB) error {
		total := uint(0)

		order = models.Order{
			UserID: userID,
			Status: models.OrderPending,
		}

		if err := s.orderRepo.Create(tx, &order); err != nil {
			return err
		}

		for _, item := range cart.Items {
			if item.Quantity > item.Product.Stock {
				return errors.New("not enough stock")
			}

			price := item.Product.Price
			total += price * item.Quantity

			orderItems := models.OrderItems{
				OrderID:   order.ID,
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
				Price:     price,
			}

			if err := tx.Create(&orderItems).Error; err != nil {
				return err
			}

			if err := tx.Model(&models.Product{}).
				Where("id = ?", item.ProductID).
				Update("stock", gorm.Expr("stock - ?", item.Quantity)).Error; err != nil {
				return err
			}
		}

		if err := tx.Model(&order).
			Update("total_price", total).Error; err != nil {
			return err
		}

		if err := tx.Where("cart_id = ?", cart.ID).Delete(&models.CartItems{}).Error; err != nil {
			return err
		}

		return nil
	})
	return order, err
}

func (s *OrderService) ListByUser(userID uint) ([]models.Order, error) {
	return s.orderRepo.ListByUser(userID)
}
