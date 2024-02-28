package carts

import (
	"online-store/common/dto"

	"github.com/gofiber/fiber/v2"
)

func HandlerGet(s Service) fiber.Handler {

	return func(c *fiber.Ctx) error {
		resp := dto.DefaultResponse{}

		data, err := s.Get(c)
		if err != nil {
			resp.PrepareStatusFailed(err.Error())
			return c.Status(fiber.StatusUnprocessableEntity).JSON(resp)
		}

		resp.PrepareStatusSuccess("success get data")
		return c.JSON(dto.GetCartResponse{
			DefaultResponse: resp,
			Data:            data,
		})
	}

}

func HandlerAdd(s Service) fiber.Handler {

	return func(c *fiber.Ctx) error {
		resp := dto.DefaultResponse{}

		err := s.Add(c)
		if err != nil {
			resp.PrepareStatusFailed(err.Error())
			return c.Status(fiber.StatusUnprocessableEntity).JSON(resp)
		}

		resp.PrepareStatusSuccess("success add cart")
		return c.JSON(dto.AddCartResponse{
			DefaultResponse: resp,
		})
	}

}

// func HandlerList(s Service) fiber.Handler {

// 	return func(c *fiber.Ctx) error {
// 		resp := dto.DefaultResponse{}

// 		data, err := s.List(c)
// 		if err != nil {
// 			resp.PrepareStatusFailed(err.Error())
// 			return c.Status(fiber.StatusUnprocessableEntity).JSON(resp)
// 		}

// 		resp.PrepareStatusSuccess("success get list")
// 		return c.JSON(dto.ListCartResponse{
// 			DefaultResponse: resp,
// 			Data:            data,
// 		})
// 	}

// }

// func HandlerDelete(s Service) fiber.Handler {

// 	return func(c *fiber.Ctx) error {
// 		resp := dto.DefaultResponse{}

// 		err := s.Delete(c)
// 		if err != nil {
// 			resp.PrepareStatusFailed(err.Error())
// 			return c.Status(fiber.StatusUnprocessableEntity).JSON(resp)
// 		}

// 		resp.PrepareStatusSuccess("success delete cart")
// 		return c.JSON(dto.DeleteCartResponse{
// 			DefaultResponse: resp,
// 		})
// 	}

// }

func HandlerUpdate(s Service) fiber.Handler {

	return func(c *fiber.Ctx) error {
		resp := dto.DefaultResponse{}

		err := s.Update(c)
		if err != nil {
			resp.PrepareStatusFailed(err.Error())
			return c.Status(fiber.StatusUnprocessableEntity).JSON(resp)
		}

		resp.PrepareStatusSuccess("success update cart")
		return c.JSON(dto.UpdateCartResponse{
			DefaultResponse: resp,
		})
	}

}
