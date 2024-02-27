package dto

type (
	CategoryDBList []CategoryDB
	CategoryDB     struct {
		ID           int    `db:"id" goqu:"skipinsert"`
		CategoryName string `db:"category_name"`
		DefaultDate
	}

	CategoryDataList []CategoryData
	CategoryData     struct {
		ID           int    `json:"id"`
		CategoryName string `json:"category_name"`
	}

	CategoryDataListResp []CategoryDataResp
	CategoryDataResp     struct {
		ID           int               `json:"id"`
		CategoryName string            `json:"category_name"`
		Products     []ProductDataList `json:"products"`
	}
)

// req resp

type (
	// get
	GetCategoryResponse struct {
		DefaultResponse
		Data CategoryDataResp `json:"data"`
	}

	// list
	ListCategoryResponse struct {
		DefaultResponse
		Data CategoryDataListResp `json:"data"`
	}

	// add
	AddCategoryRequest struct {
		CategoryName string `json:"category_name"`
	}

	AddCategoryResponse struct {
		DefaultResponse
	}

	// delete
	DeleteCategoryResponse struct {
		DefaultResponse
	}

	// update
	UpdateCategoryRequest struct {
		CategoryName string `json:"category_name"`
	}
	UpdateCategoryResponse struct {
		DefaultResponse
	}
)
