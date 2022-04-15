package models

import (
	"dartscoreboard/models/types"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Select user id using user email and return user id and error
func SelectUserInfoByEmail(db *sql.DB, email string, user types.User) (types.User, error) {
	readUserIdQuery := fmt.Sprintf("SELECT id, first_name, last_name, email FROM users WHERE email = '%s'", email)
	row := db.QueryRow(readUserIdQuery)
	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		fmt.Println(err)
		return user, err
	}
	return user, nil
}

// Select user infomation by user id
func SelectUserInfoById(db *sql.DB, id int, user types.User) (types.User, error) {
	readUserIdQuery := fmt.Sprintf("SELECT first_name,last_name FROM users WHERE id = %d", id)
	row := db.QueryRow(readUserIdQuery)
	err := row.Scan(&user.FirstName, &user.LastName)
	if err != nil {
		return user, err
	}
	return user, nil
}
