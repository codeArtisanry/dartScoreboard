package models

import (
	"dartscoreboard/models/types"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Insert Games in to Games Table
func InsertGames(db *sql.DB, createrEmail string, user types.User, game types.Game, gamePlayer types.GamePlayer, gamePlayerRes types.GamePlayerResponse) (types.GameResponse, error) {
	var gameResJson types.GameResponse
	createrInfo, err := SelectUserInfoByEmail(db, createrEmail, user)
	if err != nil {
		fmt.Println(err)
		return gameResJson, err
	}
	insert, err := db.Prepare("INSERT INTO games(name, type, status, creater_user_id) VALUES(?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
		return gameResJson, err
	}
	result, err := insert.Exec(game.Name, game.Type, game.Status, createrInfo.Id)
	if err != nil {
		fmt.Println(err)
		return gameResJson, err
	}
	gameId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return gameResJson, err
	}

	id := int(gameId)
	PlayersInfo, err := InsertPlayers(db, game, id, gamePlayerRes)
	if err != nil {
		fmt.Println(err)
		return gameResJson, err
	}
	gameResJson = types.GameResponse{
		Id:          id,
		Name:        game.Name,
		Type:        game.Type,
		Status:      game.Status,
		CreaterName: createrInfo.FirstName + createrInfo.LastName,
		Players:     PlayersInfo,
	}
	return gameResJson, nil
}
