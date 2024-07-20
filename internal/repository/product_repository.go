package repository

import (
	"boilerplate/internal/domain"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProductByID(companyID, productCode string) (*domain.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetProductByID(companyID, productCode string) (result *domain.Product, err error) {
	err = r.db.Where("companyID = ? AND productCode = ?", companyID, productCode).Find(&result).Error
	return
}
