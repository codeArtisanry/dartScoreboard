package controllers

import (
	"dartscoreboard/models"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func InsertGame(ctx *fiber.Ctx) error {
	user := models.User{}
	game := models.Game{}
	ctx.BodyParser(&game)
	gameJson, err := models.InsertGames(db, game, user)
	if err != nil {
		fmt.Println(err)
		return ctx.SendStatus(500)
	}
	ctx.SendStatus(201)
	return ctx.JSON(gameJson)
}

func DeleteGame(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	gameId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return ctx.SendStatus(400)
	}
	err = models.DeleteGames(db, gameId)
	if err != nil {
		return ctx.SendStatus(500)
	}
	return ctx.SendStatus(204)
}

func UpdateGame(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	email := ctx.Params("email")
	gameId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return ctx.SendStatus(400)
	}
	game := models.Game{}
	user := models.User{}
	ctx.BodyParser(&game)
	creater_id, err := models.FindCreaterId(db, game)
	if err != nil {
		fmt.Println(err)
		return ctx.SendStatus(400)
	}
	user_id, err := models.SelectUserIdByEmail(db, email, user)
	if err != nil {
		fmt.Println(err)
		return ctx.SendStatus(400)
	}
	if creater_id == user_id {
		row, err := models.UpdateGame(db, gameId, game)
		if err != nil {
			return ctx.SendStatus(500)
		}
		ctx.SendStatus(203)
		return ctx.JSON(row)
	}
	return ctx.SendStatus(403)
}
func GetGame(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	gameId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return ctx.SendStatus(400)
	}
	game := models.Game{}
	gameJson, err := models.GetGame(db, gameId, game)
	if err != nil {
		fmt.Println(err)
		return ctx.SendStatus(404)
	}
	ctx.SendStatus(200)
	return ctx.JSON(gameJson)
}

func GetGames(ctx *fiber.Ctx) error {
	game := models.Game{}
	user := models.User{}
	email := ctx.Params("email")
	page, err := strconv.Atoi(ctx.Params("page"))
	if err != nil {
		return ctx.SendStatus(400)
	}
	games, err := models.GetGames(db, email, page, game, user)
	if err != nil {
		fmt.Println(err)
		return ctx.SendStatus(500)
	}
	ctx.SendStatus(200)
	return ctx.JSON(games)
}
