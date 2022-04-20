package models

import (
	"dartscoreboard/models/types"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func InsertScore(db *sql.DB, gameId int, score types.Score) (types.ResScore, error) {
	var (
		totalScore   int
		roundId      int
		gamePlayerId int
		gameType     string
		scoreRes     types.ResScore
	)
	activeRes := types.ActiveStatus{}
	activejson, err := GetActiveStatusRes(db, gameId, activeRes)
	if err != nil {
		fmt.Println(err)
		return scoreRes, err
	}
	if activejson.Round == 0 {
		scoreRes := types.ResScore{
			Score:       0,
			TotalScore:  0,
			FoundWinner: true,
		}
		return scoreRes, nil
	}
	verifyRoundTable := fmt.Sprintf("SELECT id FROM rounds WHERE round = %d AND game_id = %d;", activejson.Round, gameId)
	row := db.QueryRow(verifyRoundTable)
	err = row.Scan(&roundId)
	if err != nil {
		if err == sql.ErrNoRows {
			insert, err := db.Prepare("INSERT INTO rounds(round, game_id) VALUES(?, ?)")
			if err != nil {
				fmt.Println(err)
				return scoreRes, err
			}
			_, err = insert.Exec(activejson.Round, gameId)
			if err != nil {
				fmt.Println(err)
				return scoreRes, err
			}
		}
	}
	findRoundId := fmt.Sprintf("SELECT id FROM rounds WHERE round = %d AND game_id = %d;", activejson.Round, gameId)
	row = db.QueryRow(findRoundId)
	err = row.Scan(&roundId)
	if err != nil {
		fmt.Println(err)
		return scoreRes, err
	}
	findGamePlayerid := fmt.Sprintf("SELECT id FROM game_players WHERE user_id = %d AND game_id = %d;", activejson.PlayerId, gameId)
	row = db.QueryRow(findGamePlayerid)
	err = row.Scan(&gamePlayerId)
	if err != nil {
		fmt.Println(3, err)
		return scoreRes, err
	}
	totalScore, err = FindTotalScore(db, gamePlayerId)
	if err != nil {
		fmt.Println(err)
		return scoreRes, err
	}
	findGameType := fmt.Sprintf("SELECT type FROM games WHERE id = %d;", gameId)
	row = db.QueryRow(findGameType)
	err = row.Scan(&gameType)
	if err != nil {
		fmt.Println(err)
		return scoreRes, err
	}
	totalScore = totalScore + score.Score
	if gameType == "High Score" {
		insert, err := db.Prepare("INSERT INTO scores(round_id, game_player_id, throw, score, is_valid) VALUES(?, ?, ?, ?, ?)")
		if err != nil {
			fmt.Println(err)
			return scoreRes, err
		}
		_, err = insert.Exec(roundId, gamePlayerId, activejson.Throw, score.Score, "VALID")
		if err != nil {
			fmt.Println(err)
			return scoreRes, err
		}
		scoreRes := types.ResScore{
			Score:       score.Score,
			TotalScore:  totalScore,
			FoundWinner: false,
		}
		return scoreRes, nil
	} else {
		IDENT := strings.Split(gameType, "-")
		targetscore, err := strconv.Atoi(IDENT[1])
		if err != nil {
			fmt.Println(err)
			return scoreRes, err
		}
		if totalScore <= targetscore {
			insert, err := db.Prepare("INSERT INTO scores(round_id, game_player_id, throw, score, is_valid) VALUES(?, ?, ?, ?, ?)")
			if err != nil {
				fmt.Println(err)
				return scoreRes, err
			}
			_, err = insert.Exec(roundId, gamePlayerId, activejson.Throw, score.Score, "VALID")
			if err != nil {
				fmt.Println(err)
				return scoreRes, err
			}
			if totalScore == targetscore {
				query, err := db.Prepare("UPDATE games SET status = ? WHERE id = ?;")
				if err != nil {
					fmt.Println(err)
					return scoreRes, err
				}
				_, err = query.Exec("Completed", gameId)
				if err != nil {
					fmt.Println(err)
					return scoreRes, err
				}
				scoreRes := types.ResScore{
					Score:       score.Score,
					TotalScore:  targetscore - totalScore,
					FoundWinner: true,
				}
				return scoreRes, nil
			}
			totalScore, err = FindTotalScore(db, gamePlayerId)
			if err != nil {
				fmt.Println(err)
				return scoreRes, err
			}
			scoreRes := types.ResScore{
				Score:       score.Score,
				TotalScore:  targetscore - totalScore,
				FoundWinner: false,
			}
			return scoreRes, nil
		} else {
			insert, err := db.Prepare("INSERT INTO scores(round_id, game_player_id, throw, score, is_valid) VALUES(?, ?, ?, ?, ?)")
			if err != nil {
				fmt.Println(err)
				return scoreRes, err
			}
			_, err = insert.Exec(roundId, gamePlayerId, activejson.Throw, score.Score, "VALID")
			if err != nil {
				fmt.Println(err)
				return scoreRes, err
			}
			remainingThrow := (3 - activejson.Throw)
			for throw := 1; throw <= remainingThrow; throw++ {
				insert, err := db.Prepare("INSERT INTO scores(round_id, game_player_id, throw, score, is_valid) VALUES(?, ?, ?, ?, ?)")
				if err != nil {
					fmt.Println(err)
					return scoreRes, err
				}
				_, err = insert.Exec(roundId, gamePlayerId, throw, 0, "VALID")
				if err != nil {
					fmt.Println(err)
					return scoreRes, err
				}
			}
			for throw := 1; throw <= 3; throw++ {
				query, err := db.Prepare("UPDATE scores SET is_valid = 'INVALID' WHERE round_id = ? AND game_player_id = ? AND throw = ?;")
				if err != nil {
					fmt.Println(err)
					return scoreRes, err
				}
				_, err = query.Exec(roundId, gamePlayerId, throw)
				if err != nil {
					fmt.Println(err)
					return scoreRes, err
				}
			}
			totalScore, err = FindTotalScore(db, gamePlayerId)
			if err != nil {
				fmt.Println(err)
				return scoreRes, err
			}
			scoreRes = types.ResScore{
				Score:       score.Score,
				TotalScore:  targetscore - totalScore,
				FoundWinner: false,
			}
			return scoreRes, nil
		}
	}
}

func FindTotalScore(db *sql.DB, gamePlayerId int) (int, error) {
	var totalScore int
	findTotalScore := fmt.Sprintf("SELECT IFNULL(sum(scores.score),0) from scores where game_player_id = %d AND is_valid = 'VALID';", gamePlayerId)
	row := db.QueryRow(findTotalScore)
	err := row.Scan(&totalScore)
	if err != nil {
		fmt.Println(err)
		return totalScore, err
	}
	return totalScore, nil
}
