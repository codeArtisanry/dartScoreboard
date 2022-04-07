package models

import "time"

type Game struct {
	Id               int        `json:"id"`
	Name             string     `json:"game_name"`
	Type             string     `json:"game_type"`
	CreaterUserEmail string     `json:"creater_user_email,omitempty"`
	CreaterUserId    int        `json:"creater_user_id,omitempty"`
	CreatedAt        *time.Time `json:"created_at,omitempty"`
	UpdatedAt        *time.Time `json:"updated_at,omitempty"`
}
