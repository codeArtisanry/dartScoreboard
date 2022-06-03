package services

import "dartscoreboard/models/types"

func NextTurn(playersIds []int, numOfRowsPerGame int, activeRes types.ActiveStatus) (round int, playerid int, throw int) {
	mapPlayersIds := make(map[int]int)
	for playerIdIndex := 0; playerIdIndex < len(playersIds); playerIdIndex++ {
		mapPlayersIds[playersIds[playerIdIndex]] = playerIdIndex
	}
	if numOfRowsPerGame%(3*len(playersIds)) == 0 {
		activeRes.Round = activeRes.Round + 1
	}
	if activeRes.Throw%3 == 0 {
		activeRes.Throw = 1
		test := mapPlayersIds[activeRes.PlayerId]
		if test == (len(playersIds) - 1) {
			activeRes.PlayerId = playersIds[0]
		} else {
			activeRes.PlayerId = playersIds[test+1]
		}
	} else {
		activeRes.Throw = activeRes.Throw + 1
	}
	return activeRes.Round, activeRes.PlayerId, activeRes.Throw
}

func StatusCompleted(status string, activeRes types.ActiveStatus) (round int, playerid int, Throw int) {
	if status == "Completed" {
		activeRes.Round = 0
		activeRes.PlayerId = 0
		activeRes.Throw = 0
	}
	return activeRes.Round, activeRes.PlayerId, activeRes.Throw
}
