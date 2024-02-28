package auth

import (
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	Register(c *fiber.Ctx) (err error)
	Login(c *fiber.Ctx) (jwtToken string, err error)
}
