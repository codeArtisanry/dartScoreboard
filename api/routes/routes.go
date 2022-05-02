package routes

import (
	"dartscoreboard/controllers"
	"dartscoreboard/middleware"

	"github.com/gofiber/fiber/v2"
)

// Routes for diffrent Endpoints
func Setup(app *fiber.App) {
	app.Get("/auth/:provider", controllers.Signinflow)
	app.Get("/auth/:provider/callback", controllers.GoogleRedirect)
	app.Get("/logout/:provider", controllers.Signout)
	app.Get("/", controllers.Endpoint)
	group := app.Group("/", middleware.Validate())
	group.Get("api/v1/users", controllers.GetUsers)
	group.Get("api/v1/games", controllers.GetGames)
	group.Get("api/v1/games/:id", controllers.GetGame)
	group.Get("api/v1/games/:id/scoreboard", controllers.GetScoreboard)
	group.Get("api/v1/games/:id/active-status", controllers.GetActiveStatusRes)
	group.Get("api/v1/games/:id/players/:playerid/player-info", controllers.GetCurrentPlayerInfo)
	group.Post("api/v1/games/:id/score", controllers.InsertScore)
	group.Post("api/v1/games", controllers.InsertGame)
	group.Put("api/v1/games/:id/", controllers.UpdateGame)
	group.Delete("api/v1/games/:id", controllers.DeleteGame)
}
