package models

import (
	"dartscoreboard/models/types"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func InserGameQuery(db *sql.DB, id int, game types.Game) (result sql.Result) {
	insert, err := db.Prepare("INSERT INTO games(name, type, status, creater_user_id) VALUES(?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
	}
	result, err = insert.Exec(game.Name, game.Type, "Not Started", id)
	if err != nil {
		fmt.Println(err)
	}
	return result
}
