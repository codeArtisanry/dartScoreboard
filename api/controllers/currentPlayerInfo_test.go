package controllers

import (
	models "dartscoreboard/models/database"
	types "dartscoreboard/models/types"
	"fmt"
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
		4,
		"Fourth",
		"High Score",
		"Completed",
		1,
		"Payal Raviya",
		[]types.GamePlayerResponse{{
			1,
			"Payal",
			"Raviya",
			"payal@improwised.com",
		},
		},
	},
		types.CurrentPlayerInfo{
			4,
			"Fourth",
			"High Score",
			"",
			0,
			0,
			&types.ActivePlayerInfo{
				1,
				"Payal",
				"Raviya",
				"payal@improwised.com",
				99,
			},
		},
	},
}

func TestGetCurrentPlayerInfo(t *testing.T) {
	dbtest := models.Database("test.db")
	for _, test := range CurrentPlayerInfoTests {
		output, err := GetCurrentPlayerInfo(dbtest, test.gameId, test.playerid, test.gameRes)
		if err != nil {
			t.Errorf("output is not same want %q", err)
			fmt.Println(output)
		}
	}

}
