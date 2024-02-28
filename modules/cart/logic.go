package carts

import (
	"fmt"
	"online-store/common/dto"
	"online-store/common/middleware"
	"online-store/common/repository"
	"time"

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

func (s *service) Add(c *fiber.Ctx) error {
	dataBody := new(dto.AddCartRequest)

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

	dataProduct, err := s.repoProducts.Get(dataBody.ProductID)
	if err != nil {
		return err
	}

	if dataProduct.ID == 0 {
		return fmt.Errorf("product with id %v not found", dataBody.ProductID)
	}

	if dataBody.Qty > dataProduct.Stock {
		return fmt.Errorf("out of stock product %v", dataProduct.ProductName)
	}

	// get data cart
	cart, err := s.repoCart.Get(repository.ParamSearchGetCart{
		UserID: dataJwt.UserID,
	})

	if err != nil {
		return err
	}

	if cart.ID == 0 {
		// insert new cart
		cartDB := dto.CartDB{
			UserID: dataJwt.UserID,
		}
		cartDB.CreatedAt = time.Now()
		cartDB.ModifiedAt = time.Now()
		cartDB.CreatedBy = dataJwt.Username
		cartDB.ModifiedBy = dataJwt.Username

		err := s.repoCart.Add(cartDB)
		if err != nil {
			return err
		}

		cart, err = s.repoCart.Get(repository.ParamSearchGetCart{
			UserID: dataJwt.UserID,
		})
	}

	// get cart items
	dataCartItem, err := s.repoCartItems.Get(repository.ParamSearchGetCartItem{
		CartID:    cart.ID,
		ProductID: dataBody.ProductID,
	})

	if dataCartItem.ID == 0 {

		cartItemDB := dataBody.ToCartItemDB(cart.ID, dataJwt.Username)

		err = s.repoCartItems.Add(cartItemDB)
		if err != nil {
			return err
		}

	} else {

		dataCartItem.Qty = dataCartItem.Qty + dataBody.Qty
		dataCartItem.ModifiedAt = time.Now()
		dataCartItem.ModifiedBy = dataJwt.Username

		if dataCartItem.Qty > dataProduct.Stock {
			return fmt.Errorf("out of stock product %v", dataProduct.ProductName)
		}

		err := s.repoCartItems.Update(dataCartItem.ID, dataCartItem)
		if err != nil {
			return err
		}

	}

	return nil
}

func (s *service) Update(c *fiber.Ctx) error {

	dataBody := new(dto.UpdateCartRequest)

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

	// validate
	dataCart, err := s.repoCart.Get(repository.ParamSearchGetCart{
		UserID: dataJwt.UserID,
	})
	if err != nil || dataCart.ID == 0 {
		return fmt.Errorf("cart not found")
	}

	dataProduct, err := s.repoProducts.Get(dataBody.ProductID)
	if err != nil || dataProduct.ID == 0 {
		return fmt.Errorf("product not found")
	}

	dataCartItem, err := s.repoCartItems.Get(repository.ParamSearchGetCartItem{
		CartID:    dataCart.ID,
		ProductID: dataBody.ProductID,
	})

	dataBody.PrepareDataDB(&dataCartItem, dataJwt.Username)

	if dataBody.Qty > 0 {
		if dataBody.Qty > dataProduct.Stock {
			return fmt.Errorf("out of stock product %v", dataProduct.ProductName)
		}

		// update data cart items qty
		err = s.repoCartItems.Update(dataCartItem.ID, dataCartItem)
		if err != nil {
			return fmt.Errorf("failed to update cart")
		}

	} else {

		// delete cart items
		err = s.repoCartItems.Delete(dataCartItem.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

// func (s *service) Delete(c *fiber.Ctx) error {
// 	id := c.Params("id")

// 	cartID, err := strconv.Atoi(id)
// 	if err != nil {
// 		return fmt.Errorf("id must integer")
// 	}

// 	// validate
// 	data, err := s.repoCart.Get(cartID)
// 	if err != nil || data.ID == 0 {
// 		return fmt.Errorf("cart with id %v not found", id)
// 	}

// 	err = s.repoCart.Delete(cartID)
// 	if err != nil {
// 		return fmt.Errorf("failed to delete cart with id %v", cartID)
// 	}

// 	return nil
// }

func (s *service) Get(c *fiber.Ctx) (output dto.CartDataResp, err error) {

	// get data user from jwt
	dataJwt, err := middleware.GetDataJWT(c.Locals("user"))
	if err != nil {
		return output, err
	}

	// get cart
	data, err := s.repoCart.Get(repository.ParamSearchGetCart{
		UserID: dataJwt.UserID,
	})
	if err != nil {
		return output, err
	}

	if data.ID == 0 {
		return output, fmt.Errorf("there is no data on cart")
	}

	// get cart items
	cartItems, err := s.repoCartItems.DetailList(repository.ParamSearchDetailCartItem{
		CartID: data.ID,
	})
	if err != nil {
		return output, err
	}

	cartData := data.PrepareDataJSON(cartItems, dataJwt.Username)

	output = cartData

	return output, nil
}

// func (s *service) List(c *fiber.Ctx) (output dto.CartDataListResp, err error) {

// 	carts, err := s.repoCart.List()
// 	if err != nil {
// 		return output, err
// 	}

// 	// get all products
// 	products, err := s.repoProducts.List(repository.ParamSearchProductList{})
// 	if err != nil {
// 		return output, err
// 	}

// 	output = carts.PrepareDataJSON(products)

// 	return output, nil
// }
