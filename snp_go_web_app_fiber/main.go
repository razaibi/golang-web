package main

import (
	"snp_go_web_app_fiber/db"
	"snp_go_web_app_fiber/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars/v2"
)

func main() {
	db.InitDB()
	// Initialize Handlebars template engine
	engine := handlebars.New("./views", ".hbs")

	// Create a new Fiber instance with the Handlebars engine
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Serve static files from the public directory
	app.Static("/", "./public")

	// Routes
	router.SetupRoutes(app)

	// Start the server
	app.Listen(":3000")
}
