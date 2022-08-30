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

	quit := make(chan struct{})
	go func() {
		for {
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyEnter:
					close(quit)
					return
				case tcell.KeyCtrlL:
					s.Sync()
				}
			}
		}
	}()

	makeBoard(s)

	<-quit

	return nil
}

func makeBoard(s tcell.Screen) {
	s.Clear()

	_, h := s.Size()
	w := 63

	style := tcell.StyleDefault
	style = style.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)

	// Outer frame of the game board.

	for i := 1; i < w; i++ {
		s.SetContent(i, 1, '=', nil, style)
	}
	for i := 1; i < w; i++ {
		s.SetContent(i, h-2, '=', nil, style)
	}
	for i := 2; i < h-2; i++ {
		s.SetContent(1, i, '|', nil, style)
	}
	for i := 2; i < h-2; i++ {
		s.SetContent(w-1, i, '|', nil, style)
	}

	// Message center.
	for i := 1; i < w; i++ {
		s.SetContent(i, h-6, '=', nil, style)
	}
	emitStr(s, 3, h-6, style, " Message Center ")

	// Gaming controls.
	for i := 1; i < w; i++ {
		s.SetContent(i, h-10, '=', nil, style)
	}
	emitStr(s, 3, h-10, style, " Gaming Controls ")

	// Players
	emitStr(s, 3, 2, style, "Players:")
	emitStr(s, 3, 4, style, "   Me (0x6327A384)")
	emitStr(s, 3, 5, style, "   0x8e113078")
	emitStr(s, 3, 6, style, "-> 0x0070742f")

	// Bets
	emitStr(s, 30, 2, style, "Last Bet:")
	emitStr(s, 30, 4, style, "5 Two's")
	emitStr(s, 30, 5, style, "6 One's")
	emitStr(s, 30, 6, style, "")

	// Balances
	emitStr(s, 50, 2, style, "  Balances:")
	emitStr(s, 50, 4, style, "  $1032 USD")
	emitStr(s, 50, 5, style, "   $864 USD")
	emitStr(s, 50, 6, style, "$12,000 USD")

	// My Dice
	emitStr(s, 3, 8, style, "My Dice: [3] [2] [6] [6] [1]")
	emitStr(s, 35, 8, style, "Pot: $15 USD")

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
