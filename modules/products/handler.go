package products

import (
	"online-store/common/dto"

	"github.com/gofiber/fiber/v2"
)

func Get(s Service) fiber.Handler {

	return func(c *fiber.Ctx) error {
		resp := dto.DefaultResponse{}

		data, err := s.Get(c)
		if err != nil {
			resp.PrepareStatusFailed(err.Error())
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
		resp := dto.DefaultResponse{}

		err := s.Add(c)
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
		resp := dto.DefaultResponse{}

		data, err := s.List(c)
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

func Delete(s Service) fiber.Handler {

	return func(c *fiber.Ctx) error {
		resp := dto.DefaultResponse{}

		err := s.Delete(c)
		if err != nil {
			resp.PrepareStatusFailed(err.Error())
			return c.Status(fiber.StatusUnprocessableEntity).JSON(resp)
		}

		resp.PrepareStatusSuccess("success delete product")
		return c.JSON(dto.DeleteProductResponse{
			DefaultResponse: resp,
		})
	}

}

func Update(s Service) fiber.Handler {

	return func(c *fiber.Ctx) error {
		resp := dto.DefaultResponse{}

		err := s.Update(c)
		if err != nil {
			resp.PrepareStatusFailed(err.Error())
			return c.Status(fiber.StatusUnprocessableEntity).JSON(resp)
		}

		resp.PrepareStatusSuccess("success update product")
		return c.JSON(dto.UpdateProductResponse{
			DefaultResponse: resp,
		})
	}

}
