package engine

// ErrorResponse is the form used for API responses from failures in the API.
type ErrorResponse struct {
	Error  string            `json:"error"`
	Fields map[string]string `json:"fields,omitempty"`
}

// Config represents the configuration of the game engine.
type Config struct {
	Network    string `json:"network"`
	ChainID    int    `json:"chainId"`
	ContractID string `json:"contractId"`
}

// Token contains the user game token and public address of the player.
type Token struct {
	Token   string `json:"token"`
	Address string `json:"address"`
}

// Status represents the game status.
type Status struct {
	Status        string   `json:"status"`
	AnteUSD       float64  `json:"anteUSD"`
	LastOutAcctID string   `json:"lastOut"`
	LastWinAcctID string   `json:"lastWin"`
	CurrentAcctID string   `json:"currentID"`
	Round         int      `json:"round"`
	Cups          []Cup    `json:"cups"`
	CupsOrder     []string `json:"playerOrder"`
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
	LastBet   Bet    `json:"lastBet"`
	Outs      int    `json:"outs"`
}
