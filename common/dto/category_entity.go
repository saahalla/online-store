package dto

import (
	"errors"
	"strings"
	"time"
)

func (r *AddCategoryRequest) Validate() (err error) {
	var errStr []string

	if r.CategoryName == "" {
		errStr = append(errStr, "category name is required")
	}

	if len(errStr) > 0 {
		return errors.New(strings.Join(errStr, ","))
	}

	return nil
}

func (r *AddCategoryRequest) PrepareDataDB(username string) CategoryDB {
	now := time.Now()

	categoryDB := CategoryDB{
		CategoryName: r.CategoryName,
	}

	categoryDB.CreatedAt = now
	categoryDB.CreatedBy = username
	categoryDB.ModifiedAt = now
	categoryDB.ModifiedBy = username

	return categoryDB
}

// update
func (r *UpdateCategoryRequest) Validate() (err error) {
	var errStr []string

	if r.CategoryName == "" {
		errStr = append(errStr, "category name is required")
	}

	if len(errStr) > 0 {
		return errors.New(strings.Join(errStr, ","))
	}

	return nil
}

func (r *UpdateCategoryRequest) PrepareDataDB(data *CategoryDB, username string) {

	if r.CategoryName != "" {
		data.CategoryName = r.CategoryName
	}

	data.ModifiedAt = time.Now()
	data.ModifiedBy = username
}

// list
func (l CategoryDBList) PrepareDataJSON(products ProductDBList) (categories CategoryDataListResp) {

	mapProductsByCategoryID := products.ToDataMapByCategoryID()

	for _, catDB := range l {

		catData := catDB.ToDataJSON(mapProductsByCategoryID[catDB.ID])
		if catData != nil {
			categories = append(categories, *catData)
		}

	}

	return categories
}

func (pDB CategoryDB) ToDataJSON(products ProductDBList) *CategoryDataResp {

	if pDB.ID != 0 {
		categoryData := CategoryDataResp{
			ID:           pDB.ID,
			CategoryName: pDB.CategoryName,
		}

		if len(products) > 0 {
			categoryData.Products = append(categoryData.Products, products.PrepareData())
		}

		return &categoryData
	}

	return nil
}

func (pDB CategoryDB) ToData() *CategoryData {

	if pDB.ID != 0 {
		categoryData := CategoryData{
			ID:           pDB.ID,
			CategoryName: pDB.CategoryName,
		}

		return &categoryData
	}

	return nil
}

func (l CategoryDBList) ToDataMapByCategoryID() map[int]CategoryDB {

	output := map[int]CategoryDB{}

	for _, c := range l {
		output[c.ID] = c
	}

	return output
}
