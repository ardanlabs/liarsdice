package client

// ErrorResponse is the form used for API responses from failures in the API.
type ErrorResponse struct {
	Error  string            `json:"error"`
	Fields map[string]string `json:"fields,omitempty"`
}

// Config represents the configuration of the game engine.
type Config struct {
	Network    string `json:"network"`
	ChainID    int    `json:"chain_id"`
	ContractID string `json:"contract_id"`
}

// Token contains the user game token and public address of the player.
type Token struct {
	Token   string `json:"token"`
	Address string `json:"address"`
}

// Cup represents the cup response.
type Cup struct {
	Account string `json:"account"`
	Outs    int    `json:"outs"`
}

// Bet represents the bet response.
type Bet struct {
	Account string `json:"account"`
	Number  int    `json:"number"`
	Suite   int    `json:"suite"`
}

type Status struct {
	Status        string   `json:"status"`
	AnteUSD       float64  `json:"ante_usd"`
	LastOutAcct   string   `json:"last_out"`
	LastWinAcct   string   `json:"last_win"`
	CurrentPlayer string   `json:"current_player"`
	CurrentCup    int      `json:"current_cup"`
	Round         int      `json:"round"`
	Cups          []Cup    `json:"cups"`
	CupsOrder     []string `json:"player_order"`
	Bets          []Bet    `json:"bets"`
}
