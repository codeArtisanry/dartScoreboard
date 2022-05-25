package models_test

import (
	models "dartscoreboard/models/database"
	"testing"
)

type addTest struct {
	arg1     int
	arg2     int
	arg3     int
	arg4     []int
	expected models.PreviousTurn
}

var addTests = []addTest{
	addTest{1, 1, 1, []int{1, 2, 3, 4}, models.PreviousTurn{0, 4, 3}},
	addTest{3, 4, 3, []int{1, 2, 3, 4}, models.PreviousTurn{3, 4, 2}},
	addTest{2, 3, 2, []int{1, 2, 3, 4}, models.PreviousTurn{2, 3, 1}},
}

func TestFindPreviousTurn(t *testing.T) {
	for _, test := range addTests {
		if output := models.FindPreviousTurn(test.arg1, test.arg2, test.arg3, test.arg4); output != test.expected {
			t.Errorf("Output %d not equal to expected %d", output, test.expected)
		}
	}
}
