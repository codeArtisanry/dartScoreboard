package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDatabase() {
	os.ExpandEnv("")
	db, err := sql.Open("sqlite3", os.ExpandEnv("./${DB_NAME}"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("successfully connected database")
	create, _ := db.Prepare("CREATE TABLE IF NOT EXISTS dart (token INTEGER PRIMARY KEY, userEmail TEXT,userName TEXT)")
	create.Exec()
	insert, _ := db.Prepare("INSERT INTO dart (token, userEmail, userName) VALUES (?, ?, ?)")
	insert.Exec(738473847, "jeel@improwied.com", "Jeel")
}
