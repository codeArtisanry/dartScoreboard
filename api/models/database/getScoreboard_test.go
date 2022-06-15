package models_test

import (
	models "dartscoreboard/models/database"
	types "dartscoreboard/models/types"
	"fmt"
	"reflect"
	"testing"
)

type ScoreboardTest struct {
	gameId     int
	Scoreboard types.Scoreboard
}

var ScoreboardTests = []ScoreboardTest{
	ScoreboardTest{1, types.Scoreboard{
		PlayersScore: []types.PlayerScore{{
			FirstName: "Payal",
			LastName:  "Raviya",
			Rounds: []types.Rounds{{
				Round:       1,
				ThrowsScore: nil,
				CheckRound:  "VALID",
				RoundTotal:  0,
			}},
			Total: 0,
		}},
		Winner: "",
	}},
	ScoreboardTest{2, types.Scoreboard{
		PlayersScore: []types.PlayerScore{{
			FirstName: "Payal",
			LastName:  "Raviya",
			Rounds: []types.Rounds{{
				Round:       1,
				ThrowsScore: []int{11, 11, 11},
				CheckRound:  "VALID",
				RoundTotal:  33,
			},
				{
					Round:       2,
					ThrowsScore: []int{11, 11, 11},
					CheckRound:  "VALID",
					RoundTotal:  33,
				},
				{
					Round:       3,
					ThrowsScore: []int{11, 11, 11},
					CheckRound:  "VALID",
					RoundTotal:  33,
				},
			},
			Total: 99,
		}},
		Winner: "Payal  Raviya",
	}},
}

func TestGetScoreboard(t *testing.T) {
	dbtest := models.Database("testmodel.db")
	resultOfUserQuery, err := dbtest.Exec("INSERT INTO users(id, first_name, last_name,email) VALUES(1,'Payal','Raviya','payal@improwised.com'),(2,'Jeel','Rupapara','jeel@improwised.com')")
	if err != nil {
		fmt.Println(err, resultOfUserQuery)
	}
	resultOfInsertGame, err := dbtest.Exec("INSERT INTO games(name, type, status, creater_user_id) VALUES('First', 'High Score', 'Not Started', 1),('Second', 'High Score', 'Not Started', 1)")
	if err != nil {
		fmt.Println(err, resultOfInsertGame)
	}
	resultOfInsertGamePlayer, err := dbtest.Exec("INSERT INTO game_players(id, user_id, game_id) VALUES(1,1,1),(2,1,2)")
	if err != nil {
		fmt.Println(err, resultOfInsertGamePlayer)
	}
	resultOfInsertScoreFourthGame, err := dbtest.Exec("INSERT INTO scores(round_id, game_player_id, throw, score,is_valid) VALUES(1,2,1,11,'VALID'),(1,2,2,11,'VALID'),(1,2,3,11,'VALID'),(2,2,1,11,'VALID'),(2,2,2,11,'VALID'),(2,2,3,11,'VALID'),(3,2,1,11,'VALID'),(3,2,2,11,'VALID'),(3,2,3,11,'VALID')")
	if err != nil {
		fmt.Println(err, resultOfInsertScoreFourthGame)
	}
	resultOfInsertRounds, err := dbtest.Exec("INSERT INTO rounds(round,game_id) VALUES(1,2),(2,2),(3,2)")
	if err != nil {
		fmt.Println(err, resultOfInsertRounds)
	}
	for _, test := range ScoreboardTests {
		output, err := models.GetScoreboard(dbtest, test.gameId)
		if !reflect.DeepEqual(output, test.Scoreboard) {
			t.Errorf("output is not same want %v,%v", output, test.Scoreboard)
			fmt.Println(err)
		}
	}
}
