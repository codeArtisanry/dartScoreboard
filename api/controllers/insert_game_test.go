package controllers_test

import (
	"dartscoreboard/controllers"
	models "dartscoreboard/models/database"
	types "dartscoreboard/models/types"
	"fmt"
	"testing"
	"time"
)

type InsertGame struct {
	createrEmail    string
	user            types.User
	game            types.Game
	gameRes         types.GameResponse
	gamePlayer      types.GamePlayer
	gamePlayerRes   types.GamePlayerResponse
	gameResofInsert types.GameResponse
}

var InsertGameTest = []InsertGame{
	InsertGame{"payal@improwised.com", types.User{
		1,
		"Payal",
		"Raviya",
		"payal@improwised.com",
		"",
		&time.Time{},
		&time.Time{},
	}, types.Game{
		6,
		"Six",
		"Target Score-101",
		"Not Started",
		[]int{1},
		"payal@improwised.com",
		&time.Time{},
		&time.Time{},
	}, types.GameResponse{
		6,
		"Six",
		"Target Score-101",
		"Not Started",
		1,
		"Payal Raviya",
		[]types.GamePlayerResponse{{1, "Payal", "Raviya", "payal@improwised.com"}},
	}, types.GamePlayer{
		1,
		1,
		1,
		&time.Time{},
		&time.Time{},
	}, types.GamePlayerResponse{
		1, "Payal", "Raviya", "payal@improwised.com",
	}, types.GameResponse{
		6,
		"Six",
		"Target Score-101",
		"Not Started",
		1,
		"Payal Raviya",
		[]types.GamePlayerResponse{{1, "Payal", "Raviya", "payal@improwised.com"}},
	},
	},
}

func TestInsertGames(t *testing.T) {
	dbtest := models.Database("test.db")
	for _, test := range InsertGameTest {
		output, err := controllers.InsertGames(dbtest, test.createrEmail, test.user, test.game, test.gameRes, test.gamePlayer, test.gamePlayerRes)
		if err != nil {
			fmt.Println(output)
			t.Errorf("output is not same want")
			fmt.Println(err)
		}
	}
}
