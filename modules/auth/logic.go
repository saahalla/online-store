package auth

import (
	"fmt"
	"online-store/common/dto"
	"online-store/common/repository"
	"online-store/config"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
)

type service struct {
	repo      repository.AuthRepository
	repoUsers repository.UserRepository
}

func NewService(db *sqlx.DB) Service {
	return &service{
		repo:      repository.NewAuthRepo(db),
		repoUsers: repository.NewUserRepo(db),
	}
}

func (s *service) Register(c *fiber.Ctx) (err error) {
	var dataBody = new(dto.RegisterRequest)

	if err := c.BodyParser(&dataBody); err != nil {
		return err
	}

	err = dataBody.Validate()
	if err != nil {
		return err
	}

	// get existing user
	userByEmail, err := s.repoUsers.Get(repository.ParamGetUser{
		Email: dataBody.Email,
	})
	if err != nil {
		return err
	}

	if userByEmail.Email == dataBody.Email {
		return fmt.Errorf("email is exists")
	}

	userByUsername, err := s.repoUsers.Get(repository.ParamGetUser{
		Username: dataBody.Username,
	})
	if err != nil {
		return err
	}

	if userByUsername.Username == dataBody.Username {
		return fmt.Errorf("email is exists")
	}

	passwordHash, err := dataBody.GetHashPassword()
	if err != nil {
		return fmt.Errorf("error hash password")
	}

	userDB := dataBody.PrepareDataDB(passwordHash, false)

	err = s.repoUsers.Add(userDB)
	if err != nil {
		return fmt.Errorf("failed insert data user")
	}

	return nil
}

func (s *service) Login(c *fiber.Ctx) (jwtToken string, err error) {

	var (
		dataBody = new(dto.LoginRequest)
		userData dto.UserDB
	)

	if err := c.BodyParser(&dataBody); err != nil {
		return jwtToken, err
	}

	if dataBody.UsernameIsEmail() {

		userData, err = s.repoUsers.Get(repository.ParamGetUser{
			Email: dataBody.Username,
		})
		if err != nil {
			return jwtToken, fmt.Errorf("data user not found")
		}

	} else {

		userData, err = s.repoUsers.Get(repository.ParamGetUser{
			Username: dataBody.Username,
		})
		if err != nil {
			return jwtToken, fmt.Errorf("data user not found")
		}
	}

	if userData.ID == 0 {
		return jwtToken, fmt.Errorf("data user not found")
	}

	if !dataBody.CheckPasswordHash(userData.Password) {
		return jwtToken, fmt.Errorf("invalid password")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = userData.Username
	claims["user_id"] = userData.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	jwtToken, err = token.SignedString([]byte(config.GetConfig("JWT_SECRET")))
	if err != nil {
		return jwtToken, err
	}

	return jwtToken, nil
}
