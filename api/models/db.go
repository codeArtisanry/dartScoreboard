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

// Insert User Data to users Table
func InsertUserData(db *sql.DB, user User) {
	insert, err := db.Prepare("INSERT INTO users (first_name, last_name, email, avatar_url) VALUES (?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
	}
	_, err = insert.Exec(user.FirstName, user.LastName, user.Email, user.AvatarURL)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Sucessfully Data Inserted on users Table")
}

// Select user id using user email and return user id
func SelectUserByEmail(db *sql.DB, user User) int {
	email := user.Email
	readQuery := fmt.Sprintf("SELECT id FROM users WHERE email = '%s'", email)
	row := db.QueryRow(readQuery)
	row.Scan(&user.Id)
	return user.Id
}

//Verify User Exists or Not? Then Insert User Data to Users Table
func VerifyUser(user User) int {
	db := Database()
	id := SelectUserByEmail(db, user)
	if id == 0 {
		fmt.Println("user not found")
		InsertUserData(db, user)
		id := SelectUserByEmail(db, user)
		return id
	} else {
		fmt.Printf("user found and user id is : %d ", user.Id)
		return id
	}
}
