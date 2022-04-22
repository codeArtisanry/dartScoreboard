package types

// StartGame startgame
// swagger:response CurrentTurnInfo
type CurrentTurnInfo struct {
	Id               int               `json:"id"`
	Name             string            `json:"game_name"`
	Type             string            `json:"game_type"`
	Round            int               `json:"round"`
	Throw            int               `json:"throw"`
	ActivePlayerInfo *ActivePlayerInfo `json:"active_player_info"`
	Scoreboard       Scoreboard        `json:"scoreboard"`
}

// swagger:response CurrentTurnInfo
type ActivePlayerInfo struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// swagger:response CurrentTurnInfo
type Scoreboard struct {
	PlayersScore []PlayerScore `json:"players_score"`
	Winner       string        `json:"winner,omitempty"`
}

// swagger:response CurrentTurnInfo
type PlayerScore struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Round     []Rounds `json:"round "`
	Total     int      `json:"total"`
}

// swagger:response CurrentTurnInfo
type Rounds struct {
	Round       int   `json:"round"`
	ThrowsScore []int `json:"throws_score"`
	RoundTotal  int   `json:"round_total"`
}
