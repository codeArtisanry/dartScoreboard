package models

import (
	"dartscoreboard/models/types"
	"database/sql"
	"fmt"

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
		PlayerFullName := fmt.Sprintf("SELECT first_name,last_name  from users where id = %d;", PlayerId)
		rowsPlayer := db.QueryRow(PlayerFullName)
		err = rowsPlayer.Scan(&PlayerFirstName, &PlayerLastName)
		if err != nil {
			return Scoreboard, err
		}
		Total := fmt.Sprintf("select ifnull(sum(s.score),0) from scores s left join game_players gp on gp.id = s.game_player_id WHERE gp.game_id = %d AND gp.user_id = %d;", id, PlayerId)
		rowsPlayerTotal := db.QueryRow(Total)
		err = rowsPlayerTotal.Scan(&PlayerTotal)
		if err != nil {
			return Scoreboard, err
		}
		Round := fmt.Sprintf("SELECT round, SUM(s2.score) from scores s2 left join rounds r on s2.round_id = r.id  where s2.game_player_id = (SElect id from game_players gp where gp.user_id= %d and gp.game_id=%d) group by s2.round_id;", PlayerId, id)
		rowsRound, err := db.Query(Round)
		for rowsRound.Next() {
			err = rowsRound.Scan(&PerRound, &RoundTotal)
			if err != nil {
				return Scoreboard, err
			}
			dart := fmt.Sprintf("SELECT s.score from scores s join rounds r on s.round_id = r.id where r.round = %d AND game_player_id = (SELECT id FROM game_players WHERE game_id = %d AND user_id= %d) AND s.is_valid = 'VALID';", PerRound, id, PlayerId)
			rowsThrow, err := db.Query(dart)
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
		Scoreboard = types.Scoreboard{
			PlayersScore: PlayersRes,
		}
	}
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
