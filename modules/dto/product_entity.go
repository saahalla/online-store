package dto

import (
	"errors"
	"strings"
	"time"
)

func (r *AddProductRequest) Validate() (err error) {
	var errStr []string

	if r.ProductName == "" {
		errStr = append(errStr, "product name is required")
	}

	if r.Stock == nil {
		errStr = append(errStr, "stock is required")
	}

	if r.Price == nil {
		errStr = append(errStr, "price is required")
	}

	if len(errStr) > 0 {
		return errors.New(strings.Join(errStr, ","))
	}

	return nil
}

func (r *AddProductRequest) PrepareDataDB() ProductDB {
	now := time.Now()

	productDB := ProductDB{
		ProductName: r.ProductName,
		Stock:       *r.Stock,
		Price:       *r.Price,
		Image:       r.Image,
	}

	productDB.CreatedAt = now
	productDB.CreatedBy = "test"
	productDB.ModifiedAt = now
	productDB.ModifiedBy = "test"

	return productDB
}
