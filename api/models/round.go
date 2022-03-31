package models

import "time"

type Round struct {
	Id           int        `json:"id"`
	GamePlayesId int        `json:"gamePlayerId"`
	Round        int        `json:"round"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
}
