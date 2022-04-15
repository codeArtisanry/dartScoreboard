package models

import (
	types "dartscoreboard/models/types"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Get Game From Games Table By Game Id
func GetGame(db *sql.DB, id int, gameRes types.GameResponse, user types.User, gamePlayerRes types.GamePlayerResponse) (types.GameResponse, error) {
	var gameResJson types.GameResponse
	query := fmt.Sprintf("SELECT id, name, type, status, creater_user_id FROM games WHERE id = %d;", id)
	row := db.QueryRow(query)
	err := row.Scan(&gameRes.Id, &gameRes.Name, &gameRes.Type, &gameRes.Status, &gameRes.CreaterUserId)
	if err != nil {
		return gameResJson, err
	}
	createrInfo, err := SelectUserInfoById(db, gameRes.CreaterUserId, user)
	if err != nil {
		return gameResJson, err
	}
	PlayersInfo, err := GetPlayersInfoByGameId(db, id, gamePlayerRes)
	if err != nil {
		fmt.Println(err)
		return gameResJson, err
	}
	gameResJson = types.GameResponse{
		Id:          gameRes.Id,
		Name:        gameRes.Name,
		Type:        gameRes.Type,
		Status:      gameRes.Status,
		CreaterName: createrInfo.FirstName + createrInfo.LastName,
		Players:     PlayersInfo,
	}
	return gameResJson, nil
}

//Get All Games From Game Table
func GetGames(db *sql.DB, loginUserId int, page int, offset int, gameRes types.GameResponse, user types.User, gamePlayerRes types.GamePlayerResponse) ([]types.GameResponse, error) {
	var gamesResJson []types.GameResponse
	query := fmt.Sprintf("SELECT games.id, games.name, games.type, games.status, games.creater_user_id FROM games LEFT JOIN game_players ON games.id = game_players.game_id WHERE game_players.user_id = %d ORDER BY games.created_at DESC LIMIT 10 OFFSET %d;", loginUserId, offset)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return gamesResJson, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&gameRes.Id, &gameRes.Name, &gameRes.Type, &gameRes.Status, &gameRes.CreaterUserId)
		if err != nil {
			fmt.Println(err)
			return gamesResJson, err
		}
		createrInfo, err := SelectUserInfoById(db, gameRes.CreaterUserId, user)
		if err != nil {
			fmt.Println(err)
			return gamesResJson, err
		}
		playersInfo, err := GetPlayersInfoByGameId(db, gameRes.Id, gamePlayerRes)
		if err != nil {
			fmt.Println(err)
			return gamesResJson, err
		}
		gameResJson := types.GameResponse{
			Id:          gameRes.Id,
			Name:        gameRes.Name,
			Type:        gameRes.Type,
			Status:      gameRes.Status,
			CreaterName: createrInfo.FirstName + createrInfo.LastName,
			Players:     playersInfo}
		gamesResJson = append(gamesResJson, gameResJson)
	}
	return gamesResJson, nil
}
