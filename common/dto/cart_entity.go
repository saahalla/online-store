package dto

import (
	"errors"
	"strings"
	"time"
)

func (r *AddCartRequest) Validate() (err error) {
	var errStr []string

	if r.ProductID == 0 {
		errStr = append(errStr, "category id is required")
	}

	if r.Qty == 0 {
		errStr = append(errStr, "qty is required")
	}

	if len(errStr) > 0 {
		return errors.New(strings.Join(errStr, ","))
	}

	return nil
}

func (r *AddCartRequest) ToCartItemDB(cartID int, username string) CartItemDB {
	now := time.Now()

	cartItemDB := CartItemDB{
		CartID:    cartID,
		ProductID: r.ProductID,
		Qty:       r.Qty,
	}

	cartItemDB.CreatedAt = now
	cartItemDB.CreatedBy = username
	cartItemDB.ModifiedAt = now
	cartItemDB.ModifiedBy = username

	return cartItemDB
}

// update
func (r *UpdateCartRequest) Validate() (err error) {
	var errStr []string

	if r.ProductID == 0 {
		errStr = append(errStr, "category id is required")
	}

	if r.Qty == 0 {
		errStr = append(errStr, "qty is required")
	}

	if len(errStr) > 0 {
		return errors.New(strings.Join(errStr, ","))
	}

	return nil
}

func (r *UpdateCartRequest) PrepareDataDB(data *CartItemDB, username string) {

	if r.Qty != 0 {
		data.Qty = r.Qty
	}

	data.ModifiedAt = time.Now()
	data.ModifiedBy = username
}

func (c *CartDB) PrepareDataJSON(cartItems CartItemDataList, username string) CartDataResp {
	output := CartDataResp{
		ID:        c.ID,
		Username:  username,
		CartItems: cartItems,
	}

	return output
}
