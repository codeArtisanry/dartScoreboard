package models

import (
	"dartscoreboard/models/types"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type DataBase struct {
	Db *sql.DB
}

func Query(db *sql.DB, id int, activeRes types.ActiveStatus) (numOfRowsPerGame int, typeOfGame string, status string, playersIds []int) {
	var playerId int
	findNumOfRowsPerGame := fmt.Sprintf("SELECT COUNT(scores.id) from scores inner join game_players on game_players.id=scores.game_player_id WHERE game_players.game_id = %d;", id)
	row := db.QueryRow(findNumOfRowsPerGame)
	err := row.Scan(&numOfRowsPerGame)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println("12", numOfRowsPerGame)
	findGameInfo := fmt.Sprintf("SELECT type, status from games WHERE id=%d", id)
	rowoftype := db.QueryRow(findGameInfo)
	err = rowoftype.Scan(&typeOfGame, &status)
	if err != nil {
		fmt.Println()

	}
	findPlayersIdsPerGame := fmt.Sprintf("SELECT user_id from game_players WHERE game_id = %d", id)
	rows, err := db.Query(findPlayersIdsPerGame)
	if err != nil {
		fmt.Println(err)

	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&playerId)
		if err != nil {
			fmt.Println(err)
		}
		playersIds = append(playersIds, playerId)
	}
	return numOfRowsPerGame, typeOfGame, status, playersIds
}

func UpdateStatus(db *sql.DB, id int, status string) error {
	update, err := db.Prepare("UPDATE games set status=? WHERE id = ?")
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = update.Exec(status, id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}

func Find(db *sql.DB, id int, activeRes types.ActiveStatus) (Round int, PlayerId int, Throw int) {
	findCurrentTurnInfo := fmt.Sprintf("SELECT rounds.round, game_players.user_id as player_id, scores.throw FROM scores INNER JOIN rounds ON scores.round_id = rounds.id INNER JOIN game_players ON scores.game_player_id = game_players.id WHERE rounds.game_id = %d ORDER BY scores.id DESC LIMIT 1 ;", id)
	row := db.QueryRow(findCurrentTurnInfo)
	err := row.Scan(&activeRes.Round, &activeRes.PlayerId, &activeRes.Throw)
	Round = activeRes.Round
	PlayerId = activeRes.PlayerId
	Throw = activeRes.Throw
	if err != nil {
		fmt.Println(err)
	}
	return Round, PlayerId, Throw
}
