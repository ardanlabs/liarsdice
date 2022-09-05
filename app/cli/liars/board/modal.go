package board

import (
	"strings"
)

// modalWinnerLoser shows the user if they won or lost.
func (b *Board) modalWinnerLoser(win string, los string) error {
	status, err := b.engine.QueryStatus()
	if err != nil {
		return err
	}

	if status.LastWinAcctID == b.accountID {
		b.showModal(win)
		return nil
	}
	b.showModal(los)

	return nil
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
	if b.lastStatus.CupsOrder[b.lastStatus.CurrentCup] == b.accountID {
		active = true
	}

	b.drawGameBoard(active)
	b.printLables()
	b.printStatus(b.lastStatus)
}
