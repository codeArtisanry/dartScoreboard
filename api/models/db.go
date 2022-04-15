package models

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

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
func SelectUserInfoByEmail(db *sql.DB, email string, user User) (User, error) {
	readUserIdQuery := fmt.Sprintf("SELECT id, first_name, last_name, email FROM users WHERE email = '%s'", email)
	row := db.QueryRow(readUserIdQuery)
	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		fmt.Println(err)
		return user, err
	}
	return user, nil
}

// Select user infomation by user id
func SelectUserInfoById(db *sql.DB, id int, user User) (User, error) {
	readUserIdQuery := fmt.Sprintf("SELECT first_name,last_name FROM users WHERE id = %d", id)
	row := db.QueryRow(readUserIdQuery)
	err := row.Scan(&user.FirstName, &user.LastName)
	if err != nil {
		return user, err
	}
	return user, nil
}

// Verify User Exists or Not? Then Insert User Data to Users Table and return userid
func VerifyAndInsertUser(db *sql.DB, user User) (int, error) {
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

// Get All User From Users Table
func GetUsers(db *sql.DB, offset int, user User) ([]User, error) {
	var users []User
	query := fmt.Sprintf("SELECT id, first_name, last_name, email FROM users ORDER BY first_name ASC LIMIT 5 OFFSET %d;", offset)
	rows, err := db.Query(query)
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

// Get all perticuler game of players infomation list by game id
func GetPlayersInfoByGameId(db *sql.DB, id int, gamePlayerRes GamePlayerResponse) ([]GamePlayerResponse, error) {
	var PlayersInfo []GamePlayerResponse
	query := fmt.Sprintf("SELECT DISTINCT users.id,users.first_name,users.last_name,users.email from users LEFT JOIN game_players ON game_players.user_id=users.id where game_players.game_id = %d;", id)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return PlayersInfo, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&gamePlayerRes.Id, &gamePlayerRes.FirstName, &gamePlayerRes.LastName, &gamePlayerRes.Email)
		if err != nil {
			fmt.Println(err)
			return PlayersInfo, err
		}
		gamePlayerInfo := GamePlayerResponse{
			Id:        gamePlayerRes.Id,
			FirstName: gamePlayerRes.FirstName,
			LastName:  gamePlayerRes.LastName,
			Email:     gamePlayerRes.Email}
		PlayersInfo = append(PlayersInfo, gamePlayerInfo)
	}
	return PlayersInfo, nil
}

// Insert periculer game releted all players id in to game players table
func InsertPlayers(db *sql.DB, game Game, gameId int, playerRes GamePlayerResponse) ([]GamePlayerResponse, error) {
	var playersInfo []GamePlayerResponse
	gamePlayerIds := game.PlayerIds
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(gamePlayerIds), func(i, j int) { gamePlayerIds[i], gamePlayerIds[j] = gamePlayerIds[j], gamePlayerIds[i] })
	for i := 0; i <= len(gamePlayerIds)-1; i++ {
		insert, err := db.Prepare("INSERT INTO game_players(user_id, game_id) VALUES(?, ?)")
		if err != nil {
			fmt.Println(err)
			return playersInfo, err
		}
		fmt.Println(gamePlayerIds[i], gameId)
		_, err = insert.Exec(gamePlayerIds[i], gameId)
		if err != nil {
			fmt.Println(err)
			return playersInfo, err
		}
	}
	playersInfo, err := GetPlayersInfoByGameId(db, gameId, playerRes)
	if err != nil {
		fmt.Println(err)
		return playersInfo, err
	}
	fmt.Println(playersInfo)
	return playersInfo, nil
}

// Insert Games in to Games Table
func InsertGames(db *sql.DB, user User, game Game, gamePlayer GamePlayer, gamePlayerRes GamePlayerResponse) (GameResponse, error) {
	var gameResJson GameResponse
	createrInfo, err := SelectUserInfoByEmail(db, game.CreaterUserEmail, user)
	if err != nil {
		fmt.Println(err)
		return gameResJson, err
	}
	insert, err := db.Prepare("INSERT INTO games(name, type, status, creater_user_id) VALUES(?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
		return gameResJson, err
	}
	result, err := insert.Exec(game.Name, game.Type, game.Status, createrInfo.Id)
	if err != nil {
		fmt.Println(err)
		return gameResJson, err
	}
	gameId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return gameResJson, err
	}

	id := int(gameId)
	PlayersInfo, err := InsertPlayers(db, game, id, gamePlayerRes)
	if err != nil {
		fmt.Println(err)
		return gameResJson, err
	}
	gameResJson = GameResponse{
		Id:               id,
		Name:             game.Name,
		Type:             game.Type,
		Status:           game.Status,
		CreaterFirstName: createrInfo.FirstName,
		CreaterLastName:  createrInfo.LastName,
		PlayersInfo:      PlayersInfo,
	}
	return gameResJson, nil
}

// Delete Games From Games Table By Game Id
func DeleteGames(db *sql.DB, id int) error {
	fmt.Println(id)
	deleteGamePlayers, err := db.Prepare("DELETE * FROM game_players WHERE game_id = ?;")
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = deleteGamePlayers.Exec(id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	delete, err := db.Prepare("DELETE FROM games WHERE id = ?;")
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
func UpdateGame(db *sql.DB, id int, user User, game Game, playerRes GamePlayerResponse) (GameResponse, error) {
	var gameResJson GameResponse
	createrInfo, err := SelectUserInfoByEmail(db, game.CreaterUserEmail, user)
	if err != nil {
		fmt.Println(err)
		return gameResJson, err
	}
	query, err := db.Prepare("UPDATE games SET name = ?, type = ?, status = ?, creater_user_id = ? WHERE id = ?;")
	if err != nil {
		fmt.Println(err)
		return gameResJson, err
	}
	_, err = query.Exec(game.Name, game.Type, game.Status, createrInfo.Id, id)
	if err != nil {
		fmt.Println(err)
		return gameResJson, err
	}
	deleteGamePlayers, err := db.Prepare("DELETE * FROM game_players WHERE game_id = ?;")
	if err != nil {
		fmt.Println(err)
		return gameResJson, err
	}
	_, err = deleteGamePlayers.Exec(id)
	if err != nil {
		fmt.Println(err)
		return gameResJson, err
	}

	playersInfo, err := InsertPlayers(db, game, id, playerRes)
	if err != nil {
		fmt.Println(err)
		return gameResJson, err
	}
	gameResJson = GameResponse{
		Id:               id,
		Name:             game.Name,
		Type:             game.Type,
		Status:           game.Status,
		CreaterFirstName: createrInfo.FirstName,
		CreaterLastName:  createrInfo.LastName,
		PlayersInfo:      playersInfo,
	}
	return gameResJson, nil
}

// Get Game From Games Table By Game Id
func GetGame(db *sql.DB, id int, gameRes GameResponse, user User, gamePlayerRes GamePlayerResponse) (GameResponse, error) {
	var gameResJson GameResponse
	query := fmt.Sprintf("SELECT id, name, type, status, creater_user_id FROM games WHERE id = %d;", id)
	row := db.QueryRow(query)
	err := row.Scan(&gameRes.Id, &gameRes.Name, &gameRes.Type, &gameRes.Status, &gameRes.CreaterUserId)
	if err != nil {
		return gameResJson, err
	}
	createrInfo, err := SelectUserInfoById(db, gameRes.CreaterUserId, user)
	if err != nil {
		return gameResJson, err
	}
	PlayersInfo, err := GetPlayersInfoByGameId(db, id, gamePlayerRes)
	if err != nil {
		fmt.Println(err)
		return gameResJson, err
	}
	gameResJson = GameResponse{
		Id:               gameRes.Id,
		Name:             gameRes.Name,
		Type:             gameRes.Type,
		Status:           gameRes.Status,
		CreaterFirstName: createrInfo.FirstName,
		CreaterLastName:  createrInfo.LastName,
		PlayersInfo:      PlayersInfo,
	}
	return gameResJson, nil
}

//Get All Games From Game Table
func GetGames(db *sql.DB, page int, offset int, gameRes GameResponse, user User, gamePlayerRes GamePlayerResponse) ([]GameResponse, error) {
	var gamesResJson []GameResponse
	id := 1
	query := fmt.Sprintf("SELECT games.id, games.name, games.type, games.status, games.creater_user_id FROM games LEFT JOIN game_players ON games.id = game_players.game_id WHERE game_players.user_id = %d ORDER BY games.created_at DESC LIMIT 10 OFFSET %d;", id, offset)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return gamesResJson, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&gameRes.Id, &gameRes.Name, &gameRes.Type, &gameRes.Status, &gameRes.CreaterUserId)
		if err != nil {
			return gamesResJson, err
		}
		createrInfo, err := SelectUserInfoById(db, gameRes.CreaterUserId, user)
		if err != nil {
			return gamesResJson, err
		}
		playersInfo, err := GetPlayersInfoByGameId(db, gameRes.Id, gamePlayerRes)
		if err != nil {
			fmt.Println(err)
			return gamesResJson, err
		}
		gameResJson := GameResponse{
			Id:               gameRes.Id,
			Name:             gameRes.Name,
			Type:             gameRes.Type,
			Status:           gameRes.Status,
			CreaterFirstName: createrInfo.FirstName,
			CreaterLastName:  createrInfo.LastName,
			PlayersInfo:      playersInfo}
		gamesResJson = append(gamesResJson, gameResJson)
	}
	return gamesResJson, nil
}

// Find creater id by game id
func FindCreaterIdByGameId(db *sql.DB, id int, gameRes GameResponse) (int, error) {
	query := fmt.Sprintf("SELECT creater_user_id FROM games WHERE id = %d;", id)
	row := db.QueryRow(query)
	err := row.Scan(&gameRes.CreaterUserId)
	if err != nil {
		return gameRes.CreaterUserId, err
	}
	return gameRes.CreaterUserId, nil
}
