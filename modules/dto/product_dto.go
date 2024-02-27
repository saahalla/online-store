package dto

type (
	ProductDBList []ProductDB
	ProductDB     struct {
		ID          int     `db:"id" goqu:"skipinsert"`
		ProductName string  `db:"product_name"`
		Stock       int     `db:"stock"`
		Price       float64 `db:"price"`
		Image       string  `db:"image"`
		CategoryID  int     `db:"category_id"`
		DefaultDate
	}

	ProductDataList []ProductData
	ProductData     struct {
		ID          int          `json:"id"`
		ProductName string       `json:"product_name"`
		Stock       int          `json:"stock"`
		Price       float64      `json:"price"`
		Image       string       `json:"image"`
		Category    CategoryData `json:"category"`
	}
)

// req resp

type (
	// get
	GetProductResponse struct {
		DefaultResponse
		Data ProductData `json:"data"`
	}

	// list
	ListProductResponse struct {
		DefaultResponse
		Data ProductDataList `json:"data"`
	}

	// add
	AddProductRequest struct {
		ProductName string   `json:"product_name"`
		Stock       *int     `json:"stock"`
		Price       *float64 `json:"price"`
		Image       string   `json:"image"`
		CategoryID  *int     `json:"category_id"`
	}

	AddProductResponse struct {
		DefaultResponse
	}

	// delete
	DeleteProductResponse struct {
		DefaultResponse
	}

	// update
	UpdateProductRequest struct {
		ProductName string   `json:"product_name"`
		Stock       *int     `json:"stock"`
		Price       *float64 `json:"price"`
		Image       string   `json:"image"`
		CategoryID  *int     `json:"category_id"`
	}
	UpdateProductResponse struct {
		DefaultResponse
	}
)
