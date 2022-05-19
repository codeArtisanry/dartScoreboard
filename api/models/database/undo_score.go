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
type PreviousTurn struct {
	round    int
	playerId int
	turn     int
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

// Find Previous turn, playerId and round
func FindPreviousTurn(round int, playerId int, turn int, playersId []int) PreviousTurn {
	var currentPlayerIndex int
	mapPlayerid := make(map[int]int)
	for i := 0; i <= len(playersId)-1; i++ {
		mapPlayerid[playersId[i]] = i
	}
	turn = turn - 1
	if turn == 0 {
		turn = 3
		currentPlayerIndex = mapPlayerid[playerId]
		currentPlayerIndex = currentPlayerIndex - 1
		if currentPlayerIndex == -1 {
			currentPlayerIndex = len(playersId) - 1
			playerId = playersId[currentPlayerIndex]
			round = round - 1
		} else {
			playerId = playersId[currentPlayerIndex]
		}
	}
	previousTurn := PreviousTurn{
		round:    round,
		playerId: playerId,
		turn:     turn,
	}
	return previousTurn
}

// Find Last Inserted Scoreid
func FindLastScoreId(db *sql.DB, gameId int, lastScoreDetails lastScoreDetails, previosTurn PreviousTurn) (lastScoreDetails, error) {
	findLastScoreId := fmt.Sprintf("SELECT id,score,is_valid FROM scores WHERE game_player_id = (SELECT id FROM game_players WHERE game_id = %d AND user_id = %d) AND round_id = (SELECT id FROM rounds WHERE round = %d AND game_id = %d) AND throw = %d", gameId, previosTurn.playerId, previosTurn.round, gameId, previosTurn.turn)
	lastScoreId := db.QueryRow(findLastScoreId)
	err := lastScoreId.Scan(&lastScoreDetails.scoreId, &lastScoreDetails.score, &lastScoreDetails.validate)
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

// Delete Last Inserted Score By GameId for undo action
func UndoScore(db *sql.DB, gameId int, round int, playerId int, turn int) error {
	gamePlayerList, err := FindGamePlayers(db, gameId)
	if err != nil {
		log.Println(err)
		return err
	}
	previousTurn := FindPreviousTurn(round, playerId, turn, gamePlayerList)
	for i := 1; i <= 3; i++ {
		lastScoreDetails, err := FindLastScoreId(db, gameId, lastScoreDetails{}, previousTurn)
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
			// Find Previous Turns of that player in that particular round
			var count int
			countRemainingScore := fmt.Sprintf("SELECT count(score) from scores WHERE game_player_id = (SELECT id FROM game_players WHERE game_id = %d AND user_id = %d) AND round_id = (SELECT id FROM rounds WHERE game_id = %d AND round = %d);", gameId, previousTurn.playerId, gameId, previousTurn.round)
			RemainingScore := db.QueryRow(countRemainingScore)
			err := RemainingScore.Scan(&count)
			if err != nil {
				log.Println(err)
				return err
			}
			if count != 0 {
				// Update Previous Turns Validation Score for Particular Round and Player to Valid
				invalidToValid, err := db.Prepare("UPDATE scores SET is_valid = ? WHERE game_player_id = (SELECT id FROM game_players WHERE game_id = ? AND user_id = ?) AND round_id = (SELECT id FROM rounds WHERE game_id = ? AND round = ?);")
				if err != nil {
					log.Println(err)
					return err
				}
				_, err = invalidToValid.Exec("VALID", gameId, previousTurn.playerId, gameId, previousTurn.round)
				if err != nil {
					log.Println(err)
					return err
				}
			}
			break
		}
		previousTurn = FindPreviousTurn(previousTurn.round, previousTurn.playerId, previousTurn.turn, gamePlayerList)
	}
	return nil
}
