package main

import (
	"log"
	"login-system/config"
	"login-system/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// init Redis
	config.InitRedis()

	// buar aplikasi Fiber
	app := fiber.New()

	// setup routes
	routes.SetupRoutes(app)

	// Jalankan server di port 3000
	log.Fatal(app.Listen(":3000"))
}
