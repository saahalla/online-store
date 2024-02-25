package dto

type (
	CategoryDB struct {
		ID           int    `db:"id"`
		CategoryName string `db:"category_name"`
		DefaultDate
	}

	CategoryResponseList []CategoryResponse
	CategoryResponse     struct {
		ID           int    `json:"id"`
		CategoryName string `json:"category_name"`
	}
)
