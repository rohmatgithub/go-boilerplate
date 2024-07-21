package repository

import (
	"boilerplate/internal/domain"

	"gorm.io/gorm"
)

type ExampleRepository interface {
	GetExampleByID(companyID, ExampleCode string) (*domain.Example, error)
}

type exampleRepository struct {
	db *gorm.DB
}

func NewExampleRepository(db *gorm.DB) ExampleRepository {
	return &exampleRepository{db: db}
}

func (r *exampleRepository) GetExampleByID(companyID, ExampleCode string) (result *domain.Example, err error) {
	err = r.db.Where("companyID = ? AND ExampleCode = ?", companyID, ExampleCode).Find(&result).Error
	return
}
