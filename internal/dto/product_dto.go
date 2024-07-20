package dto

type ProductRequest struct {
	CompanyID   string `json:"company_id"`
	ProductCode string `json:"product_code"`
}

type ProductDetail struct {
	ProductCode string `json:"product_code"`
	ProductName string `json:"product_name"`
}
