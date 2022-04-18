package models

import (
	"dartscoreboard/models/types"
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Insert periculer game releted all players id in to game players table
func InsertPlayers(db *sql.DB, game types.Game, gameId int, playerRes types.GamePlayerResponse) ([]types.GamePlayerResponse, error) {
	var playersInfo []types.GamePlayerResponse
	gamePlayerIds := game.PlayersIds
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(gamePlayerIds), func(i, j int) { gamePlayerIds[i], gamePlayerIds[j] = gamePlayerIds[j], gamePlayerIds[i] })
	for i := 0; i <= len(gamePlayerIds)-1; i++ {
		insert, err := db.Prepare("INSERT INTO game_players(user_id, game_id) VALUES(?, ?)")
		if err != nil {
			fmt.Println(err)
			return playersInfo, err
		}
		_, err = insert.Exec(gamePlayerIds[i], gameId)
		if err != nil {
			fmt.Println(err)
			return playersInfo, err
		}
	}
	playersInfo, err := GetPlayersInfoByGameId(db, gameId, playerRes)
	if err != nil {
		fmt.Println(err)
		return playersInfo, err
	}
	return playersInfo, nil
}
