package main

import (
	"fmt"
	"inventory/cmd/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func (app *Config) AuthenticationMiddleware(c *fiber.Ctx) error {
	accessToken := c.Query("access_token")

	if accessToken == "" {
		return utils.SendResponse("access_token is required", nil, c, fiber.StatusBadRequest)
	}

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWT_SECRET), nil
	})

	if err != nil {
		return utils.SendResponse("failed to parse access_token", nil, c, fiber.StatusInternalServerError)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userEmail := claims["userEmail"].(string)
		role := claims["role"].(string)

		c.Locals("userEmail", userEmail)
		c.Locals("role", role)

		return c.Next()
	}
	return utils.SendResponse("invalid access_token", nil, c, fiber.StatusUnauthorized)
}
