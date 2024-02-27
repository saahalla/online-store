package products

import (
	"online-store/modules/dto"

	"github.com/gofiber/fiber/v2"
)

type Service interface {
	Add(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	Get(c *fiber.Ctx) (dto.ProductDataResp, error)
	List(c *fiber.Ctx) (dto.ProductDataListResp, error)
}
