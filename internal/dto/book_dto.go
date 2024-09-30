package dto

type BookRequest struct {
	ID    string `json:"id"`
	Title string `json:"title"`
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
