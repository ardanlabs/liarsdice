package gamedb

import (
	"time"

	"github.com/ardanlabs/liarsdice/business/core/game"
	"github.com/ardanlabs/liarsdice/business/data/sqldb/dbarray"
)

type dbGame struct {
	ID          string    `db:"game_id"`
	Name        string    `db:"name"`
	CreatedDate time.Time `db:"created_date"`
	State       dbState
	Cups        []dbCup
	Bets        []dbBet
	Balances    []dbBalance
}

type dbState struct {
	Round           int    `db:"round"`
	Status          string `db:"status"`
	PlayerLastOut   string `db:"player_last_out"`
	PlayerLastWin   string `db:"player_last_win"`
	PlayerTurn      string `db:"player_turn"`
	ExistingPlayers any    `db:"existing_players"`
}

type dbCup struct {
	Round    int    `db:"round"`
	Player   string `db:"player"`
	OrderIdx int    `db:"order_idx"`
	Outs     int    `db:"outs"`
	Dice     any    `db:"dice"`
}

type dbBet struct {
	Round  int    `db:"round"`
	Player string `db:"player"`
	Number int    `db:"number"`
	Suit   int    `db:"suit"`
}

type dbBalance struct {
	Round  int    `db:"round"`
	Player string `db:"player"`
	Amount string `db:"amount"`
}

func toDBGame(g *game.Game, state game.State) dbGame {
	cups := make([]dbCup, len(state.Cups))
	for player, cup := range state.Cups {
		cups = append(cups, dbCup{
			Round:    state.Round,
			Player:   player.String(),
			OrderIdx: cup.OrderIdx,
			Outs:     cup.Outs,
			Dice:     dbarray.Array(cup.Dice),
		})
	}

	bets := make([]dbBet, len(state.Bets))
	for _, bet := range state.Bets {
		bets = append(bets, dbBet{
			Round:  state.Round,
			Player: bet.Player.String(),
			Number: bet.Number,
			Suit:   bet.Suit,
		})
	}

	balances := make([]dbBalance, len(state.Balances))
	for _, balance := range state.Balances {
		balances = append(balances, dbBalance{
			Round:  state.Round,
			Player: balance.Player.String(),
			Amount: balance.Amount,
		})
	}

	return dbGame{
		ID:          state.GameID,
		Name:        state.GameID,
		CreatedDate: g.CreatedDate(),
		State: dbState{
			Round:           state.Round,
			Status:          state.Status,
			PlayerLastOut:   state.PlayerLastOut.String(),
			PlayerLastWin:   state.PlayerLastWin.String(),
			PlayerTurn:      state.PlayerTurn.String(),
			ExistingPlayers: dbarray.Array(state.ExistingPlayers),
		},
		Cups:     cups,
		Bets:     bets,
		Balances: balances,
	}
}
