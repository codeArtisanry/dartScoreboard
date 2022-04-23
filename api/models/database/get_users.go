package models

import (
	"dartscoreboard/models/types"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Get All User From Users Table
func GetUsers(db *sql.DB, offset int, searchFisrtName string, searchLastName string, user types.User) ([]types.User, error) {
	var users []types.User
	query := fmt.Sprintf("SELECT id, first_name, last_name, email FROM users WHERE first_name LIKE '%s' AND last_name LIKE '%s' ORDER BY first_name ASC LIMIT 5 OFFSET %d;", searchFisrtName, searchLastName, offset)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email)
		if err != nil {
			return users, err
		}
		userJson := types.User{
			Id:        user.Id,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email}
		users = append(users, userJson)
	}
	return users, nil
}
