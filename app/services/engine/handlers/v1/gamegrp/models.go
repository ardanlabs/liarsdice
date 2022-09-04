package gamegrp

// Status represents the game status.
type Status struct {
	Status        string   `json:"status"`
	AnteUSD       float64  `json:"ante_usd"`
	LastOutAcctID string   `json:"last_out"`
	LastWinAcctID string   `json:"last_win"`
	CurrentCup    int      `json:"current_cup"`
	Round         int      `json:"round"`
	Cups          []Cup    `json:"cups"`
	CupsOrder     []string `json:"player_order"`
	Bets          []Bet    `json:"bets"`
	Balances      []string `json:"balances"`
}

// Bet represents the bet response.
type Bet struct {
	AccountID string `json:"account"`
	Number    int    `json:"number"`
	Suite     int    `json:"suite"`
}

// Cup represents the cup response.
type Cup struct {
	AccountID string `json:"account"`
	Dice      []int  `json:"dice"`
	LastBet   Bet    `json:"last_bet"`
	Outs      int    `json:"outs"`
}
