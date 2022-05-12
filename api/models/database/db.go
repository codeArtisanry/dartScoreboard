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
	fmt.Println("Waiting For Migrations...")
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
