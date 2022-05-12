package controllers

import (
	models "dartscoreboard/models/database"
	types "dartscoreboard/models/types"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// swagger:route POST/games/{id}/score Scores addScore
// Insert Score and Round in scores and rounds table
// Responses:
//  201: ResScore
//  400: StatusCode
//  500: StatusCode
// InsertScore is insert score that post in score api by user
func InsertScore(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	gameId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(400).JSON(types.StatusCode{
			StatusCode: 400,
			Message:    "Bad Request",
		})
	}
	playerId, err := strconv.Atoi(ctx.Params("playerid"))
	if err != nil {
		fmt.Println(err)
		return ctx.Status(400).JSON(types.StatusCode{
			StatusCode: 400,
			Message:    "Bad Request",
		})
	}
	round, err := strconv.Atoi(ctx.Params("roundid"))
	if err != nil {
		fmt.Println(err)
		return ctx.Status(400).JSON(types.StatusCode{
			StatusCode: 400,
			Message:    "Bad Request",
		})
	}
	turn, err := strconv.Atoi(ctx.Params("turnid"))
	if err != nil {
		fmt.Println(err)
		return ctx.Status(400).JSON(types.StatusCode{
			StatusCode: 400,
			Message:    "Bad Request",
		})
	}
	score := types.Score{}
	err = ctx.BodyParser(&score)
	if err != nil {
		return ctx.Status(400).JSON(types.StatusCode{
			StatusCode: 400,
			Message:    "Invalid Score",
		})
	}
	fmt.Println(score.Score)
	activeRes := types.ActiveStatus{}
	currentStateOfGame, err := models.GetActiveStatusRes(db, gameId, activeRes)
	if err != nil {
		return ctx.Status(400).JSON(types.StatusCode{
			StatusCode: 500,
			Message:    "Internal Server Error",
		})
	}
	if playerId == currentStateOfGame.PlayerId && round == currentStateOfGame.Round && turn == currentStateOfGame.Throw {
		if score.Score >= 0 && score.Score <= 60 {
			scoreRes, err := models.InsertScore(db, gameId, playerId, round, turn, score)
			if err != nil {
				fmt.Println(err)
				return ctx.Status(500).JSON(types.StatusCode{
					StatusCode: 500,
					Message:    "Internal Server Error",
				})
			}
			if scoreRes.Score == 61 {
				return ctx.Status(400).JSON(types.StatusCode{
					StatusCode: 400,
					Message:    "Score Already Entered",
				})
			}
			return ctx.Status(201).JSON(scoreRes)
		} else {
			return ctx.Status(400).JSON(types.StatusCode{
				StatusCode: 400,
				Message:    "Invalid Score",
			})
		}
	} else {
		return ctx.Status(400).JSON(types.StatusCode{
			StatusCode: 400,
			Message:    "Turn is Not Matched",
		})
	}
}
