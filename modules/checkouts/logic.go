package checkouts

import (
	"fmt"
	"online-store/common/dto"
	"online-store/common/middleware"
	"online-store/common/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type service struct {
	repoCart      repository.CartRepository
	repoCartItems repository.CartItemRepository
	repoProducts  repository.ProductRepository
}

func NewService(db *sqlx.DB) Service {
	return &service{
		repoCart:      repository.NewCartRepo(db),
		repoCartItems: repository.NewCartItemRepo(db),
		repoProducts:  repository.NewProductRepo(db),
	}
}

func (s *service) Checkout(c *fiber.Ctx) error {
	dataBody := new(dto.CheckoutRequest)

	// Parse body into struct
	if err := c.BodyParser(dataBody); err != nil {
		return err
	}

	err := dataBody.Validate()
	if err != nil {
		return err
	}

	// get data user from jwt
	dataJwt, err := middleware.GetDataJWT(c.Locals("user"))
	if err != nil {
		return err
	}

	// get data cart
	cart, err := s.repoCart.Get(repository.ParamSearchGetCart{
		UserID: dataJwt.UserID,
	})

	if err != nil || cart.ID == 0 {
		return fmt.Errorf("cart not found")
	}

	// get cart items
	dataCartItems, err := s.repoCartItems.List(repository.ParamSearchGetCartItemList{
		CartItemID: dataBody.CartItems.GetCartItemIDList(),
	})
	if err != nil {
		return err
	}

	dataProducts, err := s.repoProducts.List(repository.ParamSearchProductList{
		ProductIDList: dataCartItems.GetProductIDList(),
	})

	_ = dataProducts

	// TODO:
	// - reduce stock
	// - insert checkout
	// - count payment

	return nil
}
