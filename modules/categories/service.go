package categories

import (
	"online-store/modules/dto"

	"github.com/gofiber/fiber/v2"
)

type Service interface {
	Add(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	Get(c *fiber.Ctx) (dto.CategoryData, error)
	List(c *fiber.Ctx) (dto.CategoryDataList, error)
}
