package categories

import (
	"fmt"
	"online-store/modules/dto"
	"online-store/modules/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type service struct {
	repo repository.CategoryRepository
}

func NewService(db *sqlx.DB) Service {
	return &service{
		repo: repository.NewCategoryRepo(db),
	}
}

func (s *service) Add(c *fiber.Ctx) error {
	dataBody := new(dto.AddCategoryRequest)

	// Parse body into struct
	if err := c.BodyParser(dataBody); err != nil {
		return err
	}

	err := dataBody.Validate()
	if err != nil {
		return err
	}

	categoryDB := dataBody.PrepareDataDB()

	err = s.repo.Add(categoryDB)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	dataBody := new(dto.UpdateCategoryRequest)

	// Parse body into struct
	if err := c.BodyParser(dataBody); err != nil {
		return err
	}

	err := dataBody.Validate()
	if err != nil {
		return err
	}

	categoryID, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("id must integer")
	}

	// validate
	data, err := s.repo.Get(categoryID)
	if err != nil || data.ID == 0 {
		return fmt.Errorf("category with id %v not found", id)
	}

	dataBody.PrepareDataDB(&data)

	err = s.repo.Update(categoryID, data)
	if err != nil {
		return fmt.Errorf("failed to update category with id %v", categoryID)
	}

	return nil
}

func (s *service) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	categoryID, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("id must integer")
	}

	// validate
	data, err := s.repo.Get(categoryID)
	if err != nil || data.ID == 0 {
		return fmt.Errorf("category with id %v not found", id)
	}

	err = s.repo.Delete(categoryID)
	if err != nil {
		return fmt.Errorf("failed to delete category with id %v", categoryID)
	}

	return nil
}

func (s *service) Get(c *fiber.Ctx) (output dto.CategoryData, err error) {
	id := c.Params("id")

	categoryID, err := strconv.Atoi(id)
	if err != nil {
		return output, fmt.Errorf("id must integer")
	}

	data, err := s.repo.Get(categoryID)
	if err != nil {
		return output, err
	}

	categoriesData := data.ToDataJSON()
	if categoriesData == nil {
		return output, fmt.Errorf("category with id %v not found", id)
	}

	output = *categoriesData

	return output, nil
}

func (s *service) List(c *fiber.Ctx) (output dto.CategoryDataList, err error) {

	categories, err := s.repo.List()
	if err != nil {
		return output, err
	}

	output = categories.PrepareDataJSON()

	return output, nil
}
