package auth

import (
	"online-store/common/dto"

	"github.com/gofiber/fiber/v2"
)

func HandlerRegister(s Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		resp := dto.DefaultResponse{}

		err := s.Register(c)
		if err != nil {
			resp.PrepareStatusFailed(err.Error())
			return c.Status(fiber.StatusUnprocessableEntity).JSON(resp)
		}

		resp.PrepareStatusSuccess("success get data")
		return c.JSON(dto.RegisterResponse{
			DefaultResponse: resp,
		})
	}

}

func HandlerLogin(s Service) fiber.Handler {

	return func(c *fiber.Ctx) error {
		resp := dto.DefaultResponse{}

		jwtToken, err := s.Login(c)
		if err != nil {
			resp.PrepareStatusFailed(err.Error())
			return c.Status(fiber.StatusUnprocessableEntity).JSON(resp)
		}

		resp.PrepareStatusSuccess("success get data")
		return c.JSON(dto.LoginResponse{
			DefaultResponse: resp,
			JwtToken:        jwtToken,
		})
	}

}
