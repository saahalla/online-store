package dto

type (
	CheckoutRequest struct {
		CartItems CartItemDataList `json:"cart_items"`
	}

	CheckoutResponse struct {
		DefaultResponse
	}
)
