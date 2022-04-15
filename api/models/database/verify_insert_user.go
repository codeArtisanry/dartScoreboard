package models

import (
	types "dartscoreboard/models/types"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Verify User Exists or Not? Then Insert User Data to Users Table and return userid
func VerifyAndInsertUser(db *sql.DB, user types.User) (int, error) {
	user, err := SelectUserInfoByEmail(db, user.Email, user)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("no row", err)
			err := InsertUserDetails(db, user)
			if err != nil {
				fmt.Println("InsertUserDetails", err)
			}
			user, err := SelectUserInfoByEmail(db, user.Email, user)
			if err != nil {
				fmt.Println("SelectUserInfoByEmail", err)
				return user.Id, err
			}
			return user.Id, nil
		} else {
			fmt.Println("without no row", err)
			return user.Id, err
		}
	}
	return user.Id, nil
}
