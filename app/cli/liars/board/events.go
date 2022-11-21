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

	var status engine.Status
	var err error

	switch event {
	case "start":
		status, err = b.engine.RollDice()
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
		status, err = b.modalWinnerLoser("*** WON ROUND ***", "*** LOST ROUND ***")
		if err != nil {
			b.printMessage("winner/loser", true)
		}

		status, err = b.reconcile(status)
		if err != nil {
			b.printMessage(err.Error(), true)
		}

	case "reconcile":
		b.modalWinnerLoser("*** WON GAME ***", "*** LOST GAME ***")
	}

	// If we don't have a new status, retrieve the latest.
	if status.Status == "" {
		status, err = b.engine.QueryStatus()
		if err != nil {
			return
		}
	}

	// Redraw the screen on any event to keep it up to date.
	b.drawBoard(status)
}

// reconcile the game the winner gets paid.
func (b *Board) reconcile(status engine.Status) (engine.Status, error) {
	if status.Status != "gameover" {
		return status, nil
	}

	if status.LastWinAcctID != b.accountID {
		return status, nil
	}

	newStatus, err := b.engine.Reconcile()
	if err != nil {
		return engine.Status{}, err
	}

	return newStatus, nil
}
