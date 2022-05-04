package models

import (
	"dartscoreboard/models/types"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// get scoreboard
func GetScoreboard(db *sql.DB, id int) (types.Scoreboard, error) {
	var (
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
		PlayerRes       types.PlayerScore
		GameFullType    string
		lastRound       int
		targetScore     int
		count           int
	)
	findGameType := fmt.Sprintf("SELECT type FROM games where id = %d;", id)
	rowsPlayer := db.QueryRow(findGameType)
	err := rowsPlayer.Scan(&GameFullType)
	if err != nil {
		fmt.Println(err)
		return Scoreboard, err
	}
	GameType := strings.Split(GameFullType, "-")
	if len(GameType) == 1 {
		targetScore = 0
	} else {
		targetScore, err = strconv.Atoi(GameType[1])
		if err != nil {
			fmt.Println(err)
			return Scoreboard, err
		}
	}
	gameType := GameType[0]
	gamePlayers := fmt.Sprintf("SELECT user_id FROM game_players WHERE game_id = %d;", id)
	rowsPlayersIds, err := db.Query(gamePlayers)
	if err != nil {
		return Scoreboard, err
	}
	defer rowsPlayersIds.Close()
	for rowsPlayersIds.Next() {
		err = rowsPlayersIds.Scan(&PlayerId)
		if err != nil {
			return Scoreboard, err
		}
		PlayerFullName := fmt.Sprintf("SELECT first_name,last_name from users where id = %d;", PlayerId)
		rowsPlayer := db.QueryRow(PlayerFullName)
		err := rowsPlayer.Scan(&PlayerFirstName, &PlayerLastName)
		if err != nil {
			fmt.Println(err)
			return Scoreboard, err
		}
		Total := fmt.Sprintf("select ifnull(sum(s.score),0) from scores s left join game_players gp on gp.id = s.game_player_id WHERE s.is_valid = 'VALID' AND gp.game_id = %d AND gp.user_id = %d;", id, PlayerId)
		rowsPlayerTotal := db.QueryRow(Total)
		err = rowsPlayerTotal.Scan(&PlayerTotal)
		if err != nil {
			return Scoreboard, err
		}
		findCount := fmt.Sprintf("SELECT COUNT(round) FROM rounds WHERE game_id = %d", id)
		rowsLastRound := db.QueryRow(findCount)
		err = rowsLastRound.Scan(&count)
		if err != nil {
			fmt.Println(err)
			return Scoreboard, err
		}
		if count == 0 || count == 1 {
			lastRound = 1
		} else {
			findLastRoundOfGame := fmt.Sprintf("SELECT round FROM rounds WHERE game_id = %d ORDER BY round DESC LIMIT 1", id)
			rowsLastRound := db.QueryRow(findLastRoundOfGame)
			err = rowsLastRound.Scan(&lastRound)
			if err != nil {
				fmt.Println(err)
				return Scoreboard, err
			}
		}
		for round := 1; round <= lastRound; round++ {
			findRoundTotal := fmt.Sprintf("SELECT IFNULL(SUM(scores.score),0) FROM scores WHERE round_id = (SELECT id FROM rounds WHERE round = %d AND game_id = %d) AND game_player_id = (SELECT id FROM game_players WHERE user_id = %d AND game_id = %d);", round, id, PlayerId, id)
			roundTotal := db.QueryRow(findRoundTotal)
			err = roundTotal.Scan(&RoundTotal)
			if err != nil {
				fmt.Println(err)
				return Scoreboard, err
			}
			dart := fmt.Sprintf("SELECT s.score from scores s join rounds r on s.round_id = r.id where r.round = %d AND game_player_id = (SELECT id FROM game_players WHERE game_id = %d AND user_id= %d);", round, id, PlayerId)
			rowsThrow, err := db.Query(dart)
			if err != nil {
				fmt.Println(err)
				return Scoreboard, err
			}
			defer rowsThrow.Close()
			for rowsThrow.Next() {
				err = rowsThrow.Scan(&throwScore)
				if err != nil {
					return Scoreboard, err
				}
				Throws = append(Throws, throwScore)
			}
			RoundRes := types.Rounds{
				Round:       round,
				ThrowsScore: Throws,
				RoundTotal:  RoundTotal}
			RoundsRes = append(RoundsRes, RoundRes)
			Throws = nil
		}
		if gameType == "Target Score" {
			PlayerRes = types.PlayerScore{
				FirstName: PlayerFirstName,
				LastName:  PlayerLastName,
				Rounds:    RoundsRes,
				Total:     targetScore - PlayerTotal}
		} else {
			PlayerRes = types.PlayerScore{
				FirstName: PlayerFirstName,
				LastName:  PlayerLastName,
				Rounds:    RoundsRes,
				Total:     PlayerTotal}
		}
		PlayersRes = append(PlayersRes, PlayerRes)
		RoundsRes = nil
	}
	if count == 0 || count == 1 {
		Scoreboard = types.Scoreboard{
			PlayersScore: PlayersRes}
		return Scoreboard, nil
	} else {
		Winner, err := FoundWinner(db, id)
		if err != nil {
			fmt.Println(err)
			return Scoreboard, err
		}
		Scoreboard = types.Scoreboard{
			PlayersScore: PlayersRes,
			Winner:       Winner}
		return Scoreboard, nil
	}
}

func FoundWinner(db *sql.DB, id int) (string, error) {
	var (
		first_name string
		last_name  string
		winner     string
	)
	WinnerName := fmt.Sprintf("SELECT u.first_name,u.last_name from scores s join (SELECT game_player_id, sum(scores.score) as score from scores join rounds r ON r.id = scores.round_id AND r.game_id=%d AND scores.is_valid='VALID' GROUP BY game_player_id ) as max_score on max_score.game_player_id = s.game_player_id JOIN game_players gp on gp.id = s.game_player_id join users u on u.id = gp.user_id where round_id in (select id from rounds r WHERE r.game_id=%d) AND s.is_valid='VALID' GROUP BY gp.id ,s.round_id ORDER by max_score.score DESC LIMIT 1;", id, id)
	rowsPlayer := db.QueryRow(WinnerName)
	err := rowsPlayer.Scan(&first_name, &last_name)
	if err != nil {
		return winner, err
	}
	winner = first_name + last_name
	return winner, nil
}
