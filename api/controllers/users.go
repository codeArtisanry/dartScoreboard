//  Schemes: http
//  Host: localhost
//  BasePath: /swaggerdocs
//  Version: 1.0.0
//
//  Users:
//  - application/json
//
//  Games:
//  - application/json
//
// swagger:meta
package controllers

import (
	"dartscoreboard/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// swagger:route GET/users users listUsers
// Returns a list of users
// Responses:
//  200: User
//	400: StatusCode
//  500: StatusCode
// GetUsers are get all users that login in dart-scoreboard
func GetUsers(ctx *fiber.Ctx) error {
	user := models.User{}
	page, err := strconv.Atoi(ctx.Params("page"))
	if err != nil {
		return ctx.JSON(models.StatusCode{
			StatusCode: 400,
			Message:    "Bad Request",
		})
	}
	users, err := models.GetUsers(db, page, user)
	if err != nil {
		return ctx.JSON(models.StatusCode{
			StatusCode: 500,
			Message:    "Internal Server Error",
		})
	}
	ctx.JSON(users)
	return ctx.SendStatus(200)
}
