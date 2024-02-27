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
	repo repository.ProductRepository
}

func NewService(db *sqlx.DB) Service {
	return &service{
		repo: repository.NewProductRepo(db),
	}
}

func (s *service) Add(c *fiber.Ctx) error {
	dataBody := new(dto.AddProductRequest)

	// Parse body into struct
	if err := c.BodyParser(dataBody); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	err := dataBody.Validate()
	if err != nil {
		return err
	}

	productDB := dataBody.PrepareDataDB()

	err = s.repo.Add(productDB)
	if err != nil {
		return err
	}

	return nil
}
func (s *service) Update(c *fiber.Ctx) error {
	return nil
}
func (s *service) Delete(c *fiber.Ctx) error {
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

	productData := data.ToDataJSON()
	if productData == nil {
		return output, fmt.Errorf("product with id %v not found", id)
	}

	return output, nil
}
func (s *service) List(c *fiber.Ctx) (output dto.ProductDataList, err error) {

	products, err := s.repo.List()
	if err != nil {
		return output, err
	}

	output = products.PrepareDataJSON()

	return output, nil
}
