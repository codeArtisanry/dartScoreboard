// package models

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"os"

// 	_ "github.com/mattn/go-sqlite3"
// )

// var DB *sql.DB

// func ConnectDatabase() {
// 	os.ExpandEnv("")
// 	db, err := sql.Open("sqlite3", os.ExpandEnv("./${DB_NAME}"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("successfully connected database")
// 	create, _ := db.Prepare("CREATE TABLE IF NOT EXISTS dart (token INTEGER PRIMARY KEY, userEmail TEXT,userName TEXT)")
// 	create.Exec()
// 	insert, _ := db.Prepare("INSERT INTO dart (token, userEmail, userName) VALUES (?, ?, ?)")
// 	insert.Exec(738473847, "jeel@improwied.com", "Jeel")
// }
package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	migrate "github.com/rubenv/sql-migrate"
)

func Database() {
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
}