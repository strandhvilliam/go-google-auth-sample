package config

import (
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
)

func NewJWTConfig() func() jwtware.Config {
	return func() jwtware.Config {
		envSecret := os.Getenv("JWT_SECRET")
		return jwtware.Config{
			SigningKey: jwtware.SigningKey{Key: []byte(envSecret)},
		}
	}
}
