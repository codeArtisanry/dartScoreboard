package models

import (
	"dartscoreboard/models/types"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Get active-status from for last round,player and dart to vertify with frontend URL
func GetActiveStatusRes(db *sql.DB, id int, activeRes types.ActiveStatus) (types.ActiveStatus, error) {
	var (
		activeResJson types.ActiveStatus
		playersIds  []int
		count         int
		typeOfGame    string
		status        string
		playerId      int
	)
	queryofcount := fmt.Sprintf("SELECT COUNT(scores.id) from scores inner join game_players on game_players.id=scores.game_player_id WHERE game_players.game_id = %d;", id)
	row := db.QueryRow(queryofcount)
	err := row.Scan(&count)
	if err != nil {
		fmt.Println(err)
		return activeResJson, err
	}
	queryoftype := fmt.Sprintf("SELECT type, status from games WHERE id=%d", id)
	rowoftype := db.QueryRow(queryoftype)
	err = rowoftype.Scan(&typeOfGame, &status)
	if err != nil {
		fmt.Println()
		return activeResJson, err
	}
	queryofuserid := fmt.Sprintf("SELECT user_id from game_players WHERE game_id = %d", id)
	rows, err := db.Query(queryofuserid)
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
	if count == 0 {
		insert, err := db.Prepare("UPDATE games set status='In Progress' WHERE id = ?")
		if err != nil {
			fmt.Println(err)
			return activeResJson, err
		}
		_, err = insert.Exec(id)
		if err != nil {
			fmt.Println(err)
			return activeResJson, err
		}
		activeRes.Round = 1
		activeRes.Throw = 1
		activeRes.PlayerId = playersIds[0]

	} else {
		query := fmt.Sprintf("SELECT rounds.round, game_players.user_id as player_id, scores.throw FROM scores INNER JOIN rounds ON scores.round_id = rounds.id INNER JOIN game_players ON scores.game_player_id = game_players.id WHERE rounds.game_id = %d ORDER BY scores.id DESC LIMIT 1 ;", id)
		row = db.QueryRow(query)
		err = row.Scan(&activeRes.Round, &activeRes.PlayerId, &activeRes.Throw)
		if err != nil {
			fmt.Println(err)
			return activeResJson, err
		}

		if typeOfGame == "High Score" && count%(9*len(playersIds)) == 0 {
			activeRes.Round = 0
			activeRes.PlayerId = 0
			activeRes.Throw = 0

			insert, err := db.Prepare("UPDATE games SET status = 'Completed' WHERE id = ?")
			if err != nil {
				fmt.Println(err)
				return activeResJson, err
			}
			_, err = insert.Exec(id)
			if err != nil {
				fmt.Println(err)
				return activeResJson, err
			}
		} else {
			currentplayer := activeRes.PlayerId
			if count%(3*len(playersIds)) == 0 {
				activeRes.Round = activeRes.Round + 1
			}
			if activeRes.Throw == 3 {
				activeRes.Throw = 1
			} else {
				activeRes.Throw = activeRes.Throw + 1
			}
			if count%3 == 0 {
				currentplayer = currentplayer % len(playersIds)
				nextplayer := playersIds[currentplayer]
				activeRes.PlayerId = nextplayer
			}
			if status == "Completed" {
				activeRes.Round = 0
				activeRes.PlayerId = 0
				activeRes.Throw = 0
			}
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
