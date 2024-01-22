package board

import (
	"errors"
	"fmt"
	"strconv"
)

// addBet takes the value selected on the keyboard and adds it to the
// bet slice and screen.
func (b *Board) addBet(r rune) error {
	if b.lastStatus.CurrentAcctID != b.accountID {
		return errors.New("not your turn")
	}

	if len(b.bets) > 0 && b.bets[0] != r {
		return errors.New("please use the same value")
	}

	x := betRowX
	b.bets = append(b.bets, r)
	x += len(b.bets)

	b.screen.ShowCursor(x+1, betRowY)
	b.print(x, betRowY, string(r))

	suit, err := strconv.Atoi(string(b.bets[0]))
	if err != nil {
		return err
	}

	bet := fmt.Sprintf("%d %-10s", len(b.bets), words[suit])
	b.print(potX, potY+1, bet)

	return nil
}

// subBet removes a value from the bet slice and screen.
func (b *Board) subBet() error {
	if b.lastStatus.CurrentAcctID != b.accountID {
		return errors.New("not your turn")
	}

	if len(b.bets) == 0 {
		return errors.New("nothing to delete")
	}

	x := betRowX
	x += len(b.bets)
	b.bets = b.bets[:len(b.bets)-1]

	b.screen.ShowCursor(x, betRowY)
	b.print(x, betRowY, " ")

	bet := "                 "
	if len(b.bets) > 0 {
		suit, err := strconv.Atoi(string(b.bets[0]))
		if err != nil {
			return err
		}

		bet = fmt.Sprintf("%d %-10s", len(b.bets), words[suit])
	}
	b.print(potX, potY+1, bet)

	return nil
}

// enterBet is called to submit a bet.
func (b *Board) enterBet() error {
	status, err := b.engine.QueryStatus()
	if err != nil {
		return err
	}

	if status.Status != "playing" {
		return errors.New("invalid status state: " + status.Status)
	}

	if status.CurrentAcctID != b.accountID {
		return errors.New("not your turn")
	}

	if len(b.bets) == 0 {
		return errors.New("missing bet information")
	}

	if _, err = b.engine.Bet(len(b.bets), b.bets[0]); err != nil {
		return err
	}

	b.bets = []rune{}
	b.screen.ShowCursor(betRowX+1, betRowY)
	b.print(betRowX, betRowY, "                 ")
	b.print(potX, potY+1, "                 ")

	return nil
}
