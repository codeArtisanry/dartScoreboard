package models

import (
	"dartscoreboard/models/types"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// GetActiveStatusRes from from Games table By Game Id
func GetActiveStatusRes(db *sql.DB, id int, activeRes types.ActiveStatus) (types.ActiveStatus, error) {
	var (
		activeResJson    types.ActiveStatus
		playersIds       []int
		numOfRowsPerGame int
		typeOfGame       string
		status           string
		playerId         int
	)
	findNumOfRowsPerGame := fmt.Sprintf("SELECT COUNT(scores.id) from scores inner join game_players on game_players.id=scores.game_player_id WHERE game_players.game_id = %d;", id)
	row := db.QueryRow(findNumOfRowsPerGame)
	err := row.Scan(&numOfRowsPerGame)
	if err != nil {
		fmt.Println(err)
		return activeResJson, err
	}
	findGameInfo := fmt.Sprintf("SELECT type, status from games WHERE id=%d", id)
	rowoftype := db.QueryRow(findGameInfo)
	err = rowoftype.Scan(&typeOfGame, &status)
	if err != nil {
		fmt.Println()
		return activeResJson, err
	}
	findPlayersIdsPerGame := fmt.Sprintf("SELECT user_id from game_players WHERE game_id = %d", id)
	rows, err := db.Query(findPlayersIdsPerGame)
	if err != nil {
		fmt.Println(err)
		return activeResJson, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&playerId)
		if err != nil {
			return activeRes, err
		}
		playersIds = append(playersIds, playerId)
	}
	mapPlayersIds := make(map[int]int)
	for playerIdIndex := 0; playerIdIndex < len(playersIds); playerIdIndex++ {
		mapPlayersIds[playersIds[playerIdIndex]] = playerIdIndex
	}
	if numOfRowsPerGame == 0 {
		update, err := db.Prepare("UPDATE games set status='In Progress' WHERE id = ?")
		if err != nil {
			fmt.Println(err)
			return activeResJson, err
		}
		_, err = update.Exec(id)
		if err != nil {
			fmt.Println(err)
			return activeResJson, err
		}
		activeRes.Round = 1
		activeRes.Throw = 1
		activeRes.PlayerId = playersIds[0]
	} else {
		findCurrentTurnInfo := fmt.Sprintf("SELECT rounds.round, game_players.user_id as player_id, scores.throw FROM scores INNER JOIN rounds ON scores.round_id = rounds.id INNER JOIN game_players ON scores.game_player_id = game_players.id WHERE rounds.game_id = %d ORDER BY scores.id DESC LIMIT 1 ;", id)
		row = db.QueryRow(findCurrentTurnInfo)
		err = row.Scan(&activeRes.Round, &activeRes.PlayerId, &activeRes.Throw)
		if err != nil {
			fmt.Println(err)
			return activeResJson, err
		}
		if typeOfGame == "High Score" && numOfRowsPerGame%(9*len(playersIds)) == 0 {
			activeRes.Round = 0
			activeRes.PlayerId = 0
			activeRes.Throw = 0

			update, err := db.Prepare("UPDATE games SET status = 'Completed' WHERE id = ?")
			if err != nil {
				fmt.Println(err)
				return activeResJson, err
			}
			_, err = update.Exec(id)
			if err != nil {
				fmt.Println(err)
				return activeResJson, err
			}
		} else {
			if numOfRowsPerGame%(3*len(playersIds)) == 0 {
				activeRes.Round = activeRes.Round + 1
			}
			if activeRes.Throw%3 == 0 {
				activeRes.Throw = 1
				test := mapPlayersIds[activeRes.PlayerId]
				if test == (len(playersIds) - 1) {
					activeRes.PlayerId = playersIds[0]
				} else {
					activeRes.PlayerId = playersIds[test+1]
				}
			} else {
				activeRes.Throw = activeRes.Throw + 1
			}
		}
		if status == "Completed" {
			activeRes.Round = 0
			activeRes.PlayerId = 0
			activeRes.Throw = 0
		}
	}
	activeResJson = types.ActiveStatus{
		GameId:   id,
		Round:    activeRes.Round,
		PlayerId: activeRes.PlayerId,
		Throw:    activeRes.Throw,
	}
	return activeResJson, nil
}
