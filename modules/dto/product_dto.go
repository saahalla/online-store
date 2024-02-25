package dto

type (
	ProductDB struct {
		ID          int     `db:"id"`
		ProductName string  `db:"product_name"`
		Stock       int     `db:"stock"`
		Price       float64 `db:"price"`
		Image       string  `db:"image"`
		DefaultDate
	}

	ProductResponseList []ProductResponse
	ProductResponse     struct {
		ID          int                  `json:"id"`
		ProductName string               `json:"product_name"`
		Stock       int                  `json:"stock"`
		Price       float64              `json:"price"`
		Image       string               `json:"image"`
		Categories  CategoryResponseList `json:"categories"`
	}
)
