package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func ProtectedRoute(c *fiber.Ctx) error {
	val := c.Locals("user")
	if val == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	user := val.(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	email := claims["email"].(string)

	c.Locals("id", id)
	c.Locals("email", email)

	return c.Next()
}
