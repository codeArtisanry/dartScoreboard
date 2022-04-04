package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	migrate "github.com/rubenv/sql-migrate"
)

//Database Connection and Migration
func Database() *sql.DB {
	migrations := &migrate.FileMigrationSource{
		Dir: "models/migrations",
	}
	// Database connection
	db, err := sql.Open("sqlite3", "dart.db")
	if err != nil {
		log.Fatal("Database Not Connected Due To: ", err)
	}
	// Apply Migration
	n, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up)
	if err != nil {
		log.Fatal("Migration Not Applay Due To: ", err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
	return db
}

// Insert User Data to users Table and return error
func InsertUserDetails(db *sql.DB, user User) error {
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

// Select user id using user email and return user id and error
func SelectUserIdByEmail(db *sql.DB, user User) (int, error) {
	email := user.Email
	readUserIdQuery := fmt.Sprintf("SELECT id FROM users WHERE email = '%s'", email)
	row := db.QueryRow(readUserIdQuery)
	err := row.Scan(&user.Id)
	if err != nil {
		return user.Id, err
	}
	return user.Id, nil
}

//Verify User Exists or Not? Then Insert User Data to Users Table and return userid
func VerifyAndInsertUser(db *sql.DB, user User) int {
	userId, err := SelectUserIdByEmail(db, user)
	if err != nil {
		if err == sql.ErrNoRows {
			err := InsertUserDetails(db, user)
			if err != nil {
				fmt.Println(err)
			}
			userId, err := SelectUserIdByEmail(db, user)
			if err != nil {
				fmt.Println(err)
			}
			return userId
		} else {
			fmt.Println(err)
		}
	}
	return userId
}
