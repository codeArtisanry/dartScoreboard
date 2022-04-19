package controllers

import (
	models "dartscoreboard/models/database"
	types "dartscoreboard/models/types"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetActiveStatusRes(ctx *fiber.Ctx) error {
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
	activejson, err := models.GetActiveStatusRes(db, gameId, activeRes)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(404).JSON(types.StatusCode{
			StatusCode: 404,
			Message:    "No Content",
		})
	}
	return ctx.Status(200).JSON(activejson)
}
