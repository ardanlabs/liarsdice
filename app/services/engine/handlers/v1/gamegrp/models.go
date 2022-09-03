package gamegrp

// Status represents the game status.
type Status struct {
	Status        string   `json:"status"`
	AnteUSD       float64  `json:"ante_usd"`
	LastOutAcctID string   `json:"last_out"`
	LastWinAcctID string   `json:"last_win"`
	CurrentPlayer int      `json:"current_player"`
	CurrentCup    int      `json:"current_cup"`
	Round         int      `json:"round"`
	Cups          []Cup    `json:"cups"`
	CupsOrder     []string `json:"player_order"`
	Bets          []Bet    `json:"bets"`
	Balances      []string `json:"balances"`
}

// Cup represents the cup response.
type Cup struct {
	AccountID string `json:"account"`
	Dice      []int  `json:"dice"`
	Outs      int    `json:"outs"`
}

// Bet represents the bet response.
type Bet struct {
	AccountID string `json:"account"`
	Number    int    `json:"number"`
	Suite     int    `json:"suite"`
}
