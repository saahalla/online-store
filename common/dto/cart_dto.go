package dto

type (
	CartDBList []CartDB
	CartDB     struct {
		ID     int `db:"id" goqu:"skipinsert"`
		UserID int `db:"user_id"`
		DefaultDate
	}

	CartDataList []CartData
	CartData     struct {
		ID     int `json:"id"`
		UserID int `json:"user_id"`
	}

	CartDataListResp []CartDataResp
	CartDataResp     struct {
		Username  string           `json:"username"`
		CartItems CartItemDataList `json:"cart_items"`
	}
)

// req resp

type (
	// get
	GetCartResponse struct {
		DefaultResponse
		Data CartDataResp `json:"data"`
	}

	// list
	ListCartResponse struct {
		DefaultResponse
		Data CartDataListResp `json:"data"`
	}

	// add
	AddCartRequest struct {
		ProductID int `json:"product_id"`
		Qty       int `json:"qty"`
	}

	AddCartResponse struct {
		DefaultResponse
	}

	// delete
	DeleteCartResponse struct {
		DefaultResponse
	}

	// update
	UpdateCartRequest struct {
		ProductID int `json:"product_id"`
		Qty       int `json:"qty"`
	}
	UpdateCartResponse struct {
		DefaultResponse
	}
)
