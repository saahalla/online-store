package dto

import (
	"errors"
	"strings"
	"time"
)

// add
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

// list
func (l ProductDBList) PrepareDataJSON() (products ProductDataList) {

	for _, pDB := range l {

		pData := pDB.ToDataJSON()
		if pData != nil {
			products = append(products, *pData)
		}

	}

	return products
}

func (pDB ProductDB) ToDataJSON() *ProductData {

	if pDB.ID != 0 {
		productData := ProductData{
			ID:          pDB.ID,
			ProductName: pDB.ProductName,
			Stock:       pDB.Stock,
			Price:       pDB.Price,
			Image:       pDB.Image,
		}

		return &productData
	}

	return nil
}

// update
func (r *UpdateProductRequest) Validate() (err error) {
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

func (r *UpdateProductRequest) PrepareDataDB(data *ProductDB) {

	if r.ProductName != "" {
		data.ProductName = r.ProductName
	}

	if r.Stock != nil {
		data.Stock = *r.Stock
	}

	if r.Price != nil {
		data.Price = *r.Price
	}

	if r.Image != "" {
		data.Image = r.Image
	}

	data.ModifiedAt = time.Now()

	// if r.CategoryID != nil {
	// 	data.CategoryID = r.CategoryID
	// }
}
