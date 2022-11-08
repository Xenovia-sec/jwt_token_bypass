package main

import (
	"jwt_example/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("./src/layouts/", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	routes.Setup(app)
	app.Static("/assets", "./src/public")

	app.Listen(":8000")
}
