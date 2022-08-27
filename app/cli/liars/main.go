package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
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
			case *tcell.EventResize:
				s.Sync()
			}
		}
	}()

	style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	s.SetStyle(style)
	s.Clear()

	makebox(s)

	<-quit

	return nil
}

func makebox(s tcell.Screen) {
	st := tcell.StyleDefault
	st = st.Background(tcell.ColorLightBlue)

	s.SetContent(1, 1, '=', []rune{'+'}, st)
	s.Show()
}
