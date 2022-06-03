package services

import models "dartscoreboard/models/database"

// Find Previous turn, playerId and round
func FindPreviousTurn(round int, playerId int, turn int, playersId []int) models.PreviousTurn {
	var currentPlayerIndex int
	mapPlayerid := make(map[int]int)
	for i := 0; i <= len(playersId)-1; i++ {
		mapPlayerid[playersId[i]] = i
	}
	turn = turn - 1
	if turn == 0 {
		turn = 3
		currentPlayerIndex = mapPlayerid[playerId]
		currentPlayerIndex = currentPlayerIndex - 1
		if currentPlayerIndex == -1 {
			currentPlayerIndex = len(playersId) - 1
			playerId = playersId[currentPlayerIndex]
			round = round - 1
		} else {
			playerId = playersId[currentPlayerIndex]
		}
	}
	previousTurn := models.PreviousTurn{
		Round:    round,
		PlayerId: playerId,
		Turn:     turn,
	}
	return previousTurn
}
