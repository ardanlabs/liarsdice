package game

import (
	"fmt"
	"sync"
	"time"
)

// Tables maintains the set of games in the system.
var Tables = newTables()

// tables represent the current set of tables that actively exist. The state
// of these tables can be of any state. The Add API will remove tables that are
// older than an hour.
type tables struct {
	games map[string]*Game
	mu    sync.RWMutex
}

func newTables() *tables {
	return &tables{
		games: make(map[string]*Game),
	}
}

// Add inserts the specified game into the table management system.
func (t *tables) add(game *Game) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.games[game.id] = game

	// Let's find games that are older than an hour and
	// remove them from the cache.
	hour := time.Now().Add(time.Hour)
	for k, v := range t.games {
		if v.CreatedDate().After(hour) {
			delete(t.games, k)
		}
	}
}

// Retrieve returns the specified game from the table management system.
func (t *tables) Retrieve(key string) (*Game, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	game, ok := t.games[key]
	if !ok {
		return nil, fmt.Errorf("key %q not found", key)
	}

	return game, nil
}

// Active returns the IDs for all the active games in the system.
func (t *tables) Active() []string {
	t.mu.RLock()
	defer t.mu.RUnlock()

	var ids []string

	for k, v := range t.games {
		switch v.Status() {
		case StatusPlaying, StatusNewGame, StatusRoundOver:
			ids = append(ids, k)
		}
	}

	return ids
}
