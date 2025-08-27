package routes

import (
	"login-system/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// init handler
	authHandler := handlers.NewAuthHandler()

	// proses login
	app.Post("/login", authHandler.ProcessLogin)

	// membuat user
	app.Post("/create-user", authHandler.CreateUser)
}
