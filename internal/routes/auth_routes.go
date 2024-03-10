package routes

import (
	"google-auth-go-v2/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func AuthenticationRoutes(app *fiber.App) {
	authRoute := app.Group("/auth")
	authRoute.Get("/google/login", handlers.LoginHandler)
	authRoute.Get("/google/callback", handlers.CallbackHandler)
}
