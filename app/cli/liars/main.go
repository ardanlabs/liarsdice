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
					message(s, style, "didn't provide proper bet")
				case tcell.KeyCtrlL:
					s.Sync()
				}
			}
		}
	}()

	<-quit

	return nil
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
	emitStr(s, 35, 8, style, "Pot: $15 USD")

	emitStr(s, 3, 10, style, "My Bet: ___ ______")

	emitStr(s, 65, 2, style, "<1-6>   : multiple times to set bet")
	emitStr(s, 65, 3, style, "<minus> : decrement bet")
	emitStr(s, 65, 4, style, "<enter> : place bet")
	emitStr(s, 65, 5, style, "<l>     : call liar")

	s.Show()
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
}
