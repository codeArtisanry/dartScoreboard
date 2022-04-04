package routes

import (
	"dartscoreboard/controllers"
	"dartscoreboard/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/auth/:provider", controllers.Signinflow)
	app.Get("/auth/:provider/callback", controllers.GoogleRedirect)
	app.Get("/logout/:provider", controllers.Signout)
	group := app.Group("/", middleware.Validate())
	group.Get("/", controllers.Endpoint)
}
