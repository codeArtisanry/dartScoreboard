package controllers

import (
	models "dartscoreboard/models/database"
	types "dartscoreboard/models/types"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// swagger:route GET/current-turn-info StartGame getStartgame
// Responses:
//  200: StartGameResponse
//  400: StatusCode
//  404: StatusCode
//  500: StatusCode
// StartGame are get that game which is you want to fetch
func StartGame(ctx *fiber.Ctx) error {
	fmt.Println("vatsal")
	id := ctx.Params("id")
	gameId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(400).JSON(types.StatusCode{
			StatusCode: 400,
			Message:    "Bad Request",
		})
	}
	gameRes := types.GameResponse{}
	fmt.Println("start")
	startgameJson, err := models.GetStartGame(db, gameId, gameRes)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(500).JSON(types.StatusCode{
			StatusCode: 500,
			Message:    "Internal Server Error",
		})
	}
	fmt.Println("finised")
	return ctx.Status(200).JSON(startgameJson)
}
