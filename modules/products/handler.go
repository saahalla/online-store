package products

import (
	"online-store/modules/dto"

	"github.com/gofiber/fiber/v2"
)

func Get(s Service) fiber.Handler {

	return func(c *fiber.Ctx) error {
		data, err := s.Get(c)
		resp := dto.DefaultResponse{}

		if err != nil {
			resp.PrepareStatusNotFound(err.Error())
			return c.Status(fiber.StatusUnprocessableEntity).JSON(resp)
		}

		resp.PrepareStatusSuccess()
		return c.JSON(dto.GetProductResponse{
			DefaultResponse: resp,
			Data:            data,
		})
	}

}
