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

func (r *ProductRepo) List(limit, offset, minPrice, maxPrice int, query, sort string) ([]models.Product, int64, error) {
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

	order := "created_at desc"
	switch sort {
	case "price_asc":
		order = "price asc"
	case "price desc":
		order = "price desc"
	case "newest":
		order = "newest"
	case "oldest":
		order = "oldest"
	}

	err := r.db.Limit(limit).Offset(offset).Order(order).Find(&products).Error

	return products, total, err
}

func (r *ProductRepo) FindByID(id uint) (models.Product, error) {
	var product models.Product
	err := r.db.First(&product, id).Error
	return product, err
}

func (r *ProductRepo) Update(product models.Product) (models.Product, error) {
	if err := r.db.Save(&product).Error; err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func (r *ProductRepo) Delete(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}

func (r *ProductRepo) Restore(id uint) (models.Product, error) {
	var product models.Product

	if err := r.db.Unscoped().First(&product).Error; err != nil {
		return models.Product{}, err
	}

	if err := r.db.Unscoped().Model(&product).Update("deleted_at", nil).Error; err != nil {
		return models.Product{}, err
	}

	return product, nil
}
