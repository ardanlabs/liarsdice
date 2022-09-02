package main

import (
	"os"

	"github.com/ardanlabs/liarsdice/app/cli/liars/board"
)

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	board, err := board.New()
	if err != nil {
		return err
	}
	defer board.Shutdown()

	<-board.Run()

	return nil
}
