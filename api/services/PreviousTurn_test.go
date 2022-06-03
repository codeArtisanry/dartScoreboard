package services_test

import (
	models "dartscoreboard/models/database"
	"dartscoreboard/services"
	"testing"
)

type PreviousTurnTest struct {
	round      int
	playerId   int
	turn       int
	PlayersIds []int
	expected   models.PreviousTurn
}

var previousTurnTest = []PreviousTurnTest{
	PreviousTurnTest{1, 1, 1, []int{1, 2, 3, 4}, models.PreviousTurn{0, 4, 3}},
	PreviousTurnTest{3, 4, 3, []int{1, 2, 3, 4}, models.PreviousTurn{3, 4, 2}},
	PreviousTurnTest{2, 3, 2, []int{1, 2, 3, 4}, models.PreviousTurn{2, 3, 1}},
}

func TestFindPreviousTurn(t *testing.T) {
	for _, test := range previousTurnTest {
		if output := services.FindPreviousTurn(test.round, test.playerId, test.turn, test.PlayersIds); output != test.expected {
			t.Errorf("Output %d not equal to expected %d", output, test.expected)
		}
	}
}
