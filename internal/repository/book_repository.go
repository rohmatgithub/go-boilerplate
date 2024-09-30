package repository

import (
	"boilerplate/internal/domain"

	"gorm.io/gorm"
)

type BookRepository interface {
	GetExampleByID(id int64) (*domain.Book, error)
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) GetExampleByID(id int64) (result *domain.Book, err error) {
	err = r.db.Where("id = ? ", id).Find(&result).Error
	return
}
