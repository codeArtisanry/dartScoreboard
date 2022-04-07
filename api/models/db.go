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
func VerifyAndInsertUser(db *sql.DB, user User) (int, error) {
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
				return userId, err
			}
			return userId, nil
		} else {
			fmt.Println(err)
			return userId, err
		}
	}
	return userId, nil
}

// Get All User From Users Table
func GetUsers(db *sql.DB, user User) ([]User, error) {
	var users []User
	rows, err := db.Query("SELECT id, first_name, last_name, email FROM users")
	if err != nil {
		fmt.Println(err)
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email)
		if err != nil {
			return users, err
		}
		userJson := User{
			Id:        user.Id,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email}
		users = append(users, userJson)
	}
	return users, nil
}

// Insert Games in to Games Table
func InsertGames(db *sql.DB, game Game) (Game, error) {
	insert, err := db.Prepare("INSERT INTO games(name, type, creater_user_id) VALUES(?, ?, ?)")
	if err != nil {
		fmt.Println(err)
	}
	result, err := insert.Exec(game.Name, game.Type, game.CreaterUserId)
	if err != nil {
		fmt.Println(err)
	}
	gameId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}
	gameJson := Game{
		Id:            int(gameId),
		Name:          game.Name,
		Type:          game.Type,
		CreaterUserId: game.CreaterUserId,
	}
	return gameJson, err
}

// Delete Games From Games Table By Game Id
func DeleteGames(db *sql.DB, id int) error {
	delete, err := db.Prepare("DELETE FROM games WHERE id= ?;")
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = delete.Exec(id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// Update Games From Games Table By Game Id
func UpdateGame(db *sql.DB, id int, game Game) (Game, error) {
	gameJson := Game{
		Id:            id,
		Name:          game.Name,
		Type:          game.Type,
		CreaterUserId: game.CreaterUserId,
	}
	query, err := db.Prepare("UPDATE games SET name = ?, type = ? WHERE id = ?;")
	if err != nil {
		fmt.Println(err)
	}
	_, err = query.Exec(game.Name, game.Type, id)
	if err != nil {
		fmt.Println(err)
		return gameJson, err
	}
	return gameJson, nil
}

// Get Game From Games Table By Game Id
func GetGame(db *sql.DB, id int, game Game) (Game, error) {
	query := fmt.Sprintf("SELECT id, name, type FROM games WHERE id = %d;", id)
	row := db.QueryRow(query)
	err := row.Scan(&game.Id, &game.Name, &game.Type)
	if err != nil {
		return game, err
	}
	return game, nil
}

//Get All Games From Game Table
func GetGames(db *sql.DB, game Game) ([]Game, error) {
	var games []Game
	rows, err := db.Query("SELECT id, name, type, creater_user_id FROM games")
	if err != nil {
		fmt.Println(err)
		return games, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&game.Id, &game.Name, &game.Type, &game.CreaterUserId)
		if err != nil {
			return games, err
		}
		gameJson := Game{
			Id:            game.Id,
			Name:          game.Name,
			Type:          game.Type,
			CreaterUserId: game.CreaterUserId}
		games = append(games, gameJson)
	}
	return games, nil
}
