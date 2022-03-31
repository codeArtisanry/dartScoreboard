package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	migrate "github.com/rubenv/sql-migrate"
)

func Database() *sql.DB {
	migrations := &migrate.FileMigrationSource{
		Dir: "models/migrations",
	}
	db, err := sql.Open("sqlite3", "dart.db")
	if err != nil {
		log.Fatal(err)
	}

	n, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
	return db
}

func InsertUserData(db *sql.DB, user User) {
	insert, err := db.Prepare("INSERT INTO users (id, first_name, last_name, email, avatar) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("from database", user.FirstName, user.LastName, user.Email, user.AvatarURL)
	_, err = insert.Exec(user.FirstName, user.LastName, user.Email, user.AvatarURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Sucessfully Data Inserted on users Table")
}
func InsertGameData(db *sql.DB, game Game) {
	insert, err := db.Prepare("INSERT INTO games (name, type, creater_email) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("from database", game.Name, game.Type, game.CreaterEmail)
	_, err = insert.Exec(game.Name, game.Type, game.CreaterEmail)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Sucessfully Data Inserted on games Table")
}
