package models

import (
	"dartscoreboard/models/types"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Find creater id by game id
func FindCreaterIdByGameId(db *sql.DB, id int, gameRes types.GameResponse) (int, error) {
	query := fmt.Sprintf("SELECT creater_user_id FROM games WHERE id = %d;", id)
	row := db.QueryRow(query)
	err := row.Scan(&gameRes.CreaterUserId)
	if err != nil {
		return gameRes.CreaterUserId, err
	}
	return gameRes.CreaterUserId, nil
}
