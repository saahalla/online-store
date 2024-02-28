package dto

type (
	CartItemDBList []CartDB
	CartItemDB     struct {
		ID        int `db:"id" goqu:"skipinsert"`
		CartID    int `db:"cart_id"`
		ProductID int `db:"product_id"`
		Qty       int `db:"qty"`
		DefaultDate
	}

	CartItemDataList []CartItemData
	CartItemData     struct {
		ID          int    `db:"id" json:"id"`
		ProductName string `db:"product_name" json:"product_name"`
		Qty         int    `db:"qty" json:"qty"`
	}
)

// req resp
