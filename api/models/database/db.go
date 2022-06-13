package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	migrate "github.com/rubenv/sql-migrate"
)

//Database Connection
func Database(database string) *sql.DB {
	// Database connection
	db, err := sql.Open("sqlite3", database)
	if err != nil {
		log.Fatal("Database Not Connected Due To: ", err)
	}

	return db
}

//Migration
func Migration(database string) {
	db := Database(database)
	fmt.Println("Waiting For Migrations...")
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}
	// Apply Migration
	n, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up)
	if err != nil {
		log.Fatal("Migration Not Apply Due To: ", err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
}
