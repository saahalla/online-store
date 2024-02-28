package checkouts

import (
	"online-store/common/dto"

	"github.com/gofiber/fiber/v2"
)

func HandlerCheckout(s Service) fiber.Handler {

	return func(c *fiber.Ctx) error {
		resp := dto.DefaultResponse{}

		err := s.Checkout(c)
		if err != nil {
			resp.PrepareStatusFailed(err.Error())
			return c.Status(fiber.StatusUnprocessableEntity).JSON(resp)
		}

		resp.PrepareStatusSuccess("success get data")
		return c.JSON(dto.CheckoutResponse{
			DefaultResponse: resp,
		})
	}

}
