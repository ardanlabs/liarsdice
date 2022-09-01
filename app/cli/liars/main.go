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
				case tcell.KeyCtrlL:
					s.Sync()
				case tcell.KeyRune:
					r := ev.Rune()
					switch {
					case r == rune('d'):
						enterDeposit(s, style)
					case r == rune('b'):
						enterBet(s, style)
					case (r >= rune('0') && r <= rune('9')) || r == rune('.'):
						key(s, style, ev.Rune())
					case r == rune('-'):
						minus(s, style)
					default:
						s.Beep()
					}
				}
			}
		}
	}()

	<-quit

	return nil
}

var bets []rune
var deposit []rune
var cursor string

func enter(s tcell.Screen, style tcell.Style) {
	switch cursor {
	case "bet":
		enterBet(s, style)

	case "deposit":
		enterDeposit(s, style)

	default:
		s.Beep()
	}
}

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

func key(s tcell.Screen, style tcell.Style, r rune) {
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

func enterDeposit(s tcell.Screen, style tcell.Style) {
	cursor = "deposit"
	deposit = []rune{}

	s.ShowCursor(12, 19)
	s.SetContent(12, 19, ' ', nil, style)
	s.SetCursorStyle(tcell.CursorStyleBlinkingBlock)

	emitStr(s, 12, 19, style, "                             ")
	emitStr(s, 13, 10, style, "                             ")
}

func addDeposit(s tcell.Screen, style tcell.Style, r rune) {
	if cursor != "deposit" {
		s.Beep()
		return
	}

	x := 11
	deposit = append(deposit, r)
	x += len(deposit)

	s.ShowCursor(x+1, 19)
	emitStr(s, x, 19, style, string(r))
}

func subDeposit(s tcell.Screen, style tcell.Style) {
	if cursor != "deposit" || len(deposit) == 0 {
		s.Beep()
		return
	}

	x := 11
	x += len(deposit)
	deposit = deposit[:len(deposit)-1]

	s.ShowCursor(x, 19)
	emitStr(s, x, 19, style, " ")
}

func enterBet(s tcell.Screen, style tcell.Style) {
	cursor = "bet"
	bets = []rune{}

	s.ShowCursor(13, 10)
	s.SetContent(13, 10, ' ', nil, style)
	s.SetCursorStyle(tcell.CursorStyleBlinkingBlock)

	emitStr(s, 12, 19, style, "                             ")
	emitStr(s, 13, 10, style, "                             ")
}

func subBet(s tcell.Screen, style tcell.Style) {
	if cursor != "bet" || len(bets) == 0 {
		s.Beep()
		return
	}

	x := 12
	x += len(bets)
	bets = bets[:len(bets)-1]

	s.ShowCursor(x, 10)
	emitStr(s, x, 10, style, " ")
}

func addBet(s tcell.Screen, style tcell.Style, r rune) {
	if cursor != "bet" || (len(bets) > 0 && bets[0] != r) {
		s.Beep()
		return
	}

	x := 12
	bets = append(bets, r)
	x += len(bets)

	s.ShowCursor(x+1, 10)
	emitStr(s, x, 10, style, string(r))
}

func makeBoard(s tcell.Screen, style tcell.Style) {
	s.Clear()

	for i := 1; i < 63; i++ {
		s.SetContent(i, 1, '=', nil, style)
	}
	for i := 1; i < 63; i++ {
		s.SetContent(i, 18, '=', nil, style)
	}
	for i := 2; i < 18; i++ {
		s.SetContent(1, i, '|', nil, style)
	}
	for i := 2; i < 18; i++ {
		s.SetContent(62, i, '|', nil, style)
	}

	for i := 1; i < 63; i++ {
		s.SetContent(i, 12, '=', nil, style)
	}
	emitStr(s, 3, 12, style, " Message Center ")

	emitStr(s, 3, 2, style, "Players:")
	emitStr(s, 3, 4, style, "   Me (0x6327A384)")
	emitStr(s, 3, 5, style, "   0x8e113078")
	emitStr(s, 3, 6, style, "-> 0x0070742f")

	emitStr(s, 30, 2, style, "Last Bet:")
	emitStr(s, 30, 4, style, "5 Two's")
	emitStr(s, 30, 5, style, "6 One's")
	emitStr(s, 30, 6, style, "")

	emitStr(s, 50, 2, style, "  Balances:")
	emitStr(s, 50, 4, style, "  $1032 USD")
	emitStr(s, 50, 5, style, "   $864 USD")
	emitStr(s, 50, 6, style, "$12,000 USD")

	emitStr(s, 3, 8, style, "My Dice: [3] [2] [6] [6] [1]")
	emitStr(s, 35, 8, style, "Ante:  $5 USD")
	emitStr(s, 35, 9, style, "Pot : $15 USD")

	emitStr(s, 3, 10, style, "My Bet :>")
	emitStr(s, 1, 19, style, "Deposit :>")

	emitStr(s, 65, 2, style, "<1-6>   : set/increment bet")
	emitStr(s, 65, 3, style, "<minus> : decrement bet")
	emitStr(s, 65, 4, style, "<b>     : place bet")
	emitStr(s, 65, 5, style, "<l>     : call liar")
	emitStr(s, 65, 6, style, "<d>     : deposit funds")
}

func message(s tcell.Screen, style tcell.Style, message string) {
	emitStr(s, 3, 14, style, message)
	s.Show()
}

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
