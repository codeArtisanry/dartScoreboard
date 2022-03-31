package api

import (
	"dartscoreboard/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GameApi(app *fiber.App) {
	app.Get("/api/v1/game/", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	app.Post("/api/v1/game/", func(c *fiber.Ctx) error {
		g := new(models.Game)

		if err := c.BodyParser(g); err != nil {
			return err
		}
		c.JSON(g)
		fmt.Println(g)
		db := models.Database()
		models.InsertGameData(db, *g)
		return c.SendStatus(200)
	})
}
func PointApi(app *fiber.App){
	app.Get("/api/v1/point", func(c *fiber.Ctx) error {
		p := new(models.Score)
		if err := c.BodyParser(p); err != nil {
			return err
		}
		c.JSON(p)
		fmt.Println(p)
		return c.SendStatus(200)

	})
}
