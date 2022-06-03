package services

import (
	types "dartscoreboard/models/types"
	"testing"
)

type nextTurnTest struct {
	playersIds       []int
	numOfRowsPerGame int
	activeRes        types.ActiveStatus
	round            int
	playerid         int
	throw            int
}

type StatusComTest struct {
	status    string
	activeRes types.ActiveStatus
	round     int
	playerid  int
	Throw     int
}

var NextTurnTest = []nextTurnTest{
	nextTurnTest{[]int{1, 2, 3}, 1, types.ActiveStatus{1, 1, 1, 1}, 1, 1, 2},
	nextTurnTest{[]int{1, 2, 3}, 3, types.ActiveStatus{1, 1, 1, 3}, 1, 2, 1},
	nextTurnTest{[]int{1, 2, 3, 4}, 12, types.ActiveStatus{1, 1, 4, 3}, 2, 1, 1},
}

var StatusTest = []StatusComTest{
	StatusComTest{"Completed", types.ActiveStatus{1, 1, 1, 1}, 0, 0, 0},
	StatusComTest{"Completed", types.ActiveStatus{1, 1, 1, 3}, 0, 0, 0},
}

func TestNextTrun(t *testing.T) {
	for _, test := range NextTurnTest {
		rid, pid, th := NextTurn(test.playersIds, test.numOfRowsPerGame, test.activeRes)
		if rid != test.round || pid != test.playerid || th != test.throw {
			t.Errorf("output is not same want %d,%d,%d", test.round, test.playerid, test.throw)
		}
	}

}

func TestStatusCompleted(t *testing.T) {
	for _, test := range StatusTest {
		rid, pid, th := StatusCompleted(test.status, test.activeRes)
		if rid != test.round || pid != test.playerid || th != test.Throw {
			t.Errorf("output is not same want %d,%d,%d", test.round, test.playerid, test.Throw)
		}
	}

}
