package gamegrp

import (
	"time"

	"github.com/ardanlabs/liarsdice/business/core/game"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
)

type appState struct {
	GameID          uuid.UUID        `json:"gameID"`
	GameName        string           `json:"gameName"`
	DateCreated     string           `json:"dateCreated"`
	AnteUSD         float64          `json:"anteUSD"`
	Status          string           `json:"status"`
	PlayerLastOut   common.Address   `json:"lastOut"`
	PlayerLastWin   common.Address   `json:"lastWin"`
	PlayerTurn      common.Address   `json:"currentID"`
	Round           int              `json:"round"`
	Cups            []appCup         `json:"cups"`
	ExistingPlayers []common.Address `json:"playerOrder"`
	Bets            []appBet         `json:"bets"`
	Balances        []string         `json:"balances"`
}

func toAppState(state game.State, anteUSD float64, address common.Address) appState {
	var cups []appCup
	for _, accountID := range state.ExistingPlayers {
		cup := state.Cups[accountID]

		// Don't share the dice information for other players.
		dice := []int{0, 0, 0, 0, 0}
		if accountID == address {
			dice = cup.Dice
		}
		cups = append(cups, toAppCup(cup, dice))
	}

	var bets []appBet
	for _, bet := range state.Bets {
		bets = append(bets, toAppBet(bet))
	}

	var balances []string
	for _, balance := range state.Balances {
		balances = append(balances, balance.Amount)
	}

	return appState{
		GameID:          state.GameID,
		GameName:        state.GameName,
		DateCreated:     state.DateCreated.Format(time.RFC3339),
		AnteUSD:         anteUSD,
		Status:          state.Status,
		PlayerLastOut:   state.PlayerLastOut,
		PlayerLastWin:   state.PlayerLastWin,
		PlayerTurn:      state.PlayerTurn,
		Round:           state.Round,
		Cups:            cups,
		ExistingPlayers: state.ExistingPlayers,
		Bets:            bets,
		Balances:        balances,
	}
}

type appBet struct {
	Player common.Address `json:"account"`
	Number int            `json:"number"`
	Suit   int            `json:"suit"`
}

func toAppBet(bet game.Bet) appBet {
	return appBet{
		Player: bet.Player,
		Number: bet.Number,
		Suit:   bet.Suit,
	}
}

type appCup struct {
	Player common.Address `json:"account"`
	Dice   []int          `json:"dice"`
	Outs   int            `json:"outs"`
}

func toAppCup(cup game.Cup, dice []int) appCup {
	return appCup{
		Player: cup.Player,
		Dice:   dice,
		Outs:   cup.Outs,
	}
}
