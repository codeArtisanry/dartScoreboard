package controllers

import (
	models "dartscoreboard/models/database"
	types "dartscoreboard/models/types"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func UndoScore(ctx *fiber.Ctx) error {
	gameId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(400).JSON(types.StatusCode{
			StatusCode: 400,
			Message:    "Bad Request",
		})
	}
	err = models.UndoScore(db, gameId)
	if err != nil {
		return ctx.Status(400).JSON(types.StatusCode{
			StatusCode: 400,
			Message:    "Internal Server Error",
		})
	}
	return ctx.SendStatus(204)
}
