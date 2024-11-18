package repository

import (
	"boilerplate/internal/domain"

	"gorm.io/gorm"
)

type BookRepository interface {
	GetDetailByID(id int64) (*domain.Book, error)
	Insert(book *domain.Book) error
	Update(book *domain.Book) error
	Delete(book *domain.Book) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) GetDetailByID(id int64) (result *domain.Book, err error) {
	err = r.db.Where("id = ? ", id).Find(&result).Error
	return
}

func (r *bookRepository) Insert(book *domain.Book) error {
	err := r.db.Create(book).Error
	return err
}

func (r *bookRepository) Update(book *domain.Book) error {
	err := r.db.Save(book).Error
	return err
}

func (r *bookRepository) Delete(book *domain.Book) error {
	err := r.db.Delete(book).Error
	return err
}
