package gamegrp

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/ardanlabs/liarsdice/business/web/mid"
	"github.com/google/uuid"
)

// These types exist for documentation purposes. The API will
// will accept a string.
type (
	gameID   uuid.UUID
	playerID string
)

// evts maintains the set player channels for sending messages over
// the players web socket.
var evts = newEvents()

// events maintains a mapping of unique id and channels so goroutines
// can register and receive events.
type events struct {
	players map[playerID]chan string
	games   map[gameID]map[playerID]struct{}
	mu      sync.RWMutex
}

func newEvents() *events {
	return &events{
		players: make(map[playerID]chan string),
		games:   make(map[gameID]map[playerID]struct{}),
	}
}

func (evt *events) acquire(pID string) chan string {
	evt.mu.Lock()
	defer evt.mu.Unlock()

	// Since a message will be dropped if the websocket receiver is
	// not ready to receive, this arbitrary buffer should give the receiver
	// enough time to not lose a message. Websocket send could take long.
	const messageBuffer = 100

	playID := playerID(pID)

	ch, exists := evt.players[playID]
	if !exists {
		ch = make(chan string, messageBuffer)
		evt.players[playID] = ch
	}

	return ch
}

func (evt *events) release(pID string) error {
	evt.mu.Lock()
	defer evt.mu.Unlock()

	playID := playerID(pID)

	ch, exists := evt.players[playID]
	if !exists {
		return fmt.Errorf("player id %q does not exist", pID)
	}

	delete(evt.players, playID)
	close(ch)

	return nil
}

func (evt *events) numberOfPlayers(gID uuid.UUID) (int, error) {
	evt.mu.RLock()
	defer evt.mu.RUnlock()

	gameID := gameID(gID)

	playerMap, exists := evt.games[gameID]
	if !exists {
		return 0, fmt.Errorf("game id %q does not exist", gID)
	}

	return len(playerMap), nil
}

func (evt *events) addPlayerToGame(gID uuid.UUID, pID string) error {
	evt.mu.Lock()
	defer evt.mu.Unlock()

	gameID := gameID(gID)
	playID := playerID(pID)

	if _, exists := evt.players[playID]; !exists {
		return fmt.Errorf("player id %q does not exist", pID)
	}

	playerMap, exists := evt.games[gameID]
	if !exists {
		playerMap = make(map[playerID]struct{})
		evt.games[gameID] = playerMap
	}

	playerMap[playID] = struct{}{}

	return nil
}

func (evt *events) removePlayersFromGame(gID uuid.UUID) error {
	evt.mu.Lock()
	defer evt.mu.Unlock()

	gameID := gameID(gID)

	playerMap, exists := evt.games[gameID]
	if !exists {
		return nil
	}

	for playID := range playerMap {
		delete(playerMap, playID)
	}
	delete(evt.games, gameID)

	return nil
}

// send signals a message to every registered channel for the specified
// game. Send will not block waiting for a receiver on any given channel.
func (evt *events) send(ctx context.Context, gID uuid.UUID, typ string, v ...any) {
	evt.mu.RLock()
	defer evt.mu.RUnlock()

	gameID := gameID(gID)

	playerMap, exists := evt.games[gameID]
	if !exists {
		return
	}

	var msg string
	switch v {
	case nil:
		msg = fmt.Sprintf(`{"type":"%s","address":"%s"}`, typ, mid.GetSubject(ctx))

	default:
		m := map[string]any{
			"type":    typ,
			"address": mid.GetSubject(ctx),
		}

		for i := 0; i < len(v); i = i + 2 {
			if vs, ok := v[i].(string); ok {
				m[vs] = v[i+1]
			}
		}

		data, err := json.Marshal(m)
		if err != nil {
			return
		}

		msg = string(data)
	}

	for playID := range playerMap {
		ch, exists := evt.players[playID]
		if !exists {
			continue
		}

		select {
		case ch <- msg:
		default:
		}
	}
}
