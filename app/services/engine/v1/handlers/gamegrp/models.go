package gamegrp

import (
	"github.com/ardanlabs/liarsdice/business/core/game"
	"github.com/ethereum/go-ethereum/common"
)

type appState struct {
	GameID          string           `json:"gameID"`
	Status          string           `json:"status"`
	AnteUSD         float64          `json:"anteUSD"`
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
		Status:          state.Status,
		AnteUSD:         anteUSD,
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
	Player  common.Address `json:"account"`
	Dice    []int          `json:"dice"`
	LastBet appBet         `json:"lastBet"`
	Outs    int            `json:"outs"`
}

func toAppCup(cup game.Cup, dice []int) appCup {
	return appCup{
		Player: cup.Player,
		Dice:   dice,
		Outs:   cup.Outs,
	}
}
