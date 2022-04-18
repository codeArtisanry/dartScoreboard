package models

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Delete Games From Games Table By Game Id
func DeleteGames(db *sql.DB, id int) error {
	fmt.Println(id)
	deleteGamePlayers, err := db.Prepare("DELETE FROM game_players WHERE game_id = ?;")
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = deleteGamePlayers.Exec(id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	delete, err := db.Prepare("DELETE FROM games WHERE id = ?;")
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = delete.Exec(id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
