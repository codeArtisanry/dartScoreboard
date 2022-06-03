package models

import (
	"dartscoreboard/models/types"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func (database DataBase) FindTotalScore(gamePlayerId int) (int, error) {
	var totalScore int
	db := database.Db
	findTotalScore := fmt.Sprintf("SELECT IFNULL(sum(scores.score),0) from scores where game_player_id = %d AND is_valid = 'VALID';", gamePlayerId)
	row := db.QueryRow(findTotalScore)
	err := row.Scan(&totalScore)
	if err != nil {
		fmt.Println(err)
		return totalScore, err
	}
	return totalScore, nil
}

func (database DataBase) VerifyRoundTableQuery(gameId int, round int, roundId int) {
	db := database.Db
	verifyRoundTable := fmt.Sprintf("SELECT id FROM rounds WHERE round = %d AND game_id = %d;", round, gameId)
	row := db.QueryRow(verifyRoundTable)
	err := row.Scan(&roundId)
	if err != nil {
		if err == sql.ErrNoRows {
			insert, err := db.Prepare("INSERT INTO rounds(round, game_id) VALUES(?, ?)")
			if err != nil {
				fmt.Println(err)
			}
			_, err = insert.Exec(round, gameId)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func (database DataBase) RoundGamePlayerIdQuery(gameId int, playerId int, round int) (gamePlayerId int, roundId int) {
	db := database.Db
	findRoundId := fmt.Sprintf("SELECT id FROM rounds WHERE round = %d AND game_id = %d;", round, gameId)
	row := db.QueryRow(findRoundId)
	err := row.Scan(&roundId)
	if err != nil {
		fmt.Println(err)
	}
	findGamePlayerid := fmt.Sprintf("SELECT id FROM game_players WHERE user_id = %d AND game_id = %d;", playerId, gameId)
	row = db.QueryRow(findGamePlayerid)
	err = row.Scan(&gamePlayerId)
	if err != nil {
		fmt.Println(3, err)
	}
	return gamePlayerId, roundId
}

func (database DataBase) FindGameTypeQuery(gameId int) (gameType string) {
	db := database.Db
	findGameType := fmt.Sprintf("SELECT type FROM games WHERE id = %d;", gameId)
	row := db.QueryRow(findGameType)
	err := row.Scan(&gameType)
	if err != nil {
		fmt.Println(err)
	}
	return gameType
}

func (database DataBase) InsertIntoScoreTableQuery(playerId int, round int, turnId int, score types.Score, roundId int, gamePlayerId int) {
	db := database.Db
	insert, err := db.Prepare("INSERT INTO scores(round_id, game_player_id, throw, score, is_valid) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
	}
	_, err = insert.Exec(roundId, gamePlayerId, turnId, score.Score, "VALID")
	if err != nil {
		fmt.Println(err)
	}
}

func (database DataBase) ValidateScoreQuery(gameId int, playerId int, round int, turnId int) (rowScore *sql.Row) {
	db := database.Db
	validateScore := fmt.Sprintf("SELECT id FROM scores WHERE game_player_id = (SELECT id FROM game_players WHERE game_id = %d AND user_id = %d) AND round_id = (SELECT id FROM rounds WHERE round = %d AND game_id = %d) AND throw = %d", gameId, playerId, round, gameId, turnId)
	rowScore = db.QueryRow(validateScore)
	return rowScore
}

func (database DataBase) QueryForUpdateIsValid(roundId int, gamePlayerId int, throw int) {
	db := database.Db
	query, err := db.Prepare("UPDATE scores SET is_valid = 'INVALID' WHERE round_id = ? AND game_player_id = ? AND throw = ?;")
	if err != nil {
		fmt.Println(err)
	}
	_, err = query.Exec(roundId, gamePlayerId, throw)
	if err != nil {
		fmt.Println(err)
	}
}

func (database DataBase) RemoveMultipleEntryInScore(roundId int, gamePlayerId int, throw int) {
	db := database.Db
	insert, err := db.Prepare("INSERT INTO scores(round_id, game_player_id, throw, score, is_valid) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
	}
	_, err = insert.Exec(roundId, gamePlayerId, throw, 0, "VALID")
	if err != nil {
		fmt.Println(err)
	}
}
