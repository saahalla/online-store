package products

import (
	"online-store/modules/dto"

	"github.com/gofiber/fiber/v2"
)

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) Add(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"add": true})
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
