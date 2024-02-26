package dto

type (
	CategoryDB struct {
		ID           int    `db:"id"`
		CategoryName string `db:"category_name"`
		DefaultDate
	}

	CategoryDataList []CategoryData
	CategoryData     struct {
		ID           int    `json:"id"`
		CategoryName string `json:"category_name"`
	}
)
