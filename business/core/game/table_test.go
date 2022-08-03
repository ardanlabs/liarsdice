package game

import (
	"testing"
)

func TestRemovePlayerFromGame(t *testing.T) {
	table := NewTable(1)

	playerA := Player{
		UserID:  "a",
		Address: "aaa",
		Dice:    nil,
	}

	playerB := Player{
		UserID:  "b",
		Address: "bbb",
		Dice:    nil,
	}
	table.AddPlayer(&playerA)
	table.AddPlayer(&playerB)

	err := table.StartGame()
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	tablePlayers := len(table.Players)
	gamePlayers := len(table.Game.Players)

	if tablePlayers != 2 && gamePlayers != 2 {
		t.Fatalf("expecting 2 players in the table and game; got table %d and game %d", tablePlayers, gamePlayers)
	}

	err = table.Game.RemovePlayer(playerA.UserID)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	// Count players again.
	tablePlayers = len(table.Players)
	gamePlayers = len(table.Game.Players)

	if tablePlayers != 1 && gamePlayers != 1 {
		t.Fatalf("expecting 1 player in the table and game; got table %d and game %d", tablePlayers, gamePlayers)
	}
}

func TestGameFlow(t *testing.T) {
	table := NewTable(10)

	// Player A rolls the dices.
	playerA := Player{
		UserID:  "a",
		Address: "aaa",
		Dice:    nil,
	}
	playerA.Dice = []int{3, 4, 3, 1, 1}

	// Player B rolls the dices.
	playerB := Player{
		UserID:  "b",
		Address: "bbb",
		Dice:    nil,
	}
	playerB.Dice = []int{6, 6, 2, 4, 5}

	// Player C rolls the dices.
	playerC := Player{
		UserID:  "c",
		Address: "ccc",
		Dice:    nil,
	}
	playerC.Dice = []int{3, 1, 3, 2, 1}

	// Add players to the table.
	table.AddPlayer(&playerA)
	table.AddPlayer(&playerB)
	table.AddPlayer(&playerC)

	// StartGame will also add the players to the Game.
	err := table.StartGame()
	if err != nil {
		t.Fatalf("StartGame: %s", err)
	}

	// Player A makes a claim.
	claimA1 := Claim{
		Player: &playerA,
		Number: 2,
		Suite:  1,
	}

	err = table.MakeClaim(claimA1)
	if err != nil {
		t.Fatalf("MakeClaim A: %s", err)
	}

	// Player B makes a claim.
	claimB1 := Claim{
		Player: &playerB,
		Number: 2,
		Suite:  4,
	}

	err = table.MakeClaim(claimB1)
	if err != nil {
		t.Fatalf("MakeClaim B: %s", err)
	}

	// Player C makes a claim.
	claimC1 := Claim{
		Player: &playerC,
		Number: 3,
		Suite:  3,
	}

	err = table.MakeClaim(claimC1)
	if err != nil {
		t.Fatalf("MakeClaim C: %s", err)
	}

	//==========================================================================
	// Player A calls Player C a liar.
	// Player C is the winner.
	winner, loser, err := table.CallLiar(&playerA)
	if err != nil {
		t.Fatalf("CallLiar: %s", err)
	}

	//==========================================================================
	// Some required validations before the next round.
	if winner != &playerC {
		t.Fatalf("expecting player %s to be the winner; got %s", playerC.UserID, winner.UserID)
	}

	if loser != &playerA {
		t.Fatalf("expecting player %s to be the loser; got %s", playerA.UserID, loser.UserID)
	}

	if table.Game.Outs[&playerA] != 1 {
		t.Fatalf("expecting loser to have 1 out; got %d", table.Game.Outs[&playerA])
	}

	nextPlayer := table.NextPlayer()
	if &playerA != nextPlayer {
		t.Fatalf("expecting '%s' to be the next player; got '%s'", playerA.UserID, nextPlayer.UserID)
	}

	//==========================================================================

	// NEXT ROUND

	//==========================================================================

	playersLeft, err := table.NewRound()
	if err != nil {
		t.Fatalf("NewRound: %s", err)
	}

	if playersLeft != 3 {
		t.Fatalf("expect 3 players left; got %d", playersLeft)
	}

	// Players new dice rolls.
	playerA.Dice = []int{5, 1, 5, 2, 2}
	claimA2 := Claim{
		Player: &playerA,
		Number: 2,
		Suite:  2,
	}
	table.MakeClaim(claimA2)

	//==========================================================================
	// Player B calls Player A a liar.
	// Player C should be the winner.
	winner, loser, err = table.CallLiar(&playerB)
	if err != nil {
		t.Fatalf("CallLiar: %s", err)
	}

	//==========================================================================
	// Some required validations before the next round.
	if winner != &playerA {
		t.Fatalf("expecting player %s to be the winner; got %s", playerA.UserID, winner.UserID)
	}

	if loser != &playerB {
		t.Fatalf("expecting player %s to be the loser; got %s", playerB.UserID, loser.UserID)
	}

	if table.Game.Outs[&playerB] != 1 {
		t.Fatalf("expecting loser to have 1 out; got %d", table.Game.Outs[&playerB])
	}

	nextPlayer = table.NextPlayer()
	if &playerB != nextPlayer {
		t.Fatalf("expecting '%s' to be the next player; got '%s'", playerB.UserID, nextPlayer.UserID)
	}
}
