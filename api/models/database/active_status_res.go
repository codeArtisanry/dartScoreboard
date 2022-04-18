package models

import (
	"dartscoreboard/models/types"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Get active-status from for last round,player and dart to vertify with frontend URL
func GetActiveStatusRes(db *sql.DB, id int, activeRes types.ActiveStatus, players types.Extra) (types.ActiveStatus, error) {
	var activeResJson types.ActiveStatus
	var arrofplayers []int
	var count int
	var typeofgame string
	queryofcount := fmt.Sprintf("SELECT COUNT(scores.id) from scores inner join game_players on game_players.id=scores.game_player_id WHERE game_players.game_id = %d;", id)
	row := db.QueryRow(queryofcount)
	err := row.Scan(&players.Count)
	if err != nil {
		fmt.Println(err)
		return activeResJson, err
	}
	queryoftype := fmt.Sprintf("SELECT type from games WHERE id=%d", id)
	rowoftype := db.QueryRow(queryoftype)
	erroftype := rowoftype.Scan(&players.Type)
	if erroftype != nil {
		fmt.Println()
		return activeResJson, erroftype
	}
	queryofuserid := fmt.Sprintf("SELECT user_id from game_players WHERE game_id = %d", id)
	rows, err := db.Query(queryofuserid)
	if err != nil {
		fmt.Println(err)
		return activeResJson, erroftype
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&players.Player)
		if err != nil {
			return activeRes, err
		}
		playerJson := types.Extra{
			Count:  players.Count,
			Player: players.Player,
			Type:   players.Type,
		}
		arrofplayers = append(arrofplayers, playerJson.Player)
		count = playerJson.Count
		typeofgame = playerJson.Type
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
		activeRes.PlayerId = arrofplayers[0]

	} else {
		query := fmt.Sprintf("SELECT rounds.round , game_players.user_id as player_id, scores.throw FROM scores INNER JOIN rounds ON scores.round_id =rounds.id INNER JOIN game_players ON scores.game_player_id = game_players.id WHERE rounds.game_id = %d ORDER BY scores.id DESC LIMIT 1 ; ", id)
		row = db.QueryRow(query)
		err = row.Scan(&activeRes.Round, &activeRes.PlayerId, &activeRes.Throw)
		if err != nil {
			fmt.Println(err)
			return activeResJson, err
		}

		if typeofgame == "High Score" && count%(9*len(arrofplayers)) == 0 {
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
			if count%(3*len(arrofplayers)) == 0 {
				activeRes.Round = activeRes.Round + 1
			}
			if activeRes.Throw == 3 {
				activeRes.Throw = 1
			} else {
				activeRes.Throw = activeRes.Throw + 1
			}
			if count%3 == 0 {
				currentplayer = currentplayer % len(arrofplayers)
				nextplayer := arrofplayers[currentplayer]
				activeRes.PlayerId = nextplayer
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
