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
	score := types.Score{}
	err = ctx.BodyParser(&score)
	if err != nil {
		return ctx.Status(400).JSON(types.StatusCode{
			StatusCode: 400,
			Message:    "Invalid Score",
		})
	}
	fmt.Println(score.Score)
	if score.Score >= 0 && score.Score <= 60 {
		scoreRes, err := models.InsertScore(db, gameId, score)
		if err != nil {
			fmt.Println(err)
			return ctx.Status(500).JSON(types.StatusCode{
				StatusCode: 500,
				Message:    "Internal Server Error",
			})
		}
		return ctx.Status(201).JSON(scoreRes)
	} else {
		return ctx.Status(400).JSON(types.StatusCode{
			StatusCode: 400,
			Message:    "Invalid Score",
		})
	}
}
