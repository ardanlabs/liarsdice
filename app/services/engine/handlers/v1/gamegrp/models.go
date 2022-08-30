package gamegrp

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
