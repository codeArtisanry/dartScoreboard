package models

import "time"

// swagger:response PlayerResponce
type GamePlayerResponce struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// swagger:model GamePlayer
type GamePlayer struct {
	Id        int        `json:"id"`
	GameId    int        `json:"game_id"`
	UserId    int        `json:"user_id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
