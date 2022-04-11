package controllers

import (
	"dartscoreboard/models"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// swagger:route POST/games games
// Insert game in game table and also players in game_player table
// Responses:
//  201: GameResponse
//  500: StatusCode
// InsertGame are insert game that register in dart-scoreboard
func InsertGame(ctx *fiber.Ctx) error {
	user := models.User{}
	game := models.Game{}
	gameRes := models.GameResponce{}
	gamePlayer := models.GamePlayer{}
	gamePlayerRes := models.GamePlayerResponce{}
	ctx.BodyParser(&game)
	gameJson, err := models.InsertGames(db, user, game, gameRes, gamePlayer, gamePlayerRes)
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(models.StatusCode{
			StatusCode: 500,
			Message:    "Internal Server Error",
		})
	}
	ctx.SendStatus(201)
	return ctx.JSON(gameJson)
}

// swagger:route DELETE/games/{id} games
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
		return ctx.JSON(models.StatusCode{
			StatusCode: 400,
			Message:    "Bad Request",
		})
	}
	email := LoginUser.Email
	user := models.User{}
	gameRes := models.GameResponce{}
	gameId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(models.StatusCode{
			StatusCode: 400,
			Message:    "Bad Request",
		})
	}
	creater_id, err := models.FindCreaterIdByGameId(db, gameId, gameRes)
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(models.StatusCode{
			StatusCode: 400,
			Message:    "Bad Request",
		})
	}
	user, err = models.SelectUserInfoByEmail(db, email, user)
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(models.StatusCode{
			StatusCode: 400,
			Message:    "Bad Request",
		})
	}
	if creater_id == user.Id {
		err = models.DeleteGames(db, gameId)
		if err != nil {
			return ctx.JSON(models.StatusCode{
				StatusCode: 500,
				Message:    "Internal Server Error",
			})
		}
		return ctx.JSON(models.StatusCode{
			StatusCode: 204,
			Message:    "No Content",
		})
	}
	return ctx.JSON(models.StatusCode{
		StatusCode: 403,
		Message:    "Forbidden",
	})
}

// swagger:route PUT/games/{id} games
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
		return ctx.JSON(models.StatusCode{
			StatusCode: 400,
			Message:    "Bad Request",
		})
	}
	email := LoginUser.Email
	gameId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(models.StatusCode{
			StatusCode: 400,
			Message:    "Bad Request",
		})
	}
	playerRes := models.GamePlayerResponce{}
	gameRes := models.GameResponce{}
	game := models.Game{}
	user := models.User{}
	ctx.BodyParser(&game)
	creater_id, err := models.FindCreaterIdByGameId(db, gameId, gameRes)
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(models.StatusCode{
			StatusCode: 400,
			Message:    "Bad Request",
		})
	}
	user, err = models.SelectUserInfoByEmail(db, email, user)
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(models.StatusCode{
			StatusCode: 400,
			Message:    "Bad Request",
		})
	}
	if creater_id == user.Id {
		row, err := models.UpdateGame(db, gameId, user, game, playerRes)
		if err != nil {
			return ctx.JSON(models.StatusCode{
				StatusCode: 500,
				Message:    "Internal Server Error",
			})
		}
		ctx.SendStatus(201)
		return ctx.JSON(row)
	}
	return ctx.JSON(models.StatusCode{
		StatusCode: 403,
		Message:    "Forbidden",
	})
}

// swagger:route GET/games/{id} games
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
		return ctx.JSON(models.StatusCode{
			StatusCode: 400,
			Message:    "Bad Request",
		})
	}
	gameRes := models.GameResponce{}
	user := models.User{}
	gamePlayerRes := models.GamePlayerResponce{}

	gameJson, err := models.GetGame(db, gameId, gameRes, user, gamePlayerRes)
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(models.StatusCode{
			StatusCode: 404,
			Message:    "No Content",
		})
	}
	ctx.SendStatus(200)
	return ctx.JSON(gameJson)
}

// swagger:route GET/games games
// Get all the game that is by login user
// Responses:
//  200: GameResponse
//  400: StatusCode
//  500: StatusCode
// GetGames are get that games which participate and register by perticuler user
func GetGames(ctx *fiber.Ctx) error {
	gamePlayer := models.GamePlayerResponce{}
	game := models.GameResponce{}
	user := models.User{}
	page, err := strconv.Atoi(ctx.Params("page"))
	if err != nil {
		return ctx.JSON(models.StatusCode{
			StatusCode: 400,
			Message:    "Bad Request",
		})
	}
	games, err := models.GetGames(db, page, game, user, gamePlayer)
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(models.StatusCode{
			StatusCode: 500,
			Message:    "Internal Server Error",
		})
	}
	ctx.SendStatus(200)
	return ctx.JSON(games)
}
