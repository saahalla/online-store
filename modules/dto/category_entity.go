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

func (r *AddCategoryRequest) PrepareDataDB() CategoryDB {
	now := time.Now()

	categoryDB := CategoryDB{
		CategoryName: r.CategoryName,
	}

	categoryDB.CreatedAt = now
	categoryDB.CreatedBy = "test"
	categoryDB.ModifiedAt = now
	categoryDB.ModifiedBy = "test"

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

func (r *UpdateCategoryRequest) PrepareDataDB(data *CategoryDB) {

	if r.CategoryName != "" {
		data.CategoryName = r.CategoryName
	}

	data.ModifiedAt = time.Now()

	// if r.CategoryID != nil {
	// 	data.CategoryID = r.CategoryID
	// }
}

// list
func (l CategoryDBList) PrepareDataJSON() (categories CategoryDataList) {

	for _, pDB := range l {

		pData := pDB.ToDataJSON()
		if pData != nil {
			categories = append(categories, *pData)
		}

	}

	return categories
}

func (pDB CategoryDB) ToDataJSON() *CategoryData {

	if pDB.ID != 0 {
		categoryData := CategoryData{
			ID:           pDB.ID,
			CategoryName: pDB.CategoryName,
		}

		return &categoryData
	}

	return nil
}
