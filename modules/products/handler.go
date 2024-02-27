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

		resp.PrepareStatusSuccess("success get data")
		return c.JSON(dto.GetProductResponse{
			DefaultResponse: resp,
			Data:            data,
		})
	}

}

func Add(s Service) fiber.Handler {

	return func(c *fiber.Ctx) error {
		err := s.Add(c)
		resp := dto.DefaultResponse{}

		if err != nil {
			resp.PrepareStatusFailed(err.Error())
			return c.Status(fiber.StatusUnprocessableEntity).JSON(resp)
		}

		resp.PrepareStatusSuccess("success add product")
		return c.JSON(dto.AddProductResponse{
			DefaultResponse: resp,
		})
	}

}

func List(s Service) fiber.Handler {

	return func(c *fiber.Ctx) error {
		data, err := s.List(c)
		resp := dto.DefaultResponse{}

		if err != nil {
			resp.PrepareStatusFailed(err.Error())
			return c.Status(fiber.StatusUnprocessableEntity).JSON(resp)
		}

		resp.PrepareStatusSuccess("success get list")
		return c.JSON(dto.ListProductResponse{
			DefaultResponse: resp,
			Data:            data,
		})
	}

}
