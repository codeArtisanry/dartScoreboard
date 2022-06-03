package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type LastScoreDetails struct {
	ScoreId  int
	Score    int
	Validate string
}
type PreviousTurn struct {
	Round    int
	PlayerId int
	Turn     int
}

// Find players of particular game by gameId
func FindGamePlayers(db *sql.DB, gameId int) ([]int, error) {
	var playersList []int
	var player int
	query := fmt.Sprintf("SELECT user_id FROM game_players WHERE game_id = %d;", gameId)
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return playersList, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&player)
		if err != nil {
			log.Println(err)
			return playersList, err
		}
		playersList = append(playersList, player)
	}
	return playersList, nil
}

// Find Last Inserted Scoreid
func FindLastScoreId(db *sql.DB, gameId int, lastScoreDetails LastScoreDetails, previosTurn PreviousTurn) (LastScoreDetails, error) {
	findLastScoreId := fmt.Sprintf("SELECT id,score,is_valid FROM scores WHERE game_player_id = (SELECT id FROM game_players WHERE game_id = %d AND user_id = %d) AND round_id = (SELECT id FROM rounds WHERE round = %d AND game_id = %d) AND throw = %d", gameId, previosTurn.PlayerId, previosTurn.Round, gameId, previosTurn.Turn)
	lastScoreId := db.QueryRow(findLastScoreId)
	err := lastScoreId.Scan(&lastScoreDetails.ScoreId, &lastScoreDetails.Score, &lastScoreDetails.Validate)
	if err != nil {
		log.Println(err)
		return lastScoreDetails, err
	}
	return lastScoreDetails, nil
}

// Delete last Score by Particular Score Id
func DeleteScore(db *sql.DB, scoreId int) error {
	delete, err := db.Prepare("DELETE FROM scores WHERE id = ?;")
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = delete.Exec(scoreId)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func ChangeValid(db *sql.DB, gameId int, previousTurn PreviousTurn) {
	var count int
	//db := database.Db
	countRemainingScore := fmt.Sprintf("SELECT count(score) from scores WHERE game_player_id = (SELECT id FROM game_players WHERE game_id = %d AND user_id = %d) AND round_id = (SELECT id FROM rounds WHERE game_id = %d AND round = %d);", gameId, previousTurn.PlayerId, gameId, previousTurn.Round)
	RemainingScore := db.QueryRow(countRemainingScore)
	err := RemainingScore.Scan(&count)
	if err != nil {
		log.Println(err)

	}
	if count != 0 {
		// Update Previous Turns Validation Score for Particular Round and Player to Valid
		invalidToValid, err := db.Prepare("UPDATE scores SET is_valid = ? WHERE game_player_id = (SELECT id FROM game_players WHERE game_id = ? AND user_id = ?) AND round_id = (SELECT id FROM rounds WHERE game_id = ? AND round = ?);")
		if err != nil {
			log.Println(err)
		}
		_, err = invalidToValid.Exec("VALID", gameId, previousTurn.PlayerId, gameId, previousTurn.Round)
		if err != nil {
			log.Println(err)
		}
	}
}
