package models

import (
	"dartscoreboard/models/types"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Update Games From Games Table By Game Id
func UpdateGame(db *sql.DB, id int, user types.User, game types.Game, playerRes types.GamePlayerResponse) (types.GameResponse, error) {
	var gameResJson types.GameResponse
	createrInfo, err := SelectUserInfoByEmail(db, game.CreaterUserEmail, user)
	if err != nil {
		fmt.Println(err)
		return gameResJson, err
	}
	query, err := db.Prepare("UPDATE games SET name = ?, type = ?, status = ?, creater_user_id = ? WHERE id = ?;")
	if err != nil {
		fmt.Println(err)
		return gameResJson, err
	}
	_, err = query.Exec(game.Name, game.Type, game.Status, createrInfo.Id, id)
	if err != nil {
		fmt.Println(err)
		return gameResJson, err
	}
	deleteGamePlayers, err := db.Prepare("DELETE FROM game_players WHERE game_id = ?;")
	if err != nil {
		fmt.Println(err)
		return gameResJson, err
	}
	_, err = deleteGamePlayers.Exec(id)
	if err != nil {
		fmt.Println(err)
		return gameResJson, err
	}

	playersInfo, err := InsertPlayers(db, game, id, playerRes)
	if err != nil {
		fmt.Println(err)
		return gameResJson, err
	}
	gameResJson = types.GameResponse{
		Id:               id,
		Name:             game.Name,
		Type:             game.Type,
		Status:           game.Status,
		CreaterName: createrInfo.FirstName+createrInfo.LastName,
		Players:          playersInfo,
	}
	return gameResJson, nil
}
