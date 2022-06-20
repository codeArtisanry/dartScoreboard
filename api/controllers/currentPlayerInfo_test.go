package controllers

import (
	models "dartscoreboard/models/database"
	types "dartscoreboard/models/types"
	"fmt"
	"reflect"
	"testing"
)

type CurrentPlayerInfoTest struct {
	gameId        int
	playerid      int
	gameRes       types.GameResponse
	CurrentPlayer types.CurrentPlayerInfo
}

var CurrentPlayerInfoTests = []CurrentPlayerInfoTest{
	CurrentPlayerInfoTest{4, 1, types.GameResponse{
		Id:            4,
		Name:          "Fourth",
		Type:          "High Score",
		Status:        "Completed",
		CreaterUserId: 1,
		CreaterName:   "Payal Raviya",
		Players: []types.GamePlayerResponse{{
			Id:        1,
			FirstName: "Payal",
			LastName:  "Raviya",
			Email:     "payal@improwised.com",
		}},
	},
		types.CurrentPlayerInfo{
			Id:    1,
			Name:  "Payal Raviya",
			Email: "payal@improwised.com",
			Round: 0,
			Throw: 0,
			Score: 99,
			Game: types.GameResponse{
				Id:            4,
				Name:          "Fourth",
				Type:          "High Score",
				Status:        "Completed",
				CreaterUserId: 1,
				CreaterName:   "Payal Raviya",
				Players: []types.GamePlayerResponse{{
					Id:        1,
					FirstName: "Payal",
					LastName:  "Raviya",
					Email:     "payal@improwised.com",
				}},
			},
		},
	},
}

func TestGetCurrentPlayerInfo(t *testing.T) {
	dbtest := models.Database("test.db")
	for _, test := range CurrentPlayerInfoTests {
		output, err := GetCurrentPlayerInfo(dbtest, test.gameId, test.playerid, test.gameRes)
		if !reflect.DeepEqual(output, test.CurrentPlayer) {
			t.Errorf("output is not same want %v,%v", output, test.CurrentPlayer)
			fmt.Println(err)
		}
	}

}
