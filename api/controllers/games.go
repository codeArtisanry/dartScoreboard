package controllers

import (
	"dartscoreboard/models"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// swagger:route POST/games Games addGame
// Insert game in game table and also players in game_player table
// Responses:
//  201: GameResponse
//  500: StatusCode
// InsertGame are insert game that register in dart-scoreboard
func InsertGame(ctx *fiber.Ctx) error {
	user := models.User{}
	game := models.Game{}
	gamePlayer := models.GamePlayer{}
	gamePlayerRes := models.GamePlayerResponse{}
	ctx.BodyParser(&game)
	gameJson, err := models.InsertGames(db, user, game, gamePlayer, gamePlayerRes)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(500).JSON(models.StatusCode{
			StatusCode: 500,
			Message:    "Internal Server Error",
		})
	}
	return ctx.Status(201).JSON(gameJson)
}

// swagger:route DELETE/games/{id} Games deleteGame
// Delete game using game id
// Responses:
//  204: StatusCode
//  400: StatusCode
//  403: StatusCode
//  500: StatusCode
// DeleteGame are Delete that game which is you want to delete
func DeleteGame(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	LoginUser, err := UserJson(ctx)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(400).JSON(models.StatusCode{
			StatusCode: 400,
			Message:    "Bad Request",
		})
	}
	email := LoginUser.Email
	user := models.User{}
	gameRes := models.GameResponse{}
	gameId, err := strconv.Atoi(id)
	fmt.Println(gameId)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(400).JSON(models.StatusCode{
			StatusCode: 400,
			Message:    "Bad Request",
		})
	}
	creater_id, err := models.FindCreaterIdByGameId(db, gameId, gameRes)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(400).JSON(models.StatusCode{
			StatusCode: 400,
			Message:    "Bad Request",
		})
	}
	user, err = models.SelectUserInfoByEmail(db, email, user)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(400).JSON(models.StatusCode{
			StatusCode: 400,
			Message:    "Bad Request",
		})
	}
	if creater_id == user.Id {
		err = models.DeleteGames(db, gameId)
		if err != nil {
			return ctx.Status(500).JSON(models.StatusCode{
				StatusCode: 500,
				Message:    "Internal Server Error",
			})
		}
		return ctx.Status(204).JSON(models.StatusCode{
			StatusCode: 204,
			Message:    "No Content",
		})
	}
	return ctx.Status(403).JSON(models.StatusCode{
		StatusCode: 403,
		Message:    "Forbidden",
	})
}

// swagger:route PUT/games/{id} Games editGame
// Update game using game id
// Responses:
//  201: GameResponse
//  400: StatusCode
//  403: StatusCode
//  500: StatusCode
// UpdateGame are Update that game which is you want to Update
func UpdateGame(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	LoginUser, err := UserJson(ctx)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(400).JSON(models.StatusCode{
			StatusCode: 400,
			Message:    "Bad Request",
		})
	}
	email := LoginUser.Email
	gameId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(400).JSON(models.StatusCode{
			StatusCode: 400,
			Message:    "Bad Request",
		})
	}
	playerRes := models.GamePlayerResponse{}
	gameRes := models.GameResponse{}
	game := models.Game{}
	user := models.User{}
	ctx.BodyParser(&game)
	creater_id, err := models.FindCreaterIdByGameId(db, gameId, gameRes)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(400).JSON(models.StatusCode{
			StatusCode: 400,
			Message:    "Bad Request",
		})
	}
	user, err = models.SelectUserInfoByEmail(db, email, user)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(400).JSON(models.StatusCode{
			StatusCode: 400,
			Message:    "Bad Request",
		})
	}
	if creater_id == user.Id {
		row, err := models.UpdateGame(db, gameId, user, game, playerRes)
		if err != nil {
			return ctx.Status(500).JSON(models.StatusCode{
				StatusCode: 500,
				Message:    "Internal Server Error",
			})
		}
		ctx.SendStatus(201)
		return ctx.Status(201).JSON(row)
	}
	return ctx.Status(403).JSON(models.StatusCode{
		StatusCode: 403,
		Message:    "Forbidden",
	})
}

// swagger:route GET/games/{id} Games getGame
// Get game using game id
// Responses:
//  200: GameResponse
//  400: StatusCode
//  404: StatusCode
//  500: StatusCode
// GetGame are get that game which is you want to fetch
func GetGame(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	gameId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(400).JSON(models.StatusCode{
			StatusCode: 400,
			Message:    "Bad Request",
		})
	}
	gameRes := models.GameResponse{}
	user := models.User{}
	gamePlayerRes := models.GamePlayerResponse{}

	gameJson, err := models.GetGame(db, gameId, gameRes, user, gamePlayerRes)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(404).JSON(models.StatusCode{
			StatusCode: 404,
			Message:    "No Content",
		})
	}
	return ctx.Status(200).JSON(gameJson)
}

// swagger:route GET/games Games ListGame
// Get all the game that is by login user
// Responses:
//  200: GamesPaginationResponse
//  400: StatusCode
//  500: StatusCode
// GetGames are get that games which participate and register by perticuler user
func GetGames(ctx *fiber.Ctx) error {
	gamePlayer := models.GamePlayerResponse{}
	game := models.GameResponse{}
	user := models.User{}
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		return ctx.Status(400).JSON(models.StatusCode{
			StatusCode: 400,
			Message:    "Bad Request",
		})
	}
	offset := (page - 1) * 10
	games, err := models.GetGames(db, page, offset, game, user, gamePlayer)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(500).JSON(models.StatusCode{
			StatusCode: 500,
			Message:    "Internal Server Error",
		})
	}
	prePage := page - 1
	postPage := page + 1
	prePageLink := fmt.Sprintf("/api/v1/games?page=%d", prePage)
	postPageLink := fmt.Sprintf("/api/v1/games?page=%d", postPage)
	if len(games) < 10 {
		postPageLink = "cross limits"
	}
	if prePage == 0 {
		prePageLink = "cross limits"
	}
	return ctx.Status(200).JSON(models.GamesPaginationResponse{
		GameResponses: games,
		PrePageLink:   prePageLink,
		PostPageLink:  postPageLink,
	})
}
