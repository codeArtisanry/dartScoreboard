package routes

import (
	"dartscoreboard/controllers"
	"dartscoreboard/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("api/v1/games", controllers.GetGames)
	app.Get("api/v1/games/:id", controllers.GetGame)
	app.Put("api/v1/games/:id/", controllers.UpdateGame)
	app.Delete("api/v1/games/:id", controllers.DeleteGame)
	app.Post("api/v1/games", controllers.InsertGame)
	app.Get("api/v1/users", controllers.GetUsers)
	app.Get("/auth/:provider", controllers.Signinflow)
	app.Get("/auth/:provider/callback", controllers.GoogleRedirect)
	app.Get("/logout/:provider", controllers.Signout)
	group := app.Group("/", middleware.Validate())
	group.Get("/", controllers.Endpoint)
}
