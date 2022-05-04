package models

import (
	"dartscoreboard/models/types"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// get scoreboard
func GetScoreboard(db *sql.DB, id int) (types.Scoreboard, error) {
	var (
		PerRound        int
		PlayerId        int
		throwScore      int
		PlayerFirstName string
		PlayerLastName  string
		PlayerTotal     int
		RoundTotal      int
		Scoreboard      types.Scoreboard
		Throws          []int
		RoundsRes       []types.Rounds
		PlayersRes      []types.PlayerScore
	)
	game := fmt.Sprintf("SELECT user_id FROM game_players WHERE game_id = %d;", id)
	rowsPlayersIds, err := db.Query(game)
	if err != nil {
		return Scoreboard, err
	}
	for rowsPlayersIds.Next() {
		err = rowsPlayersIds.Scan(&PlayerId)
		if err != nil {
			return Scoreboard, err
		}
		PlayerFullName := fmt.Sprintf("SELECT first_name,last_name FROM users where id = %d;", PlayerId)
		rowsPlayer := db.QueryRow(PlayerFullName)
		err = rowsPlayer.Scan(&PlayerFirstName, &PlayerLastName)
		if err != nil {
			return Scoreboard, err
		}
		Total := fmt.Sprintf("select ifnull(sum(s.score),0) from scores s left join game_players gp on gp.id = s.game_player_id WHERE gp.game_id = %d AND gp.user_id = %d  AND s.is_valid = 'VALID';", id, PlayerId)
		rowsPlayerTotal := db.QueryRow(Total)
		err = rowsPlayerTotal.Scan(&PlayerTotal)
		if err != nil {
			return Scoreboard, err
		}
		Round := fmt.Sprintf("SELECT round, SUM(s2.score) from scores s2 left join rounds r on s2.round_id = r.id  where s2.game_player_id = (SElect id from game_players gp where gp.user_id= %d and gp.game_id=%d) group by s2.round_id;", PlayerId, id)
		rowsRound, err := db.Query(Round)
		if err != nil {
			return Scoreboard, err
		}
		for rowsRound.Next() {
			err = rowsRound.Scan(&PerRound, &RoundTotal)
			if err != nil {
				return Scoreboard, err
			}
			dart := fmt.Sprintf("SELECT s.score from scores s join rounds r on s.round_id = r.id where r.round = %d AND game_player_id = (SELECT id FROM game_players WHERE game_id = %d AND user_id= %d) AND s.is_valid = 'VALID';", PerRound, id, PlayerId)
			rowsThrow, err := db.Query(dart)
			if err != nil {
				return Scoreboard, err
			}
			for rowsThrow.Next() {
				err = rowsThrow.Scan(&throwScore)
				if err != nil {
					return Scoreboard, err
				}
				Throws = append(Throws, throwScore)
			}
			RoundRes := types.Rounds{
				Round:       PerRound,
				ThrowsScore: Throws,
				RoundTotal:  RoundTotal,
			}
			Throws = nil
			RoundsRes = append(RoundsRes, RoundRes)
		}
		PlayerRes := types.PlayerScore{
			FirstName: PlayerFirstName,
			LastName:  PlayerLastName,
			Rounds:    RoundsRes,
			Total:     PlayerTotal,
		}
		RoundsRes = nil
		PlayersRes = append(PlayersRes, PlayerRes)
		Winner, err := FoundWinner(db, id)
		if err != nil {
			fmt.Println(err)
		}
		Scoreboard = types.Scoreboard{
			PlayersScore: PlayersRes,
			Winner:       Winner,
		}
	}
	return Scoreboard, nil
}

func FoundWinner(db *sql.DB, id int) (string, error) {
	var (
		first_name   string
		last_name    string
		winner       string
		fullGameType string
	)
	findGameType := fmt.Sprintf("SELECT type FROM games WHERE id = %d;", id)
	rowGameType := db.QueryRow(findGameType)
	err := rowGameType.Scan(&fullGameType)
	if err != nil {
		return winner, err
	}
	if fullGameType == "High Score" {
		WinnerName := fmt.Sprintf("SELECT u.first_name,u.last_name from scores s join (SELECT game_player_id, sum(scores.score) as score from scores join rounds r ON r.id = scores.round_id AND r.game_id=%d AND scores.is_valid='VALID' GROUP BY game_player_id ) as max_score on max_score.game_player_id = s.game_player_id JOIN game_players gp on gp.id = s.game_player_id join users u on u.id = gp.user_id where round_id in (select id from rounds r WHERE r.game_id=%d) AND s.is_valid='VALID' GROUP BY gp.id ,s.round_id ORDER by max_score.score DESC LIMIT 1;", id, id)
		rowsPlayer := db.QueryRow(WinnerName)
		err = rowsPlayer.Scan(&first_name, &last_name)
		if err != nil {
			return winner, err
		}
		winner = first_name + last_name
		return winner, nil
	} else {
		WinnerName := fmt.Sprintf("SELECT u.first_name,u.last_name from scores s join (SELECT game_player_id, sum(scores.score) as score from scores join rounds r ON r.id = scores.round_id AND r.game_id=%d AND scores.is_valid='VALID' GROUP BY game_player_id ) as max_score on max_score.game_player_id = s.game_player_id JOIN game_players gp on gp.id = s.game_player_id join users u on u.id = gp.user_id where round_id in (select id from rounds r WHERE r.game_id=%d) AND s.is_valid='VALID' GROUP BY gp.id ,s.round_id ORDER by max_score.score ASC LIMIT 1;", id, id)
		rowsPlayer := db.QueryRow(WinnerName)
		err = rowsPlayer.Scan(&first_name, &last_name)
		if err != nil {
			return winner, err
		}
		winner = first_name + last_name
		return winner, nil
	}
}