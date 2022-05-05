package types

// swagger:response CurrentPlayerInfo
type CurrentPlayerInfo struct {
	Id               int               `json:"id"`
	Name             string            `json:"game_name"`
	Type             string            `json:"game_type"`
	Status           string            `json:"game_status"`
	Round            int               `json:"round,omitempty"`
	Throw            int               `json:"throw,omitempty"`
	// in: body
	ActivePlayerInfo *ActivePlayerInfo `json:"active_player_info"`
}

// swagger:response ActivePlayerInfo
type ActivePlayerInfo struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Score     int    `json:"score"`
}

// swagger:response Scoreboard
type Scoreboard struct {
	// in: body
	PlayersScore []PlayerScore `json:"players_score"`
	Winner       string        `json:"winner,omitempty"`
}

// swagger:response PlayerScore
type PlayerScore struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	// in: body
	Rounds []Rounds `json:"rounds"`
	Total  int      `json:"total"`
}

// swagger:response Rounds
type Rounds struct {
	Round       int    `json:"round"`
	ThrowsScore []int  `json:"throws_score"`
	CheckRound  string `json:"check_round"`
	RoundTotal  int    `json:"round_total"`
}
