package board

import (
	"fmt"
	"strings"

	"github.com/ardanlabs/liarsdice/app/cli/liars/engine"
)

// webEvents handles any events from the websocket.
func (b *Board) webEvents(event string, address string) {
	if !strings.Contains(address, "read tcp") {
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
		b.modalWinnerLoser("*** WON ROUND ***", "*** LOST ROUND ***")

		status, err = b.reconcile()
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
	b.printStatus(status)
}

// reconcile the game the winner gets paid.
func (b *Board) reconcile() (engine.Status, error) {
	status, err := b.engine.QueryStatus()
	if err != nil {
		return status, err
	}

	if status.Status != "gameover" {
		return status, nil
	}

	if status.LastWinAcctID != b.accountID {
		return status, nil
	}

	if _, err := b.engine.Reconcile(); err != nil {
		return engine.Status{}, err
	}

	return status, nil
}
