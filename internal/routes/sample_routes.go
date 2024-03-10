package routes

import (
	"google-auth-go-v2/internal/handlers"
	"google-auth-go-v2/internal/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SampleRoutes(app *fiber.App) {
	sampleRoute := app.Group("/sample")
	sampleRoute.Get("/", middlewares.ProtectedRoute, handlers.SampleHandler)
}
