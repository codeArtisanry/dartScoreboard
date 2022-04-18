package models

import (
	"dartscoreboard/models/types"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Get all perticuler game of players infomation list by game id
func GetPlayersInfoByGameId(db *sql.DB, id int, gamePlayerRes types.GamePlayerResponse) ([]types.GamePlayerResponse, error) {
	var PlayersInfo []types.GamePlayerResponse
	query := fmt.Sprintf("SELECT DISTINCT users.id,users.first_name,users.last_name,users.email from users LEFT JOIN game_players ON game_players.user_id=users.id where game_players.game_id = %d;", id)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return PlayersInfo, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&gamePlayerRes.Id, &gamePlayerRes.FirstName, &gamePlayerRes.LastName, &gamePlayerRes.Email)
		if err != nil {
			fmt.Println(err)
			return PlayersInfo, err
		}
		gamePlayerInfo := types.GamePlayerResponse{
			Id:        gamePlayerRes.Id,
			FirstName: gamePlayerRes.FirstName,
			LastName:  gamePlayerRes.LastName,
			Email:     gamePlayerRes.Email}
		PlayersInfo = append(PlayersInfo, gamePlayerInfo)
	}
	return PlayersInfo, nil
}
