package controllers

import (
	"dartscoreboard/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(ctx *fiber.Ctx) error {
	user := models.User{}
	page, err := strconv.Atoi(ctx.Params("page"))
	if err != nil {
		return ctx.SendStatus(400)
	}
	users, err := models.GetUsers(db, page, user)
	if err != nil {
		ctx.SendStatus(500)
	}
	ctx.JSON(users)
	return ctx.SendStatus(200)
}
