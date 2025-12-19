package services

import (
	"errors"
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

func (s *ProductService) List() ([]models.Product, error) {
	return s.repo.GetAll()
}
