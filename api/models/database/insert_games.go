package models

import (
	"dartscoreboard/models/types"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Insert Games in to Games Table
func InsertGames(db *sql.DB, createrEmail string, user types.User, game types.Game, gameRes types.GameResponse, gamePlayer types.GamePlayer, gamePlayerRes types.GamePlayerResponse) (types.GameResponse, error) {
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
	result, err := insert.Exec(game.Name, game.Type, "Not Started", createrInfo.Id)
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
	_, err = InsertPlayers(db, game, id, gamePlayerRes)
	if err != nil {
		fmt.Println(err)
		return gameResJson, err
	}
	gameResJson, err = GetGame(db, id, gameRes, user, gamePlayerRes)
	if err != nil {
		fmt.Println(err)
		return gameResJson, err
	}
	return gameResJson, nil
}
