package models

import (
	"dartscoreboard/models/types"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Get active-status from for last round,player and dart to vertify with frontend URL
func GetActiveStatusRes(db *sql.DB, id int, activeRes types.ActiveStatus, players Extra) (ActiveStatus, error) {
	var activeResJson types.ActiveStatus

	var count int
	var type1 string

	query2 := fmt.Sprintf("SELECT  COUNT(scores.id) from scores inner join game_players on game_players.id=scores.player_id WHERE game_players.game_id = %d;", id)
	row2 := db.QueryRow(query2)
	err2 := row2.Scan(&players.Count)
	fmt.Println(players.Count)
	if err2 != nil {
		return activeResJson, err2
	}

	query5 := fmt.Sprintf("SELECT type  from games WHERE id=%d", id)
	row3 := db.QueryRow(query5)
	err3 := row3.Scan(&players.Type)
	fmt.Println(players.Type)
	if err3 != nil {
		return activeResJson, err3
	}

	var arrofplayers []int

	query3 := fmt.Sprintf("SELECT user_id from game_players WHERE game_id = %d", id)
	rows, err := db.Query(query3)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&players.Player)
		if err != nil {
			return activeRes, err
		}
		userJson := types.Extra{
			Count:  players.Count,
			Player: players.Player,
			Type:   players.Type,
		}
		arrofplayers = append(arrofplayers, userJson.Player)
		count = userJson.Count
		type1 = userJson.Type

	}
	fmt.Println(type1)

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
		fmt.Println("Sucessfully Data Inserted on users Table")
		activeRes.Round = 1
		activeRes.Throw = 1
		activeRes.PlayerId = arrofplayers[0]

	} else {

		query := fmt.Sprintf("SELECT rounds.round , game_players.user_id as player_id, scores.dart from game_players inner join rounds on game_players.game_id =rounds.game_id INNER  JOIN scores on scores.player_id = game_players.id WHERE game_players.game_id = %d GROUP  BY rounds.round ,game_players.id , scores.dart ORDER BY rounds.round DESC ,game_players.id DESC , scores.dart DESC LIMIT 1; ", id)
		row := db.QueryRow(query)
		err22 := row.Scan(&activeRes.Round, &activeRes.PlayerId, &activeRes.Throw)
		if err22 != nil {
			return activeResJson, err22
		}

		counter := activeRes.PlayerId
		fmt.Println(counter)
		if count%(3*len(arrofplayers)) == 0 {
			activeRes.Round = activeRes.Round + 1
		}
		if activeRes.Throw == 3 {
			activeRes.Throw = 1
		} else {
			activeRes.Throw = activeRes.Throw + 1
		}
		if count%3 == 0 {

			counter = counter % len(arrofplayers)
			arrofplayers1 := arrofplayers[counter]
			fmt.Println(arrofplayers1)
			activeRes.PlayerId = arrofplayers1
		}
	}
	activeResJson = ActiveStatus{
		GameId:   id,
		Round:    activeRes.Round,
		PlayerId: activeRes.PlayerId,
		Throw:    activeRes.Throw,
	}
	return activeResJson, nil
}
