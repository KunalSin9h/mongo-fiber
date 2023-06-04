package models

type User struct {
	Image       string `json:"image"` // Image url from object store like Amazon S3
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone"`
	Password    string `json:"password"`
	Role        string `json:"role"`
	Note        string `json:"note"`
}

type Item struct {
	Image       string  `json:"image"` // Image url from object store like Amazon S3
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	SubCategory string  `json:"sub_category"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	Supplier    string  `json:"supplier"`
	MinStock    int     `json:"min_stock"`
	Note        string  `json:"note"`
}

type SalesRecord struct {
	CustomerName  string  `json:"name"`
	CustomerEmail string  `json:"email"`
	CustomerPhone string  `json:"phone"`
	Date          string  `json:"date"`
	Time          string  `json:"time"`
	Product       string  `json:"product"`
	Quantity      int     `json:"quantity"`
	Price         float64 `json:"price"`
	Note          string  `json:"note"`
}
