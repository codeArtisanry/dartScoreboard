package controllers

import (
	models "dartscoreboard/models/database"
	types "dartscoreboard/models/types"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// swagger:route POST/games/{id}/score Scores addScore
// Insert Score and Round in scores and rounds table
// Responses:
//  201: ResScore
//  400: StatusCode
//  500: StatusCode
// InsertScore is insert score that post in score api by user
func InsertScoreAPI(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	db := models.Database("dart.db")
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
	currentStateOfGame, err := GetActiveStatusRes(db, gameId, activeRes)
	if err != nil {
		return ctx.Status(400).JSON(types.StatusCode{
			StatusCode: 500,
			Message:    "Internal Server Error",
		})
	}
	if playerId == currentStateOfGame.PlayerId && round == currentStateOfGame.Round && turn == currentStateOfGame.Throw {
		if score.Score >= 0 && score.Score <= 60 {
			activeRes := types.ActiveStatus{}
			activejson, err := GetActiveStatusRes(db, gameId, activeRes)
			if err != nil {
				fmt.Println(err)

			}
			if activejson.Round == 0 {
				scoreRes := types.ResScore{
					Score:       0,
					TotalScore:  0,
					FoundWinner: true,
				}
				fmt.Println(scoreRes)
			}
			scoreRes, err := InsertScore(db, gameId, playerId, round, turn, score)
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

func InsertScore(db *sql.DB, gameId int, playerId int, round int, turnId int, score types.Score) (types.ResScore, error) {
	var (
		totalScore   int
		roundId      int
		gamePlayerId int
		scoresId     int
		gameType     string
		scoreRes     types.ResScore
	)

	models.VerifyRoundTableQuery(db, gameId, round, roundId)
	gamePlayerId, roundId = models.RoundGamePlayerIdQuery(db, gameId, playerId, round)
	totalScore, err := models.FindTotalScore(db, gamePlayerId)
	if err != nil {
		fmt.Println(err)
		return scoreRes, err
	}
	gameType = models.FindGameTypeQuery(db, gameId)
	totalScore = totalScore + score.Score
	rowScore := models.ValidateScoreQuery(db, gameId, playerId, round, turnId)
	err = rowScore.Scan(&scoresId)
	if err != nil {
		if err == sql.ErrNoRows {
			if gameType == "High Score" {
				models.InsertIntoScoreTableQuery(db, playerId, round, turnId, score, roundId, gamePlayerId)
				scoreRes = types.ResScore{
					Score:       score.Score,
					TotalScore:  totalScore,
					FoundWinner: false,
				}
				return scoreRes, nil
			} else {
				fullGameType := strings.Split(gameType, "-")
				targetscore, err := strconv.Atoi(fullGameType[1])
				if err != nil {
					fmt.Println(err)
					return scoreRes, err
				}
				if totalScore <= targetscore {
					models.InsertIntoScoreTableQuery(db, playerId, round, turnId, score, roundId, gamePlayerId)
					if totalScore == targetscore {
						status := "Completed"
						err = models.UpdateStatus(db, gameId, status)
						if err != nil {
							fmt.Println(err)
						}
						scoreRes = types.ResScore{
							Score:       score.Score,
							TotalScore:  targetscore - totalScore,
							FoundWinner: true,
						}
						return scoreRes, nil
					}
					totalScore, err = models.FindTotalScore(db, gamePlayerId)
					if err != nil {
						fmt.Println(err)
						return scoreRes, err
					}
					scoreRes = types.ResScore{
						Score:       score.Score,
						TotalScore:  targetscore - totalScore,
						FoundWinner: false,
					}
					return scoreRes, nil
				} else {
					models.InsertIntoScoreTableQuery(db, playerId, round, turnId, score, roundId, gamePlayerId)
					for throw := turnId + 1; throw <= 3; throw++ {
						models.RemoveMultipleEntryInScore(db, roundId, gamePlayerId, throw)
					}
					for throw := 1; throw <= 3; throw++ {
						models.QueryForUpdateIsValid(db, roundId, gamePlayerId, throw)
					}
					totalScore, err = models.FindTotalScore(db, gamePlayerId)
					if err != nil {
						fmt.Println(err)
						return scoreRes, err
					}
					scoreRes = types.ResScore{
						Score:       score.Score,
						TotalScore:  targetscore - totalScore,
						FoundWinner: false,
					}
					return scoreRes, nil
				}
			}
		}
		return scoreRes, err
	} else {
		scoreRes = types.ResScore{
			Score:       61,
			TotalScore:  0,
			FoundWinner: false,
		}
		return scoreRes, nil
	}
}
