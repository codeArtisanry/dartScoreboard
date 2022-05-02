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
//  200: CurrentPlayerInfo
//  400: StatusCode
//  404: StatusCode
//  500: StatusCode
// StartGame are get that game which is you want to fetch
func GetCurrentPlayerInfo(ctx *fiber.Ctx) error {
	fmt.Println("problem")
	gameId, err := strconv.Atoi(ctx.Params("id"))
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
	gameRes := types.GameResponse{}
	activePlayerInfo, err := models.GetCurrentPlayerInfo(db, gameId, playerId, gameRes)
	if err != nil {
		return ctx.Status(500).JSON(types.StatusCode{
			StatusCode: 500,
			Message:    "Internal Server Error",
		})
	}
	return ctx.Status(200).JSON(activePlayerInfo)
}
