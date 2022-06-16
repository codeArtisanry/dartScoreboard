package controllers

import (
	models "dartscoreboard/models/database"
	types "dartscoreboard/models/types"
	services "dartscoreboard/services"
	"database/sql"
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
func GetCurrentPlayerInfoAPI(ctx *fiber.Ctx) error {

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
	if playerId == 0 {
		return ctx.Status(400).JSON(types.StatusCode{
			StatusCode: 400,
			Message:    "Invalid Player",
		})
	}
	gameRes := types.GameResponse{}
	activePlayerInfo, err := GetCurrentPlayerInfo(db, gameId, playerId, gameRes)
	if err != nil {
		return ctx.Status(500).JSON(types.StatusCode{
			StatusCode: 500,
			Message:    "Internal Server Error",
		})
	}
	return ctx.Status(200).JSON(activePlayerInfo)
}

func GetCurrentPlayerInfo(db *sql.DB, gameId int, playerId int, gameRes types.GameResponse) (types.CurrentPlayerInfo, error) {
	var (
		CurrentPlayerInfo types.CurrentPlayerInfo
		activeRes         types.ActiveStatus
	)
	game := models.GameResQuery(db, gameId, gameRes)
	ActiveStatus, err := GetActiveStatusRes(db, gameId, activeRes)
	if err != nil {
		return CurrentPlayerInfo, err
	}
	id, name, email := models.QueryForPlayer(db, playerId)
	score := models.ActivePlayerTotal(db, gameId, playerId)
	score = services.RemainScore(game.Type, score)
	if ActiveStatus.Round == 0 {
		CurrentPlayerInfo = types.CurrentPlayerInfo{
			Id:    id,
			Name:  name,
			Email: email,
			Score: score,
			Game:  game,
		}
		return CurrentPlayerInfo, nil
	} else {
		round := ActiveStatus.Round
		throw := ActiveStatus.Throw
		CurrentPlayerInfo = types.CurrentPlayerInfo{
			Id:    id,
			Name:  name,
			Email: email,
			Round: round,
			Throw: throw,
			Score: score,
			Game:  game,
		}
		return CurrentPlayerInfo, nil
	}

}
