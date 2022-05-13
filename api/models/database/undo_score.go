package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type lastScoreDetails struct {
	scoreId  int
	score    int
	validate string
}

// Delete Games From Games Table By Game Id
func UndoScore(db *sql.DB, id int) error {
	for i := 1; i <= 3; i++ {
		lastScoreDetails, err := FindLastScoreId(db, id, lastScoreDetails{})
		if err != nil {
			log.Println(err)
			return err
		}
		if lastScoreDetails.validate == "INVALID" && lastScoreDetails.score == 0 {
			err = DeleteScore(db, lastScoreDetails.scoreId)
			if err != nil {
				log.Println(err)
				return err
			}
			log.Println("Successfully Deleted", lastScoreDetails.scoreId, "ScoreId")
		} else {
			err = DeleteScore(db, lastScoreDetails.scoreId)
			if err != nil {
				log.Println(err)
				return err
			}
			log.Println("Successfully Deleted", lastScoreDetails.scoreId, "ScoreId")
			break
		}
	}
	return nil
}

// Find Last Inserted Scoreid
func FindLastScoreId(db *sql.DB, id int, lastScoreDetails lastScoreDetails) (lastScoreDetails, error) {
	findLastScoreId := fmt.Sprintf("SELECT id,score,is_valid FROM scores WHERE game_player_id = (SELECT id FROM game_players WHERE game_id = %d) ORDER BY id DESC LIMIT 1", id)
	lastScoreId := db.QueryRow(findLastScoreId)
	err := lastScoreId.Scan(&lastScoreDetails.scoreId, &lastScoreDetails.score, &lastScoreDetails.validate)
	if err != nil {
		log.Println(err)
		return lastScoreDetails, err
	}
	return lastScoreDetails, nil
}

// Delete Score'srow by Perticuler Score Id
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
