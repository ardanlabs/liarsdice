package gamegrp

import (
	"fmt"
	"sync"
	"time"

	"github.com/ardanlabs/liarsdice/business/core/game"
)

type games struct {
	mp map[string]*game.Game
	mu sync.RWMutex
}

func initGames() *games {
	return &games{
		mp: make(map[string]*game.Game),
	}
}

func (g *games) add(key string, gm *game.Game) {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.mp[key] = gm

	// Let's find games that are older than an hour and
	// remove them from the cache.
	hour := time.Now().Add(time.Hour)
	for k, v := range g.mp {
		if v.CreatedDate().After(hour) {
			delete(g.mp, k)
		}
	}
}

func (g *games) delete(key string) {
	g.mu.Lock()
	defer g.mu.Unlock()

	delete(g.mp, key)
}

func (g *games) retrieve(key string) (*game.Game, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()

	game, ok := g.mp[key]
	if !ok {
		return nil, fmt.Errorf("key %q not found", key)
	}

	return game, nil
}

func (g *games) active() []string {
	g.mu.RLock()
	defer g.mu.RUnlock()

	var ids []string

	for k, v := range g.mp {
		switch v.Status() {
		case game.StatusPlaying, game.StatusNewGame, game.StatusRoundOver:
			ids = append(ids, k)
		}
	}

	return ids
}
