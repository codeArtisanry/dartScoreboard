package models

type User struct {
	Id      string `json:"id"`
	Email   string `json:"email"`
	Picture string `json:"picture"`
}

type Game struct {
	GameName        string
	GameType        string
	PlayersNames    []string
	GameTargetScore string
}
