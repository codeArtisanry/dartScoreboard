package models

import (
	"dartscoreboard/models/types"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func (database DataBase) GameResQuery(gameId int, gameRes types.GameResponse) (Id int, Name string, Type string, Status string) {
	db := database.Db
	query := fmt.Sprintf("SELECT id, name, type, status FROM games WHERE id = %d;", gameId)
	row := db.QueryRow(query)
	err := row.Scan(&gameRes.Id, &gameRes.Name, &gameRes.Type, &gameRes.Status)
	if err != nil {
		fmt.Println(err)
	}
	Id = gameRes.Id
	Name = gameRes.Name
	Type = gameRes.Type
	Status = gameRes.Status
	fmt.Println(Name, "fdfdfd")
	return Id, Name, Type, Status
}

func (database DataBase) ActivePlayerTotal(gameId int, playerId int) (Score int) {
	db := database.Db
	var ActivePlayerInfo types.ActivePlayerInfo
	ActivePlayerTotal := fmt.Sprintf("select ifnull(sum(s.score),0) from scores s left join game_players gp on gp.id = s.game_player_id WHERE gp.game_id = %d AND gp.user_id = %d  AND s.is_valid = 'VALID';", gameId, playerId)
	rowsPlayerTotal := db.QueryRow(ActivePlayerTotal)
	err := rowsPlayerTotal.Scan(&ActivePlayerInfo.Score)
	if err != nil {
		fmt.Println(err)
	}
	Score = ActivePlayerInfo.Score
	return Score
}

func (database DataBase) QueryForPlayer(playerId int) (Id int, FirstName string, LastName string, Email string) {
	db := database.Db
	var ActivePlayerInfo types.ActivePlayerInfo
	queryPlayer := fmt.Sprintf("SELECT id, first_name, last_name, email FROM users WHERE id = %d", playerId)
	rowsPlayers := db.QueryRow(queryPlayer)
	err := rowsPlayers.Scan(&ActivePlayerInfo.Id, &ActivePlayerInfo.FirstName, &ActivePlayerInfo.LastName, &ActivePlayerInfo.Email)
	if err != nil {
		fmt.Println(err)
	}
	Id = ActivePlayerInfo.Id
	FirstName = ActivePlayerInfo.FirstName
	LastName = ActivePlayerInfo.LastName
	Email = ActivePlayerInfo.Email
	return Id, FirstName, LastName, Email
}
