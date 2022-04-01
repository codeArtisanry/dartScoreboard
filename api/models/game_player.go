package models

import "time"

type GamePlayer struct {
	Id        int        `json:"id"`
	GameId    int        `json:"gameId"`
	UserId    int        `json:"userId"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
