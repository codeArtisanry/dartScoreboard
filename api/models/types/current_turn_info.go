package types

// swagger:response CurrentPlayerInfo
type CurrentPlayerInfo struct {
	Id    int          `json:"id"`
	Name  string       `json:"name"`
	Email string       `json:"email"`
	Round int          `json:"round,omitempty"`
	Throw int          `json:"throw,omitempty"`
	Score int          `json:"score"`
	// in: body
	Game  GameResponse `json:"game"`
}

// swagger:response Scoreboard
type Scoreboard struct {
	// in: body
	PlayersScore []PlayerScore `json:"players"`
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
