// Package controllers Dart-Scoreboard APIs.
//
//  Schemes: http
//  Host: localhost:8080
//  BasePath: /api/v1
//  Version: v1
//
//  Games:
//  - application/json
//
//  Users:
//  - application/json
//
// swagger:meta
package controllers

import (
	models "dartscoreboard/models/database"
	types "dartscoreboard/models/types"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// swagger:route GET/users Users ListUsers
// Returns a list of users
// Responses:
//  200: UsersPaginationResponse
//	400: StatusCode
//  500: StatusCode
// GetUsers are get all users that login in dart-scoreboard
func GetUsers(ctx *fiber.Ctx) error {
	user := types.User{}
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		return ctx.JSON(types.StatusCode{
			StatusCode: 400,
			Message:    "Bad Request",
		})
	}
	offset := (page - 1) * 5
	users, err := models.GetUsers(db, offset, user)
	if err != nil {
		return ctx.JSON(types.StatusCode{
			StatusCode: 500,
			Message:    "Internal Server Error",
		})
	}
	prePage := page - 1
	postPage := page + 1
	prePageLink := fmt.Sprintf("/api/v1/users?page=%d", prePage)
	postPageLink := fmt.Sprintf("/api/v1/users?page=%d", postPage)
	if len(users) < 5 {
		postPageLink = "cross limits"
	}
	if prePage == 0 {
		prePageLink = "cross limits"
	}
	ctx.SendStatus(200)
	return ctx.JSON(types.UsersPaginationResponse{
		UserResponses: users,
		PrePageLink:   prePageLink,
		PostPageLink:  postPageLink,
	})
}
