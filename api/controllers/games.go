package controllers

import (
	"dartscoreboard/models"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func InsertGame(ctx *fiber.Ctx) error {
	db := models.Database()
	game := models.Game{}
	ctx.BodyParser(&game)
	gameJson, err := models.InsertGames(db, game)
	if err != nil {
		fmt.Println(err)
		ctx.SendStatus(500)
	}
	ctx.SendStatus(201)
	return ctx.JSON(gameJson)
}

func DeleteGame(ctx *fiber.Ctx) error {
	db := models.Database()
	id := ctx.Params("id")
	gameId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
	}
	err = models.DeleteGames(db, gameId)
	if err != nil {
		return ctx.SendStatus(500)
	}
	return ctx.SendStatus(204)
}

func UpdateGame(ctx *fiber.Ctx) error {
	db := models.Database()
	id := ctx.Params("id")
	gameId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
	}
	game := models.Game{}
	ctx.BodyParser(&game)
	row, err := models.UpdateGame(db, gameId, game)
	if err != nil {
		return ctx.SendStatus(500)
	}
	return ctx.JSON(row)
}
func GetGame(ctx *fiber.Ctx) error {
	db := models.Database()
	id := ctx.Params("id")
	gameId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
	}
	game := models.Game{}
	gameJson, err := models.GetGame(db, gameId, game)
	if err != nil {
		return ctx.SendStatus(500)
	}
	ctx.SendStatus(200)
	return ctx.JSON(gameJson)
}

func GetGames(ctx *fiber.Ctx) error {
	db := models.Database()
	game := models.Game{}
	games, err := models.GetGames(db, game)
	if err != nil {
		fmt.Println(err)
		return ctx.SendStatus(500)
	}
	ctx.SendStatus(200)
	return ctx.JSON(games)
}
