package products

import (
	"fmt"
	"online-store/modules/dto"
	"online-store/modules/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type service struct {
	repo         repository.ProductRepository
	repoCategory repository.CategoryRepository
}

func NewService(db *sqlx.DB) Service {
	return &service{
		repo:         repository.NewProductRepo(db),
		repoCategory: repository.NewCategoryRepo(db),
	}
}

func (s *service) Add(c *fiber.Ctx) error {
	dataBody := new(dto.AddProductRequest)

	// Parse body into struct
	if err := c.BodyParser(dataBody); err != nil {
		return err
	}

	err := dataBody.Validate()
	if err != nil {
		return err
	}

	// validate category
	category, err := s.repoCategory.Get(*dataBody.CategoryID)
	if err != nil || category.ID == 0 {
		return fmt.Errorf("category with id %v not found", *dataBody.CategoryID)
	}

	productDB := dataBody.PrepareDataDB()

	err = s.repo.Add(productDB)
	if err != nil {
		return err
	}

	return nil
}
func (s *service) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	dataBody := new(dto.UpdateProductRequest)

	// Parse body into struct
	if err := c.BodyParser(dataBody); err != nil {
		return err
	}

	err := dataBody.Validate()
	if err != nil {
		return err
	}

	productID, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("id must integer")
	}

	// validate
	data, err := s.repo.Get(productID)
	if err != nil || data.ID == 0 {
		return fmt.Errorf("product with id %v not found", id)
	}

	// validate category
	category, err := s.repoCategory.Get(*dataBody.CategoryID)
	if err != nil || category.ID == 0 {
		return fmt.Errorf("category with id %v not found", *dataBody.CategoryID)
	}

	dataBody.PrepareDataDB(&data)

	err = s.repo.Update(productID, data)
	if err != nil {
		return fmt.Errorf("failed to update product with id %v", productID)
	}

	return nil
}
func (s *service) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	productID, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("id must integer")
	}

	// validate
	data, err := s.repo.Get(productID)
	if err != nil || data.ID == 0 {
		return fmt.Errorf("product with id %v not found", id)
	}

	err = s.repo.Delete(productID)
	if err != nil {
		return fmt.Errorf("failed to delete product with id %v", productID)
	}

	return nil
}
func (s *service) Get(c *fiber.Ctx) (output dto.ProductData, err error) {
	id := c.Params("id")

	productID, err := strconv.Atoi(id)
	if err != nil {
		return output, fmt.Errorf("id must integer")
	}

	data, err := s.repo.Get(productID)
	if err != nil {
		return output, err
	}

	category, err := s.repoCategory.Get(data.CategoryID)
	if err != nil || data.ID == 0 {
		return output, fmt.Errorf("category with id %v not found", id)
	}

	productData := data.ToDataJSON(category.ToDataJSON())
	if productData == nil {
		return output, fmt.Errorf("product with id %v not found", id)
	}

	output = *productData

	return output, nil
}
func (s *service) List(c *fiber.Ctx) (output dto.ProductDataList, err error) {

	products, err := s.repo.List()
	if err != nil {
		return output, err
	}

	categories, err := s.repoCategory.List()
	if err != nil {
		return output, err
	}

	output = products.PrepareDataJSON(categories)

	return output, nil
}
