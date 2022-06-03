package services

import (
	"fmt"
	"strconv"
	"strings"
)

func RemainScore(gameType string, totalScore int) (remaingscore int) {
	GameType := strings.Split(gameType, "-")
	if GameType[0] == "Target Score" {
		TargetScore, err := strconv.Atoi(GameType[1])
		if err != nil {

			fmt.Println(err)
		}
		remaingscore = TargetScore - totalScore
	} else {
		remaingscore = totalScore
	}
	return remaingscore
}
