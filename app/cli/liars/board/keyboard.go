package board

import (
	"errors"

	"github.com/gdamore/tcell/v2"
)

// pollEvents starts a goroutine to handle terminal events.
func (b *Board) pollEvents() chan struct{} {
	quit := make(chan struct{})

	go func() {
		for {
			event := b.screen.PollEvent()

			// Check if we received a key event.
			ev, isEventKey := event.(*tcell.EventKey)
			if !isEventKey {
				continue
			}

			// Check if the escape key was selected.
			keyType := ev.Key()
			if keyType == tcell.KeyEscape {
				if b.modalUp {
					b.closeModal()
					continue
				}
				close(quit)
				return
			}

			// If the modal is up, ignore the keystroke.
			if b.modalUp {
				b.screen.Beep()
				continue
			}

			// Process the specified keys.
			var err error
			switch keyType {
			case tcell.KeyDEL, tcell.KeyDelete:
				err = b.subBet()

			case tcell.KeyEnter:
				err = b.enterBet()

			case tcell.KeyRune:
				err = b.processKeyEvent(ev.Rune())
			}

			// Print an error message that was returned.
			if err != nil {
				b.printMessage(err.Error(), true)
			}
		}
	}()

	return quit
}

// processKeyEvent is the first line of processing for any key that is
// pressed during the game.
func (b *Board) processKeyEvent(r rune) error {
	var err error

	switch {
	case (r >= rune('0') && r <= rune('9')) || r == rune('.'):
		err = b.value(r)

	case r == rune('n'):
		err = b.newGame()

	case r == rune('j'):
		err = b.joinGame()

	case r == rune('s'):
		err = b.startGame()

	case r == rune('l'):
		err = b.callLiar()

	default:
		err = errors.New("invalid selection")
	}

	return err
}

// value processes the keystroke based on the mode.
func (b *Board) value(r rune) error {
	if r >= rune('1') && r <= rune('6') {
		return b.addBet(r)
	}

	return errors.New("invalid selection")
}

// newGame starts a new game.
func (b *Board) newGame() error {
	if _, err := b.engine.NewGame(); err != nil {
		return err
	}

	return nil
}

// joinGame adds the account to the game.
func (b *Board) joinGame() error {
	status, err := b.engine.QueryStatus()
	if err != nil {
		return err
	}

	if status.Status != "newgame" {
		return errors.New("invalid status state: " + status.Status)
	}

	if _, err = b.engine.JoinGame(); err != nil {
		return err
	}

	return nil
}

// startGame start the game so it can be played.
func (b *Board) startGame() error {
	status, err := b.engine.QueryStatus()
	if err != nil {
		return err
	}

	if status.Status != "newgame" {
		return errors.New("invalid status state: " + status.Status)
	}

	if _, err := b.engine.StartGame(); err != nil {
		return err
	}

	return nil
}

// callLiar calls the last bet a lie.
func (b *Board) callLiar() error {
	status, err := b.engine.QueryStatus()
	if err != nil {
		return err
	}

	if status.Status != "playing" {
		return errors.New("invalid status state: " + status.Status)
	}

	if status.CupsOrder[status.CurrentCup] != b.accountID {
		return errors.New("not your turn")
	}

	if _, err := b.engine.Liar(); err != nil {
		return err
	}

	return nil
}
