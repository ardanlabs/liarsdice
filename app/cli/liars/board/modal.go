package board

import (
	"strings"

	"github.com/ardanlabs/liarsdice/app/cli/liars/engine"
)

// modalWinnerLoser shows the user if they won or lost.
func (b *Board) modalWinnerLoser(win string, los string) (engine.Status, error) {
	status, err := b.engine.QueryStatus()
	if err != nil {
		return engine.Status{}, err
	}

	if status.LastWinAcctID == b.accountID {
		b.showModal(win)
		return status, nil
	}
	b.showModal(los)

	return status, nil
}

// showModal displays a modal dialog box.
func (b *Board) showModal(message string) error {
	b.modalUp = true
	b.modalMsg = message

	b.screen.HideCursor()
	b.drawBox(9, 3, 55, 8)

	words := strings.Split(message, " ")
	var msg string
	h := 5
	for _, word := range words {
		msg += word + " "
		if len(msg) >= 40 {
			l := len(msg)
			x := 32 - (l / 2)
			b.print(x, h, msg)
			h++
			msg = ""
		}
	}
	l := len(msg)
	x := 32 - (l / 2)
	b.print(x, h, msg)

	return nil
}

// closeModal closes the modal dialog box.
func (b *Board) closeModal() {
	b.modalUp = false
	b.modalMsg = ""

	active := false
	if b.lastStatus.CurrentAcctID == b.accountID {
		active = true
	}

	b.drawGameBox(active)
	b.drawLables()
	b.drawBoard(b.lastStatus)
}
