package dto

import "fmt"

func (r *CheckoutRequest) Validate() error {

	if len(r.CartItems) == 0 {
		return fmt.Errorf("cart items is required")
	}

	return nil
}
