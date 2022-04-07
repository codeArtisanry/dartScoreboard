package controllers

import (
	"dartscoreboard/models"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(ctx *fiber.Ctx) error {
	db := models.Database()
	user := models.User{}
	users, err := models.GetUsers(db, user)
	if err != nil {
		ctx.SendStatus(500)
	}
	ctx.JSON(users)
	return ctx.SendStatus(200)
}
