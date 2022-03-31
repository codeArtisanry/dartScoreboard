package models

import "time"

type User struct {
	Id        int        `json:"id"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Email     string     `json:"email"`
	AvatarURL string     `json:"avatarUrl"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type Game struct {
	Id            int        `json:"id"`
	Name          string     `json:"gameName"`
	Type          string     `json:"gameType"`
	CreaterUserId string     `json:"createrUserId"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
}

type Score struct {
	Id        int        `json:"id"`
	RoundId   int        `json:"roundId"`
	Dart      int        `json:"dart"`
	Score     int        `json:"score"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type Round struct {
	Id           int        `json:"id"`
	GamePlayesId int        `json:"gamePlayerId"`
	Round        int        `json:"round"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
}

type GamePlayes struct {
	Id        int        `json:"id"`
	GameId    int        `json:"gameId"`
	UserId    int        `json:"userId"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
