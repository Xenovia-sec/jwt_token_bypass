package routes

import (
	"jwt_example/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Index)
	app.Get("/getjwt", controllers.GetJWT)
	app.Post("/checkJwt", controllers.CheckJwt)

}
