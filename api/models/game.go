package models

import "time"

type Game struct {
	Id            int        `json:"id"`
	Name          string     `json:"gameName"`
	Type          string     `json:"gameType"`
	Status        string     `json:"gameStatus"`
	CreaterUserId int        `json:"createrUserId"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
}
