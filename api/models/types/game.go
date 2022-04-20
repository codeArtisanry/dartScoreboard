package types

import "time"

// Game defines the structure for an API response of GameResponse

// Game game
// swagger:model game
type Game struct {
	Id               int        `json:"id"`
	Name             string     `json:"game_name"`
	Type             string     `json:"game_type,omitempty"`
	Status           string     `json:"game_status"`
	PlayersIds       []int      `json:"players_ids"`
	CreaterUserEmail string     `json:"creater_user_email"`
	CreatedAt        *time.Time `json:"created_at,omitempty"`
	UpdatedAt        *time.Time `json:"updated_at,omitempty"`
}

// swagger:response GameResponse
type GameResponse struct {
	// in: body
	Id            int    `json:"id"`
	Name          string `json:"game_name"`
	Type          string `json:"game_type"`
	Status        string `json:"game_status"`
	CreaterUserId int    `json:"creater_user_id,omitempty"`
	CreaterName   string `json:"creater_name"`
	// in: body
	Players      []GamePlayerResponse `json:"players"`
	PreviousPage string               `json:"previous_page,omitempty"`
	NextPage     string               `json:"next_page,omitempty"`
}

// swagger:response GamesPaginationResponse
type GamesPaginationResponse struct {
	// in: body
	GameResponses []GameResponse `json:"game_responses"`
	PrePageLink   string         `json:"pre_page_link"`
	PostPageLink  string         `json:"post_page_link"`
}

// swagger:parameters getGame editGame deleteGame addScore
type _ struct {

	// in: path
	Id int `json:"id"`
}
