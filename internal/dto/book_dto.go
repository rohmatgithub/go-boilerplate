package dto

type BookRequest struct {
	Title         string  `json:"title" validate:"required"`
	Author        string  `json:"author" validate:"required"`
	PublishedDate string  `json:"published_date" validate:"required"`
	CategoryID    int32   `json:"category_id" validate:"required"`
	Price         float64 `json:"price" validate:"required"`
	Stock         int64   `json:"stock" validate:"required,numeric,min=5,max=130"`
}

type BookDetail struct {
	ID            int64  `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedDate string `json:"published_date"`
	CategoryName  string `json:"category_name"`
	Price         string `json:"price"`
	Stock         string `json:"stock"`
}
