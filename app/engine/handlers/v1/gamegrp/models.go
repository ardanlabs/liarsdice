package gamegrp

type Game struct {
	ID      string   `json:"id"`
	Status  string   `json:"status"`
	Players []Player `json:"players,omitempty"`
}

type Player struct {
	Wallet string `json:"wallet"`
	Out    uint8  `json:"out"`
}

type Table struct {
	Games map[string]Game `json:"games"`
}

func NewTable() Table {
	table := Table{}
	table.Games = make(map[string]Game)

	return table
}
