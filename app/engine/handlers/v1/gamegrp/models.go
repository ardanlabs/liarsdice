package gamegrp

type Game struct {
	ID            string   `json:"id"`
	Status        string   `json:"status"`
	Players       []Player `json:"players,omitempty"`
	CurrentPlayer int      `json:"current_player"`
	Claims        []Claim  `json:"claims"`
}

type Player struct {
	Wallet string `json:"wallet"`
	Outs   uint8  `json:"outs"`
	Dices  []int  `json:"dices,omitempty"`
}

type Table struct {
	Games map[string]Game `json:"games"`
}

func NewTable() Table {
	table := Table{}
	table.Games = make(map[string]Game)

	return table
}

type Claim struct {
	Wallet string `json:"wallet"`
	Number int    `json:"number"`
	Suite  int    `json:"suite"`
}
