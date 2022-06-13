package controllers

import (
	models "dartscoreboard/models/database"
	types "dartscoreboard/models/types"
	"fmt"
	"testing"
)

type ActiveStatusTest struct {
	gameId int

	ActiveRes    types.ActiveStatus
	ActiveStatus types.ActiveStatus
}

var ActiveStatusTests = []ActiveStatusTest{
	ActiveStatusTest{1, types.ActiveStatus{1, 0, 1, 0}, types.ActiveStatus{1, 1, 1, 1}},
	ActiveStatusTest{2, types.ActiveStatus{2, 1, 1, 3}, types.ActiveStatus{2, 1, 2, 1}},
	ActiveStatusTest{3, types.ActiveStatus{3, 1, 2, 3}, types.ActiveStatus{3, 2, 1, 1}},
	ActiveStatusTest{4, types.ActiveStatus{4, 3, 1, 3}, types.ActiveStatus{4, 0, 0, 0}},
}

func TestGetActiveStatusRes(t *testing.T) {
	dbtest := models.Database("test.db")
	resultOfUserQuery, err := dbtest.Exec("INSERT INTO users(id, first_name, last_name,email) VALUES(1,'Payal','Raviya','payal@improwised.com'),(2,'Jeel','Rupapara','jeel@improwised.com')")
	if err != nil {
		fmt.Println(err, resultOfUserQuery)
	}
	resultOfInsertGame, err := dbtest.Exec("INSERT INTO games(name, type, status, creater_user_id) VALUES('First', 'High Score', 'Not Started', 1),('Second', 'High Score', 'Not Started', 1),('Third', 'High Score', 'Not Started', 1),('Fourth', 'High Score', 'Not Started', 1)")
	if err != nil {
		fmt.Println(err, resultOfInsertGame)
	}
	resultOfInsertGamePlayer, err := dbtest.Exec("INSERT INTO game_players(id, user_id, game_id) VALUES(1,1,1),(2,1,2),(3,2,2),(4,1,3),(5,2,3),(6,1,4)")
	if err != nil {
		fmt.Println(err, resultOfInsertGamePlayer)
	}
	resultOfInsertScore, err := dbtest.Exec("INSERT INTO scores(round_id, game_player_id, throw, score) VALUES(1, 2, 1, 11),(2, 2, 2, 11),(3, 2, 3, 11),(4, 4, 1, 11),(5, 4, 2, 11),(6, 4, 3, 11),(7, 5, 1, 11),(8, 5, 2, 11),(9, 5, 3, 11)")
	if err != nil {
		fmt.Println(err, resultOfInsertScore)
	}
	resultOfInsertScoreFourthGame, err := dbtest.Exec("INSERT INTO scores(round_id, game_player_id, throw, score) VALUES(10, 6, 1, 11),(11, 6, 2, 11),(12, 6, 3, 11),(13, 6, 1, 11),(14, 6, 2, 11),(15, 6, 3, 11),(16, 6, 1, 11),(17, 6, 2, 11),(18, 6, 3, 11)")
	if err != nil {
		fmt.Println(err, resultOfInsertScoreFourthGame)
	}

	for _, test := range ActiveStatusTests {
		output, err := GetActiveStatusRes(dbtest, test.gameId, test.ActiveRes)
		if output != test.ActiveStatus {
			t.Errorf("output is not same want %d,%d", test.ActiveStatus, output)
			fmt.Println(err)
		}
	}
}
