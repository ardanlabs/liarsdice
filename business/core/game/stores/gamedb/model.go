package gamedb

import (
	"fmt"
	"time"

	"github.com/ardanlabs/liarsdice/business/core/game"
	"github.com/ardanlabs/liarsdice/business/data/sqldb/dbarray"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
)

type dbState struct {
	ID              uuid.UUID      `db:"game_id"`
	Name            string         `db:"name"`
	DateCreated     time.Time      `db:"date_created"`
	Round           int            `db:"round"`
	Status          string         `db:"status"`
	PlayerLastOut   string         `db:"player_last_out"`
	PlayerLastWin   string         `db:"player_last_win"`
	PlayerTurn      string         `db:"player_turn"`
	ExistingPlayers dbarray.String `db:"existing_players"`
	Cups            []dbCup
	Bets            []dbBet
	Balances        []dbBalance
}

type dbCup struct {
	ID       uuid.UUID `db:"game_id"`
	Round    int       `db:"round"`
	Player   string    `db:"player"`
	OrderIdx int       `db:"order_idx"`
	Outs     int       `db:"outs"`
	Dice     any       `db:"dice"`
}

type dbBet struct {
	ID       uuid.UUID `db:"game_id"`
	Round    int       `db:"round"`
	BetOrder int       `db:"bet_order"`
	Player   string    `db:"player"`
	Number   int       `db:"number"`
	Suit     int       `db:"suit"`
}

type dbBalance struct {
	ID     uuid.UUID `db:"game_id"`
	Round  int       `db:"round"`
	Player string    `db:"player"`
	Amount string    `db:"amount"`
}

func toDBState(state game.State) dbState {
	cups := make([]dbCup, 0, len(state.Cups))
	for _, cup := range state.Cups {
		cups = append(cups, dbCup{
			ID:       state.GameID,
			Round:    state.Round,
			Player:   cup.Player.String(),
			OrderIdx: cup.OrderIdx,
			Outs:     cup.Outs,
			Dice:     dbarray.Array(cup.Dice),
		})
	}

	bets := make([]dbBet, len(state.Bets))
	for i, bet := range state.Bets {
		bets[i] = dbBet{
			ID:       state.GameID,
			Round:    state.Round,
			BetOrder: i,
			Player:   bet.Player.String(),
			Number:   bet.Number,
			Suit:     bet.Suit,
		}
	}

	balances := make([]dbBalance, len(state.Balances))
	for i, balance := range state.Balances {
		balances[i] = dbBalance{
			ID:     state.GameID,
			Round:  state.Round,
			Player: balance.Player.String(),
			Amount: balance.Amount,
		}
	}

	existingPlayers := make([]string, len(state.ExistingPlayers))
	for i, ep := range state.ExistingPlayers {
		existingPlayers[i] = ep.String()
	}

	return dbState{
		ID:              state.GameID,
		Name:            state.GameID.String(),
		DateCreated:     state.DateCreated,
		Round:           state.Round,
		Status:          state.Status,
		PlayerLastOut:   state.PlayerLastOut.String(),
		PlayerLastWin:   state.PlayerLastWin.String(),
		PlayerTurn:      state.PlayerTurn.String(),
		ExistingPlayers: existingPlayers,
		Cups:            cups,
		Bets:            bets,
		Balances:        balances,
	}
}

func toCoreState(dbState dbState) (game.State, error) {
	if !common.IsHexAddress(dbState.PlayerLastOut) {
		return game.State{}, fmt.Errorf("player last out is not a common address: %q", dbState.PlayerLastOut)
	}

	if !common.IsHexAddress(dbState.PlayerLastWin) {
		return game.State{}, fmt.Errorf("player last win is not a common address: %q", dbState.PlayerLastWin)
	}

	if !common.IsHexAddress(dbState.PlayerTurn) {
		return game.State{}, fmt.Errorf("player turn is not a common address: %q", dbState.PlayerTurn)
	}

	existingPlayers := make([]common.Address, len(dbState.ExistingPlayers))
	for i, player := range dbState.ExistingPlayers {
		existingPlayers[i] = common.HexToAddress(player)
	}

	state := game.State{
		GameID:          dbState.ID,
		GameName:        dbState.Name,
		DateCreated:     dbState.DateCreated,
		Round:           dbState.Round,
		Status:          dbState.Status,
		PlayerLastOut:   common.HexToAddress(dbState.PlayerLastOut),
		PlayerLastWin:   common.HexToAddress(dbState.PlayerLastWin),
		PlayerTurn:      common.HexToAddress(dbState.PlayerTurn),
		ExistingPlayers: existingPlayers,
	}

	return state, nil
}
