package client

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
