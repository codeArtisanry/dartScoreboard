package models

import (
	"dartscoreboard/models/types"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func GetStartGame(db *sql.DB, id int, gameRes types.GameResponse) (types.CurrentTurnInfo, error) {
	var (
		GetStartGame     types.CurrentTurnInfo
		activeRes        types.ActiveStatus
		ActivePlayerInfo types.ActivePlayerInfo
		Scoreboard       types.Scoreboard
	)
	Scoreboard, err := GetScoreboard(db, id)
	if err != nil {
		fmt.Println(err)
		return GetStartGame, err
	}
	query := fmt.Sprintf("SELECT id, name, type FROM games WHERE id = %d;", id)
	row := db.QueryRow(query)
	err = row.Scan(&gameRes.Id, &gameRes.Name, &gameRes.Type)
	if err != nil {
		return GetStartGame, err
	}
	ActiveStatus, err := GetActiveStatusRes(db, id, activeRes)
	if err != nil {
		fmt.Println(err)
		return GetStartGame, err
	}
	if ActiveStatus.Round == 0 {
		GetStartGame = types.CurrentTurnInfo{
			Id:               gameRes.Id,
			Name:             gameRes.Name,
			Type:             gameRes.Type,
			Round:            ActiveStatus.Round,
			Throw:            ActiveStatus.Throw,
			ActivePlayerInfo: &ActivePlayerInfo,
			Scoreboard:       Scoreboard,
		}
		return GetStartGame, nil
	}
	queryPlayer := fmt.Sprintf("SELECT id, first_name, last_name, email FROM users WHERE id = %d", ActiveStatus.PlayerId)
	rowsPlayers := db.QueryRow(queryPlayer)
	err = rowsPlayers.Scan(&ActivePlayerInfo.Id, &ActivePlayerInfo.FirstName, &ActivePlayerInfo.LastName, &ActivePlayerInfo.Email)
	if err != nil {
		return GetStartGame, err
	}
	GetStartGame = types.CurrentTurnInfo{
		Id:               gameRes.Id,
		Name:             gameRes.Name,
		Type:             gameRes.Type,
		Round:            ActiveStatus.Round,
		Throw:            ActiveStatus.Throw,
		ActivePlayerInfo: &ActivePlayerInfo,
		Scoreboard:       Scoreboard,
	}
	return GetStartGame, nil
}

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
		findLastRoundOfGame := fmt.Sprintf("SELECT round FROM rounds WHERE game_id = %d ORDER BY round DESC LIMIT 1", id)
		rowsLastRound := db.QueryRow(findLastRoundOfGame)
		err = rowsLastRound.Scan(&lastRound)
		if err != nil {
			if err == sql.ErrNoRows {
				lastRound = 1
			}
			fmt.Println(err)
		}
		fmt.Println(lastRound)
		for round := 1; round <= lastRound; round++ {
			findRoundTotal := fmt.Sprintf("SELECT IFNULL(SUM(scores.score),0) FROM scores WHERE scores.is_valid = 'VALID' AND round_id = (SELECT id FROM rounds WHERE round = %d AND game_id = %d) AND game_player_id = (SELECT id FROM game_players WHERE user_id = %d AND game_id = %d);", round, id, PlayerId, id)
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
	Scoreboard = types.Scoreboard{
		PlayersScore: PlayersRes}

	return Scoreboard, nil
}

func FoundWinner(db *sql.DB, id int) (types.Scoreboard, error) {
	var (
		Scoreboard types.Scoreboard
		first_name string
		last_name  string
		win        int
	)
	Winner := fmt.Sprintf("SELECT u.first_name,u.last_name, max_score.score as total from scores s left join (SELECT game_player_id, sum(scores.score) as score from scores GROUP BY game_player_id ) as max_score on max_score.game_player_id = s.game_player_id LEFT JOIN game_players gp on gp.id = s.game_player_id left join users u on u.id = gp.user_id  where s.round_id in (select round from rounds r where game_id= %d) GROUP BY s.game_player_id,s.round_id  ORDER by s.score DESC;", id)
	rowsPlayer := db.QueryRow(Winner)
	err := rowsPlayer.Scan(&first_name, &last_name, &win)
	if err != nil {
		return Scoreboard, err
	}
	Scoreboard = types.Scoreboard{
		Winner: first_name + last_name,
	}

	return Scoreboard, nil
}
