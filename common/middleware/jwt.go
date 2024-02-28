package middleware

import (
	"online-store/common/dto"
	"online-store/config"
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// Protected protect routes
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(config.GetConfig("JWT_SECRET"))},
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}

// get data jwt
func GetDataJWT(token interface{}) *dto.ClaimJWTData {
	jwtToken := token.(*jwt.Token)
	dataJwt := ClaimJWT(jwtToken)

	return dataJwt
}

func ClaimJWT(t *jwt.Token) *dto.ClaimJWTData {

	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))
	exp := int(claims["exp"].(float64))
	username := claims["username"].(string)

	if uid > 0 && username != "" {
		data := dto.ClaimJWTData{
			UserID:   uid,
			Username: username,
			Exp:      exp,
			ExpDate:  time.Unix(int64(exp), 0),
		}

		return &data
	}

	return nil
}
