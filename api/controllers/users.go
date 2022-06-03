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
func GetUsersAPI(ctx *fiber.Ctx) error {
	user := types.User{}
	page := ctx.Query("page")
	var offset string
	if page == "" {
		page = "0"
		_, err := strconv.Atoi(page)
		if err != nil {
			return ctx.JSON(types.StatusCode{
				StatusCode: 400,
				Message:    "Bad Request",
			})
		}
		offset = "ASC"
		searchFirstName := ctx.Query("sfn")
		searchLastName := ctx.Query("sln")
		searchFirstName = (searchFirstName + "%")
		searchLastName = (searchLastName + "%")
		users, err := GetUsers(offset, searchFirstName, searchLastName, user)
		if err != nil {
			return ctx.JSON(types.StatusCode{
				StatusCode: 500,
				Message:    "Internal Server Error",
			})
		}
		ctx.SendStatus(200)
		return ctx.JSON(types.UsersPaginationResponse{
			UserResponses: users,
		})
	} else {
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			return ctx.JSON(types.StatusCode{
				StatusCode: 400,
				Message:    "Bad Request",
			})
		}
		offset = fmt.Sprintf("ASC LIMIT 5 OFFSET %d", (pageInt-1)*5)
		searchFirstName := ctx.Query("sfn")
		searchLastName := ctx.Query("sln")
		searchFirstName = (searchFirstName + "%")
		searchLastName = (searchLastName + "%")
		users, err := GetUsers(offset, searchFirstName, searchLastName, user)
		if err != nil {
			return ctx.JSON(types.StatusCode{
				StatusCode: 500,
				Message:    "Internal Server Error",
			})
		}
		prePage := pageInt - 1
		postPage := pageInt + 1
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
}

// Get All User From Users Table
func GetUsers(offset string, searchFisrtName string, searchLastName string, user types.User) ([]types.User, error) {
	var users []types.User
	dbcon := models.DataBase{Db: db}
	rows := dbcon.UsersQuery(offset, searchFisrtName, searchLastName)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email)
		if err != nil {
			fmt.Println(err)
			return users, err
		}
		userJson := types.User{
			Id:        user.Id,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email}
		users = append(users, userJson)
	}
	return users, nil
}
