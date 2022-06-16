package types

import "time"

// Game defines the structure for an API response of GameResponse

// Game game
// swagger:model game
type Game struct {
	Id               int        `json:"id"`
	Name             string     `json:"name"`
	Type             string     `json:"type,omitempty"`
	Status           string     `json:"status"`
	PlayersIds       []int      `json:"players"`
	CreaterUserEmail string     `json:"creater_user_email"`
	CreatedAt        *time.Time `json:"created_at,omitempty"`
	UpdatedAt        *time.Time `json:"updated_at,omitempty"`
}

// swagger:response GameResponse
type GameResponse struct {
	// in: body
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	Status        string `json:"status"`
	CreaterUserId int    `json:"creater_user_id,omitempty"`
	CreaterName   string `json:"creater_name,omitempty"`
	// in: body
	Players []GamePlayerResponse `json:"players,omitempty"`
}

// swagger:response GamesPaginationResponse
type GamesPaginationResponse struct {
	// in: body
	GameResponses []GameResponse `json:"list"`
	PrePageLink   string         `json:"previous,omitempty"`
	PostPageLink  string         `json:"next,omitempty"`
}

// swagger:parameters getGame editGame deleteGame activeStatus addScore
type _ struct {

	// in: path
	Id int `json:"id"`
}
