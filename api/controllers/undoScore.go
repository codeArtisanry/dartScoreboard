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
	round, err := strconv.Atoi(ctx.Params("roundid"))
	if err != nil {
		return ctx.Status(400).JSON(types.StatusCode{
			StatusCode: 400,
			Message:    "Bad Request",
		})
	}
	playerId, err := strconv.Atoi(ctx.Params("playerid"))
	if err != nil {
		return ctx.Status(400).JSON(types.StatusCode{
			StatusCode: 400,
			Message:    "Bad Request",
		})
	}
	turn, err := strconv.Atoi(ctx.Params("turnid"))
	if err != nil {
		return ctx.Status(400).JSON(types.StatusCode{
			StatusCode: 400,
			Message:    "Bad Request",
		})
	}
	activeRes := types.ActiveStatus{}
	currentStateOfGame, err := models.GetActiveStatusRes(db, gameId, activeRes)
	if err != nil {
		return ctx.Status(500).JSON(types.StatusCode{
			StatusCode: 500,
			Message:    "Internal Server Error",
		})
	}
	if playerId == currentStateOfGame.PlayerId && round == currentStateOfGame.Round && turn == currentStateOfGame.Throw {
		err = models.UndoScore(db, gameId, round, playerId, turn)
		if err != nil {
			return ctx.Status(500).JSON(types.StatusCode{
				StatusCode: 500,
				Message:    "Internal Server Error",
			})
		}
		return ctx.SendStatus(204)
	} else {
		return ctx.Status(400).JSON(types.StatusCode{
			StatusCode: 400,
			Message:    "Your Turn is not Matching With Current Turn",
		})
	}
}
