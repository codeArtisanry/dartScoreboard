package routes

import (
	"dartscoreboard/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Endpoint)
	app.Get("/auth/:provider", controllers.Signinflow)
	app.Get("/auth/:provider/callback", controllers.GoogleRedirect)
	app.Get("/logout/:provider", controllers.Signout)
}
