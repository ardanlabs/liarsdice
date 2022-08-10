package game

import (
	"testing"
)

func TestNextCup(t *testing.T) {
	g := New(nil)

	g.AddAccount("a")
	g.AddAccount("b")

	// Playber B
	g.nextCup()
	if g.currentCup != 1 {
		t.Fatalf("expecting 1; got %d", g.currentCup)
	}
	if g.CupsOrder[g.currentCup] != "b" {
		t.Fatalf("expecting b; got %s", g.CupsOrder[g.currentCup])
	}

	// Playber A
	g.nextCup()
	if g.currentCup != 0 {
		t.Fatalf("expecting 0; got %d", g.currentCup)
	}
	if g.CupsOrder[g.currentCup] != "a" {
		t.Fatalf("expecting a; got %s", g.CupsOrder[g.currentCup])
	}

	// Playber B
	g.nextCup()
	if g.currentCup != 1 {
		t.Fatalf("expecting 1; got %d", g.currentCup)
	}
	if g.CupsOrder[g.currentCup] != "b" {
		t.Fatalf("expecting b; got %s", g.CupsOrder[g.currentCup])
	}
}

func TestNextCupRemovingPlayers(t *testing.T) {
	g := New(nil)

	g.AddAccount("a")
	g.AddAccount("b")
	g.AddAccount("c")
	g.AddAccount("d")

	g.RemoveAccount("b")
	g.RemoveAccount("d")

	g.nextCup()
	expectedIndex := 2
	expectedValue := "c"
	if g.currentCup != expectedIndex {
		t.Fatalf("expecting %d; got %d", expectedIndex, g.currentCup)
	}
	if g.CupsOrder[g.currentCup] != expectedValue {
		t.Fatalf("expecting %s; got %s", expectedValue, g.CupsOrder[g.currentCup])
	}
}

func TestNextCupInfiniteLoop(t *testing.T) {
	g := New(nil)

	g.AddAccount("a")

	g.RemoveAccount("a")

	err := g.nextCup()
	if err == nil {
		t.Fatal("expecting error; got nil")
	}
}
