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

