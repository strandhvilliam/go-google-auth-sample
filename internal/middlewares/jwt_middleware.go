package middlewares

import (
	"google-auth-go-v2/internal/config"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func JwtMiddleware() fiber.Handler {
	cfg := config.NewJWTConfig()
	return jwtware.New(cfg())
}
