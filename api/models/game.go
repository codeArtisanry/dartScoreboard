package models

import "time"

// Game defines the structure for an API response of GameResponse

// swagger:model Game
type Game struct {
	Id               int        `json:"id"`
	Name             string     `json:"game_name"`
	Type             string     `json:"game_type"`
	Status           string     `json:"game_status"`
	PlayerIds        []int      `json:"player_ids"`
	CreaterUserEmail string     `json:"creater_user_email"`
	CreatedAt        *time.Time `json:"created_at,omitempty"`
	UpdatedAt        *time.Time `json:"updated_at,omitempty"`
}

// swagger:response GameResponse
type GameResponse struct {
	Id               int                  `json:"id"`
	Name             string               `json:"game_name"`
	Type             string               `json:"game_type"`
	Status           string               `json:"game_status"`
	CreaterUserId    int                  `json:"creater_user_id,omitempty"`
	CreaterFirstName string               `json:"creater_first_name"`
	CreaterLastName  string               `json:"creater_last_name"`
	PlayersInfo      []GamePlayerResponse `json:"players_info"`
}
