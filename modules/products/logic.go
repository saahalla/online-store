package products

import (
	"online-store/modules/dto"
	"online-store/modules/repository"

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
	return output, nil
}
func (s *service) List(c *fiber.Ctx) (output dto.ProductDataList, err error) {
	return output, nil
}
