package models

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func (database DataBase) UsersQuery(offset string, searchFisrtName string, searchLastName string) (rows *sql.Rows) {
	db := database.Db
	query := fmt.Sprintf("SELECT id, first_name, last_name, email FROM users WHERE first_name LIKE '%s' AND last_name LIKE '%s' ORDER BY first_name %s;", searchFisrtName, searchLastName, offset)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	return rows
}
