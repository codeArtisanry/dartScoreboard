package services_test

import (
	"dartscoreboard/services"
	"testing"
)

type RemainingScoreTest struct {
	gameType       string
	totalScore     int
	remainingScore int
}

var remainingScoreTest = []RemainingScoreTest{
	RemainingScoreTest{"Target Score-101", 33, 68},
	RemainingScoreTest{"Target Score-301", 33, 268},
	RemainingScoreTest{"High Score", 33, 33},
}

func TestRemaingScore(t *testing.T) {
	for _, test := range remainingScoreTest {
		r := services.RemainScore(test.gameType, test.totalScore)
		if r != test.remainingScore {
			t.Errorf("output is got %d and want %d", test.remainingScore, r)
		}
	}

}
