package controllers

import (
	models "dartscoreboard/models/database"
	types "dartscoreboard/models/types"
	services "dartscoreboard/services"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// swagger:route GET/games/{id}/active-status ActiveStatus activeStatus
// Get activestatus using game id
// Responses:
//  200: ActiveStatus
//  400: StatusCode
//  404: StatusCode
//  500: StatusCode
// GetActiveStatus are get that res which is you want to fetch
func GetActiveStatusResAPI(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	gameId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(400).JSON(types.StatusCode{
			StatusCode: 400,
			Message:    "Bad Request",
		})
	}
	activeRes := types.ActiveStatus{}
	activejson, err := GetActiveStatusRes(gameId, activeRes)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(404).JSON(types.StatusCode{
			StatusCode: 404,
			Message:    "No Content",
		})
	}
	return ctx.Status(200).JSON(activejson)
}

func GetActiveStatusRes(id int, activeRes types.ActiveStatus) (types.ActiveStatus, error) {
	dbcon := models.DataBase{Db: db}
	numOfRowsPerGame, typeOfGame, status, playersIds := dbcon.Query(id, activeRes)

	if numOfRowsPerGame == 0 {
		status := "In Progress"
		err := dbcon.UpdateStatus(id, status)
		if err != nil {
			fmt.Println(err)
		}
		activeRes.Round = 1
		activeRes.Throw = 1
		activeRes.PlayerId = playersIds[0]
	} else {
		activeRes.Round, activeRes.PlayerId, activeRes.Throw = dbcon.Find(id, activeRes)

		if typeOfGame == "High Score" && numOfRowsPerGame%(9*len(playersIds)) == 0 {
			activeRes.Round = 0
			activeRes.PlayerId = 0
			activeRes.Throw = 0
			status := "Completed"
			err := dbcon.UpdateStatus(id, status)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			activeRes.Round, activeRes.PlayerId, activeRes.Throw = services.NextTurn(playersIds, numOfRowsPerGame, activeRes)

		}
		activeRes.Round, activeRes.PlayerId, activeRes.Throw = services.StatusCompleted(status, activeRes)
	}
	activeResJson := types.ActiveStatus{
		GameId:   id,
		Round:    activeRes.Round,
		PlayerId: activeRes.PlayerId,
		Throw:    activeRes.Throw,
	}
	return activeResJson, nil
}
