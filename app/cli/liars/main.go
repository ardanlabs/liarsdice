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

	board.Init()
	board.SetAnte(5.0)

	board.AddPlayer("0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd", "$1,000.00")
	board.AddPlayer("0x8e113078adf6888b7ba84967f299f29aece24c55", "$235.65")
	board.AddPlayer("0x0070742ff6003c3e809e78d524f0fe5dcc5ba7f7", "$12,765.44")

	board.ActivePlayer("0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd")

	if err := board.SetDice([]int{1, 3, 2, 5, 5}); err != nil {
		return err
	}

	<-board.StartEventLoop()
	return nil
}
