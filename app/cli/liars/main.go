package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalln(err)
	}

	if err := s.Init(); err != nil {
		log.Fatalln(err)
	}
	defer s.Fini()

	style := tcell.StyleDefault
	style = style.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	makeBoard(s, style)

	quit := make(chan struct{})
	go func() {
		for {
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape:
					close(quit)
					return

				case tcell.KeyEnter:
					enter(s, style)

				case tcell.KeyRune:
					r := ev.Rune()
					processKeyEvent(s, style, r)
				}
			}
		}
	}()

	<-quit

	return nil
}

// processKeyEvent is the first line of processing for any key that is
// pressed during the game.
func processKeyEvent(s tcell.Screen, style tcell.Style, r rune) {
	switch {
	case r == rune('d'):
		depositMode(s, style)

	case r == rune('b'):
		betMode(s, style)

	case (r >= rune('0') && r <= rune('9')) || r == rune('.'):
		value(s, style, r)

	case r == rune('-'):
		minus(s, style)

	default:
		s.Beep()
	}
}

// =============================================================================

var bets []rune
var deposit []rune
var cursor string

var depositRowX = 11
var depositRowY = 19
var betRowX = 13
var betRowY = 10
var boardWidth = 63
var boardHeight = 18

var messageHeight = 12
var columnHeight = 2
var playersX = 3
var betX = 30
var balX = 50
var myDiceX = 12
var myDiceY = 8
var anteX = 40
var anteY = 8
var potX = 40
var potY = 9
var helpX = 65

// =============================================================================

// enter is called to submit a bet or deposit.
func enter(s tcell.Screen, style tcell.Style) {
	switch cursor {
	case "bet":
		betMode(s, style)

	case "deposit":
		depositMode(s, style)

	default:
		s.Beep()
	}
}

// minus is called to remove the latest value from the bet or deposit.
func minus(s tcell.Screen, style tcell.Style) {
	switch cursor {
	case "bet":
		subBet(s, style)

	case "deposit":
		subDeposit(s, style)

	default:
		s.Beep()
	}
}

// value processes the keystroke based on the mode.
func value(s tcell.Screen, style tcell.Style, r rune) {
	switch cursor {
	case "bet":
		if r >= rune('1') && r <= rune('6') {
			addBet(s, style, r)
			return
		}
		s.Beep()

	case "deposit":
		addDeposit(s, style, r)

	default:
		s.Beep()
	}
}

// =============================================================================

// depositMode puts the UI into the mode to accept deposit information and
// process a deposit.
func depositMode(s tcell.Screen, style tcell.Style) {
	cursor = "deposit"
	deposit = []rune{}

	s.ShowCursor(depositRowX+1, depositRowY)
	s.SetContent(depositRowX, depositRowY, ' ', nil, style)
	s.SetCursorStyle(tcell.CursorStyleBlinkingBlock)

	emitStr(s, depositRowX, depositRowY, style, "                             ")
	emitStr(s, betRowX, betRowY, style, "                             ")
}

// addDeposit takes the value selected on the keyboard and adds it to the
// deposit slice and screen.
func addDeposit(s tcell.Screen, style tcell.Style, r rune) {
	if cursor != "deposit" {
		s.Beep()
		return
	}

	x := depositRowX
	deposit = append(deposit, r)
	x += len(deposit)

	s.ShowCursor(x+1, depositRowY)
	emitStr(s, x, depositRowY, style, string(r))
}

// subDeposit removes a value from the deposit slice and screen.
func subDeposit(s tcell.Screen, style tcell.Style) {
	if cursor != "deposit" || len(deposit) == 0 {
		s.Beep()
		return
	}

	x := depositRowX
	x += len(deposit)
	deposit = deposit[:len(deposit)-1]

	s.ShowCursor(x, depositRowY)
	emitStr(s, x, depositRowY, style, " ")
}

// =============================================================================

// betMode puts the UI into the mode to accept bet information and
// process a bet.
func betMode(s tcell.Screen, style tcell.Style) {
	cursor = "bet"
	bets = []rune{}

	s.ShowCursor(betRowX+1, betRowY)
	s.SetContent(betRowX, betRowY, ' ', nil, style)
	s.SetCursorStyle(tcell.CursorStyleBlinkingBlock)

	emitStr(s, depositRowX, depositRowY, style, "                             ")
	emitStr(s, betRowX, betRowY, style, "                             ")
}

// addBet takes the value selected on the keyboard and adds it to the
// bet slice and screen.
func addBet(s tcell.Screen, style tcell.Style, r rune) {
	if cursor != "bet" || (len(bets) > 0 && bets[0] != r) {
		s.Beep()
		return
	}

	x := betRowX
	bets = append(bets, r)
	x += len(bets)

	s.ShowCursor(x+1, betRowY)
	emitStr(s, x, betRowY, style, string(r))
}

// subBet removes a value from the bet slice and screen.
func subBet(s tcell.Screen, style tcell.Style) {
	if cursor != "bet" || len(bets) == 0 {
		s.Beep()
		return
	}

	x := betRowX
	x += len(bets)
	bets = bets[:len(bets)-1]

	s.ShowCursor(x, betRowY)
	emitStr(s, x, betRowY, style, " ")
}

// =============================================================================

// makeBoard generates the initial game board.
func makeBoard(s tcell.Screen, style tcell.Style) {
	s.Clear()

	for i := 1; i < boardWidth; i++ {
		s.SetContent(i, 1, '=', nil, style)
	}
	for i := 1; i < boardWidth; i++ {
		s.SetContent(i, boardHeight, '=', nil, style)
	}
	for i := 2; i < boardHeight; i++ {
		s.SetContent(1, i, '|', nil, style)
	}
	for i := 2; i < boardHeight; i++ {
		s.SetContent(boardWidth-1, i, '|', nil, style)
	}

	for i := 1; i < boardWidth; i++ {
		s.SetContent(i, messageHeight, '=', nil, style)
	}
	emitStr(s, 3, messageHeight, style, " Message Center ")

	emitStr(s, playersX, columnHeight, style, "Players:")

	emitStr(s, betX, columnHeight, style, "Last Bet:")

	emitStr(s, balX, columnHeight, style, "  Balances:")

	emitStr(s, myDiceX-8, myDiceY, style, "My Dice:")

	emitStr(s, anteX-6, anteY, style, "Ante:")
	emitStr(s, potX-6, potY, style, "Pot :")

	emitStr(s, betRowX-9, betRowY, style, "My Bet :>")
	emitStr(s, depositRowX-10, depositRowY, style, "Deposit :>")

	emitStr(s, helpX, 2, style, "<1-6>   : set/increment bet")
	emitStr(s, helpX, 3, style, "<minus> : decrement bet")
	emitStr(s, helpX, 4, style, "<b>     : place bet")
	emitStr(s, helpX, 5, style, "<l>     : call liar")
	emitStr(s, helpX, 6, style, "<d>     : deposit funds")
}

// message adds a message to the message center.
func message(s tcell.Screen, style tcell.Style, message string) {
	emitStr(s, 3, messageHeight+2, style, message)
	s.Show()
}

// emitStr knows how to print a string on the screen.
func emitStr(s tcell.Screen, x, y int, style tcell.Style, str string) {
	for _, c := range str {
		var comb []rune
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}
		s.SetContent(x, y, c, comb, style)
		x += w
	}
	s.Show()
}
