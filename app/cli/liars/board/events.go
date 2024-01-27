package board

import (
	"fmt"
	"strings"

	"github.com/ardanlabs/liarsdice/app/cli/liars/engine"
	"github.com/ethereum/go-ethereum/common"
)

// webEvents handles any events from the websocket.
func (b *Board) webEvents(event string, address common.Address) {
	if !strings.Contains(event, "read tcp") {
		message := fmt.Sprintf("addr: %s type: %s", b.fmtAddress(address), event)
		b.printMessage(message, true)
	}

	var state engine.State
	var err error

	switch event {
	case "start":
		state, err = b.engine.RollDice(b.lastState.GameID)
		if err != nil {
			b.printMessage("error rolling dice", true)
		}

	case "rolldice":

		// Another player rolling the dice does not affect
		// our display.
		if address != b.accountID {
			return
		}

	case "callliar":
		state, err = b.modalWinnerLoser("*** WON ROUND ***", "*** LOST ROUND ***")
		if err != nil {
			b.printMessage("winner/loser", true)
		}

		state, err = b.reconcile(state)
		if err != nil {
			b.printMessage(err.Error(), true)
		}

	case "reconcile":
		b.modalWinnerLoser("*** WON GAME ***", "*** LOST GAME ***")
	}

	// If we don't have a new status, retrieve the latest.
	if state.Status == "" {
		state, err = b.engine.QueryState(b.lastState.GameID)
		if err != nil {
			return
		}
	}

	// Redraw the screen on any event to keep it up to date.
	b.drawBoard(state)
}

// reconcile the game the winner gets paid.
func (b *Board) reconcile(status engine.State) (engine.State, error) {
	if status.Status != "gameover" {
		return status, nil
	}

	if status.LastWinAcctID != b.accountID {
		return status, nil
	}

	newState, err := b.engine.Reconcile(status.GameID)
	if err != nil {
		return engine.State{}, err
	}

	return newState, nil
}
