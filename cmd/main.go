package main

import (
	"google-auth-go-v2/internal/config"
	"google-auth-go-v2/internal/infra"
	"google-auth-go-v2/internal/middlewares"
	"google-auth-go-v2/internal/routes"
	"google-auth-go-v2/internal/utils"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := infra.GetDb().Connect(); err != nil {
		log.Fatal(err)
	}
	err = utils.SeedDb()
	if err != nil {
		panic(err)
	}

	infra.NewAuthProvider(config.GoogleAuthConfig())
	infra.InitSessionStorage()
	app := fiber.New()
	app.Use(middlewares.JwtMiddleware())

	initRoutes(app)
	shutdownRoutine(app)

	app.Listen(":8080")
}

func initRoutes(app *fiber.App) {
	routes.AuthenticationRoutes(app)
	routes.SampleRoutes(app)
}

func shutdownRoutine(app *fiber.App) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		<-sigChan
		infra.GetDb().Close()
		_ = app.Shutdown()
	}()
}
