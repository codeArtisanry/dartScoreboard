package models

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func InsertScore(db *sql.DB, gameId int) {
	// login to next
	var totalScore int
	var roundId int
	var gamePlayerId int
	var gameType string
	round := 1
	findRoundId := fmt.Sprintf("SELECT id FROM rounds WHERE round = %d AND game_id = %d;", round, gameId)
	row := db.QueryRow(findRoundId)
	err := row.Scan(&roundId)
	if err != nil {
		fmt.Println(err)
	}
	playerid := 2
	findGamePlayerid := fmt.Sprintf("SELECT id FROM game_players WHERE user_id = %d AND game_id = %d);", playerid, gameId)
	row = db.QueryRow(findGamePlayerid)
	err = row.Scan(&gamePlayerId)
	if err != nil {
		fmt.Println(err)
	}
	current_score := 40
	findTotal := fmt.Sprintf("SELECT sum(scores.score) from scores where game_player_id = %d;", gamePlayerId)
	row = db.QueryRow(findTotal)
	err = row.Scan(&totalScore)
	if err != nil {
		fmt.Println(err)
	}
	findGameType := fmt.Sprintf("SELECT sum(scores.score) from scores where game_player_id = %d;", gamePlayerId)
	row = db.QueryRow(findGameType)
	err = row.Scan(&gameType)
	if err != nil {
		fmt.Println(err)
	}
	IDENT := strings.Split(gameType, "-")
	targetscore, err := strconv.Atoi(IDENT[1])
	if err != nil {
		fmt.Println(err)
	}
	totalScore = totalScore + current_score
	if totalScore <= targetscore {
		// insert()
		if totalScore == targetscore {
			// responce winner
		}
		//response total score, score
	} else {
		// insert()
		var throw int
		remainingThrow := (3 - throw)
		for i:=1; i<=remainingThrow; i++{
			// insert by throw
		}
		for j:=1; j<=3; j++{
			// update by throw
		}
		// response total score, score
	}
}
