package checkouts

import (
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	Checkout(c *fiber.Ctx) error
}
