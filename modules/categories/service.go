package categories

import (
	"online-store/modules/dto"

	"github.com/gofiber/fiber/v2"
)

type Service interface {
	Add(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	Get(c *fiber.Ctx) (dto.CategoryDataResp, error)
	List(c *fiber.Ctx) (dto.CategoryDataListResp, error)
}
