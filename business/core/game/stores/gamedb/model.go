package gamedb

type dbGame struct {
	ID   string `db:"game_id"`
	Name string `db:"name"`
}

type dbGameStatus struct {
	ID            string `db:"game_id"`
	Iteration     int    `db:"iteration"`
	Status        string `db:"status"`
	PlayerLastOut string `db:"player_last_out"`
	PlayerLastWin string `db:"player_last_win"`
	PlayerTurn    string `db:"player_turn"`
	Round         int    `db:"round"`
}

type dbGameCup struct {
	ID        string `db:"game_id"`
	Iteration int    `db:"iteration"`
	Player    string `db:"player"`
	OrderIdx  int    `db:"order_idx"`
	Outs      int    `db:"outs"`
}

type dbGameDice struct {
	ID        string `db:"game_id"`
	Iteration int    `db:"iteration"`
	Player    string `db:"player"`
	Dice      int    `db:"dice"`
}

type dbGameExistingPlayers struct {
	ID        string `db:"game_id"`
	Iteration int    `db:"iteration"`
	Player    string `db:"player"`
}

type dbGameBets struct {
	ID        string `db:"game_id"`
	Iteration int    `db:"iteration"`
	Player    string `db:"player"`
	Number    int    `db:"number"`
	Suit      int    `db:"suit"`
}

type dbGameBalances struct {
	ID        string `db:"game_id"`
	Iteration int    `db:"iteration"`
	Player    string `db:"player"`
	Balance   string `db:"balance"`
}

// func toDBGame(status game.State)
