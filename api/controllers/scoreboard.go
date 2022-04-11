package controllers

import (
	"dartscoreboard/models"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetScoreboard(c *fiber.Ctx) error {
	db := models.Database()
	id := c.Params("id")
	gameId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
	}
	game := models.Game{}
	c.BodyParser(&game)
}
