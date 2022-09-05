package board

import (
	"errors"
	"fmt"
	"strings"
)

// Events handles any events from the websocket.
func (b *Board) Events(event string, address string) {
	if !strings.Contains(address, "read tcp") {
		message := fmt.Sprintf("addr: %s type: %s", b.fmtAddress(address), event)
		b.printMessage(message, true)
	}

	switch event {
	case "start":
		if _, err := b.engine.RollDice(); err != nil {
			b.printMessage("error rolling dice", true)
		}

	case "callliar":
		b.modalWinnerLoser("*** WON ROUND ***", "*** LOST ROUND ***")

		if err := b.reconcile(); err != nil {
			b.printMessage(err.Error(), true)
		}

	case "reconcile":
		b.modalWinnerLoser("*** WON GAME ***", "*** LOST GAME ***")
	}

	status, err := b.engine.QueryStatus()
	if err != nil {
		return
	}
	b.printStatus(status)
}

// =============================================================================

// reconcile the game the winner gets paid.
func (b *Board) reconcile() error {
	status, err := b.engine.QueryStatus()
	if err != nil {
		return err
	}

	if status.Status != "gameover" {
		return errors.New("invalid status state: " + status.Status)
	}

	if status.LastWinAcctID != b.accountID {
		return errors.New("gameover: winner will call reconcile")
	}

	if _, err := b.engine.Reconcile(); err != nil {
		return err
	}

	return nil
}
