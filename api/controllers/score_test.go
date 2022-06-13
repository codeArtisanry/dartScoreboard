package controllers_test

import (
	"dartscoreboard/controllers"
	models "dartscoreboard/models/database"
	types "dartscoreboard/models/types"
	"fmt"
	"testing"
)

type InserScoreTest struct {
	gameId int

	Playerid   int
	round      int
	TurnId     int
	Score      types.Score
	ResOfScore types.ResScore
}

var InserScore = []InserScoreTest{
	InserScoreTest{1, 1, 1, 1, types.Score{43}, types.ResScore{43, 43, false}},
	InserScoreTest{2, 2, 1, 1, types.Score{43}, types.ResScore{43, 76, false}},
	InserScoreTest{5, 1, 1, 2, types.Score{43}, types.ResScore{43, 43, false}},
}

func TestInsertScore(t *testing.T) {
	dbtest := models.Database("test.db")
	resultOfInsertGame, err := dbtest.Exec("INSERT INTO games(name, type, status, creater_user_id) VALUES('Five', 'Target Score-101', 'Not Started', 1)")
	if err != nil {
		fmt.Println(err, resultOfInsertGame)
	}
	resultOfInsertGamePlayer, err := dbtest.Exec("INSERT INTO game_players(id, user_id, game_id) VALUES(7,1,5)")
	if err != nil {
		fmt.Println(err, resultOfInsertGamePlayer)
	}
	resultOfInsertScore, err := dbtest.Exec("INSERT INTO scores(round_id, game_player_id, throw,score) VALUES(7,7,1,43)")
	if err != nil {
		fmt.Println(err, resultOfInsertScore)
	}

	for _, test := range InserScore {
		output, err := controllers.InsertScore(dbtest, test.gameId, test.Playerid, test.round, test.TurnId, test.Score)
		if err != nil {
			t.Errorf("output is not same want")
			fmt.Println(output)
			fmt.Println(err)
		}
	}

}
