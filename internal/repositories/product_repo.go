package repositories

import (
	"shop/internal/models"

	"gorm.io/gorm"
)

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{db: db}
}

func (r *ProductRepo) Create(product models.Product) (models.Product, error) {
	err := r.db.Create(&product).Error
	return product, err
}

func (r *ProductRepo) List(limit, offset, minPrice, maxPrice int, query string) ([]models.Product, int64, error) {
	var products []models.Product
	var total int64

	db := r.db.Model(&models.Product{})

	if minPrice > 0 {
		db = db.Where("price >= ?", minPrice)
	}
	if maxPrice > 0 {
		db = db.Where("price <= ?", maxPrice)
	}
	if query != "" {
		db = db.Where("name LIKE ?", "%"+query+"%")
	}

	db.Count(&total)

	err := r.db.Limit(limit).Offset(offset).Order("created_at").Find(&products).Error

	return products, total, err
}

func (r *ProductRepo) FindByID(id uint) (models.Product, error) {
	var product models.Product
	err := r.db.First(&product, id).Error
	return product, err
}

func (r *ProductRepo) Update(product models.Product) (models.Product, error) {
	return product, r.db.Save(&product).Error
}

func (r *ProductRepo) Delete(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}
