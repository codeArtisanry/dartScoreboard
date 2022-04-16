package models

import (
	types "dartscoreboard/models/types"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Insert User Data to users Table and return error
func InsertUserDetails(db *sql.DB, user types.User) error {
	insert, err := db.Prepare("INSERT INTO users (first_name, last_name, email, avatar_url) VALUES (?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = insert.Exec(user.FirstName, user.LastName, user.Email, user.AvatarURL)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Sucessfully Data Inserted on users Table")
	return nil
}

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
