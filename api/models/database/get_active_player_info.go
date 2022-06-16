package models

import (
	"dartscoreboard/models/types"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func GameResQuery(db *sql.DB, gameId int, game types.GameResponse) (gameRes types.GameResponse) {
	query := fmt.Sprintf("SELECT id, name, type, status FROM games WHERE id = %d;", gameId)
	row := db.QueryRow(query)
	err := row.Scan(&game.Id, &game.Name, &game.Type, &game.Status)
	if err != nil {
		fmt.Println(err)
	}
	return game
}

func ActivePlayerTotal(db *sql.DB, gameId int, playerId int) (Score int) {
	ActivePlayerTotal := fmt.Sprintf("select ifnull(sum(s.score),0) from scores s left join game_players gp on gp.id = s.game_player_id WHERE gp.game_id = %d AND gp.user_id = %d  AND s.is_valid = 'VALID';", gameId, playerId)
	rowsPlayerTotal := db.QueryRow(ActivePlayerTotal)
	err := rowsPlayerTotal.Scan(&Score)
	if err != nil {
		fmt.Println(err)
	}
	return Score
}

func QueryForPlayer(db *sql.DB, playerId int) (Id int, Name string, Email string) {
	var playerFirstName string
	var playerLastName string
	queryPlayer := fmt.Sprintf("SELECT id, first_name, last_name, email FROM users WHERE id = %d", playerId)
	rowsPlayers := db.QueryRow(queryPlayer)
	err := rowsPlayers.Scan(&Id, &playerFirstName, &playerLastName, &Email)
	if err != nil {
		fmt.Println(err)
	}
	Name = playerFirstName + " " + playerLastName
	return Id, Name, Email
}
