package models

import (
	"dartscoreboard/models/types"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// get current player info
func GetCurrentPlayerInfo(db *sql.DB, gameId int, playerId int, gameRes types.GameResponse) (types.CurrentPlayerInfo, error) {
	var (
		CurrentPlayerInfo types.CurrentPlayerInfo
		activeRes         types.ActiveStatus
		ActivePlayerInfo  types.ActivePlayerInfo
	)
	query := fmt.Sprintf("SELECT id, name, type, status FROM games WHERE id = %d;", gameId)
	row := db.QueryRow(query)
	err := row.Scan(&gameRes.Id, &gameRes.Name, &gameRes.Type, &gameRes.Status)
	if err != nil {
		return CurrentPlayerInfo, err
	}
	ActiveStatus, err := GetActiveStatusRes(db, gameId, activeRes)
	if err != nil {
		fmt.Println(err)
		return CurrentPlayerInfo, err
	}
	ActivePlayerTotal := fmt.Sprintf("select ifnull(sum(s.score),0) from scores s left join game_players gp on gp.id = s.game_player_id WHERE gp.game_id = %d AND gp.user_id = %d  AND s.is_valid = 'VALID';", gameId, playerId)
	rowsPlayerTotal := db.QueryRow(ActivePlayerTotal)
	err = rowsPlayerTotal.Scan(&ActivePlayerInfo.Score)
	if err != nil {
		return CurrentPlayerInfo, err
	}
	queryPlayer := fmt.Sprintf("SELECT id, first_name, last_name, email FROM users WHERE id = %d", playerId)
	rowsPlayers := db.QueryRow(queryPlayer)
	err = rowsPlayers.Scan(&ActivePlayerInfo.Id, &ActivePlayerInfo.FirstName, &ActivePlayerInfo.LastName, &ActivePlayerInfo.Email)
	if err != nil {
		return CurrentPlayerInfo, err
	}
	if ActiveStatus.Round == 0 {
		CurrentPlayerInfo = types.CurrentPlayerInfo{
			Id:               gameRes.Id,
			Name:             gameRes.Name,
			Type:             gameRes.Type,
			Status:           gameRes.Status,
			ActivePlayerInfo: &ActivePlayerInfo,
		}
		return CurrentPlayerInfo, nil
	}
	CurrentPlayerInfo = types.CurrentPlayerInfo{
		Id:               gameRes.Id,
		Name:             gameRes.Name,
		Type:             gameRes.Type,
		Round:            ActiveStatus.Round,
		Throw:            ActiveStatus.Throw,
		ActivePlayerInfo: &ActivePlayerInfo,
	}
	return CurrentPlayerInfo, nil
}
