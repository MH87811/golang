package services

import (
	"errors"
	"shop/internal/dto"
	"shop/internal/models"
	"shop/internal/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepo
}

func NewProductService(repo *repositories.ProductRepo) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) Create(product models.Product) (models.Product, error) {
	if product.Price <= 0 {
		return models.Product{}, errors.New("invalid price")
	}
	return s.repo.Create(product)
}

func (s *ProductService) List(limit, page, minPrice, maxPrice int, query, sort string) ([]models.Product, int64, error) {
	offset := (page - 1) * limit
	return s.repo.List(limit, offset, minPrice, maxPrice, query, sort)
}

func (s *ProductService) Update(productID, userID uint, req dto.UpdateProductRequest) (models.Product, error) {
	product, err := s.repo.FindByID(productID)
	if err != nil {
		return models.Product{}, err
	}

	if product.UserID != userID {
		return models.Product{}, errors.New("forbidden")
	}

	if req.Name != nil {
		product.Name = *req.Name
	}
	if req.Price != nil {
		if *req.Price <= 0 {
			return models.Product{}, errors.New("invalid price")
		}
		product.Price = *req.Price
	}
	if req.Stock != nil {
		if *req.Stock < 0 {
			return models.Product{}, errors.New("invalid stock")
		}
		product.Stock = *req.Stock
	}

	return s.repo.Update(product)
}

func (s *ProductService) Delete(productID, userID uint) error {
	product, err := s.repo.FindByID(productID)
	if err != nil {
		return errors.New("product not found")
	}

	if product.UserID != userID {
		return errors.New("forbidden")
	}

	return s.repo.Delete(productID)
}

func (s *ProductService) Restore(productID, userID uint) (models.Product, error) {
	product, err := s.repo.FindUnscopedByID(productID)
	if err != nil {
		return models.Product{}, errors.New("product not fond")
	}
	if product.UserID != userID {
		return models.Product{}, errors.New("forbidden")
	}
	return s.repo.Restore(productID)
}
