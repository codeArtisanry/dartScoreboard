package controllers

import (
	models "dartscoreboard/models/database"
	types "dartscoreboard/models/types"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetWinner(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	gameId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("gameid", err)
		return ctx.Status(400).JSON(types.StatusCode{
			StatusCode: 400,
			Message:    "Bad Request",
		})
	}
	getwinnerjson, err := models.FoundWinner(db, gameId)
	if err != nil {
		fmt.Println("getwinnerjson", err)
		return ctx.Status(500).JSON(types.StatusCode{
			StatusCode: 500,
			Message:    "Internal Server Error",
		})
	}
	fmt.Println("finised")
	return ctx.Status(200).JSON(getwinnerjson)
}
