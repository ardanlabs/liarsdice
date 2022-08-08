package game

import "github.com/ardanlabs/liarsdice/business/core/bank"

type Game struct {
	ID            string
	Status        string
	Players       []Player
	CurrentPlayer int
	Claims        []Claim
	Banker        bank.Banker
}

type Player struct {
	Wallet string
	Outs   uint8
	Dice   []int
}

type Claim struct {
	Wallet string
	Number int
	Suite  int
}
