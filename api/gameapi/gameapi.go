package api

import (
	"dartscoreboard/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Gameapi(app *fiber.App) {
	app.Get("/game/api/v1", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	app.Post("/game/api/v1/", func(c *fiber.Ctx) error {
		g := new(models.Game)

		if err := c.BodyParser(g); err != nil {
			return err
		}
		c.JSON(g)
		fmt.Println(g)
		return c.SendStatus(200)
	})
}
