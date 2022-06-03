package controllers

import (
	models "dartscoreboard/models/database"
	types "dartscoreboard/models/types"
	services "dartscoreboard/services"
	"database/sql"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func UndoScoreAPI(ctx *fiber.Ctx) error {
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
	currentStateOfGame, err := GetActiveStatusRes(gameId, activeRes)
	if err != nil {
		return ctx.Status(500).JSON(types.StatusCode{
			StatusCode: 500,
			Message:    "Internal Server Error",
		})
	}
	if playerId == currentStateOfGame.PlayerId && round == currentStateOfGame.Round && turn == currentStateOfGame.Throw {
		err = UndoScore(db, gameId, round, playerId, turn)
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

func UndoScore(db *sql.DB, gameId int, round int, playerId int, turn int) error {
	gamePlayerList, err := models.FindGamePlayers(db, gameId)
	if err != nil {
		log.Println(err)
		return err
	}
	previousTurn := services.FindPreviousTurn(round, playerId, turn, gamePlayerList)
	for i := 1; i <= 3; i++ {
		LastScoreDetails, err := models.FindLastScoreId(db, gameId, models.LastScoreDetails{}, previousTurn)
		if err != nil {
			log.Println(err)
			return err
		}
		if LastScoreDetails.Validate == "INVALID" && LastScoreDetails.Score == 0 {
			err = models.DeleteScore(db, LastScoreDetails.ScoreId)
			if err != nil {
				log.Println(err)
				return err
			}
			log.Println("Successfully Deleted", LastScoreDetails.ScoreId, "ScoreId")
		} else {
			err = models.DeleteScore(db, LastScoreDetails.ScoreId)
			if err != nil {
				log.Println(err)
				return err
			}
			log.Println("Successfully Deleted", LastScoreDetails.ScoreId, "ScoreId")
			models.ChangeValid(db, gameId, previousTurn)
			break
		}
		previousTurn = services.FindPreviousTurn(previousTurn.Round, previousTurn.PlayerId, previousTurn.Turn, gamePlayerList)
	}
	return nil
}
