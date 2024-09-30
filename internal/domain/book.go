package domain

import "database/sql"

type Book struct {
	ID            int64
	Title         string
	Author        string
	PublishedDate sql.NullTime
	CategoryID    int32
	Price         float64
	Stock         int64
}

func (Book) TableName() string {
	return "book"
}

type BookCategory struct {
	ID   int64
	Name string
}

func (BookCategory) TableName() string {
	return "book_category"
}
