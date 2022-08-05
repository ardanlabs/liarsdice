package gamegrp

import (
	"context"
	"fmt"
	"net/http"

	v1Web "github.com/ardanlabs/liarsdice/business/web/v1"
	"github.com/ardanlabs/liarsdice/foundation/web"

	"github.com/google/uuid"
)

const (
	STATUSPLAYING = "playing"
	STATUSOPEN    = "open"
	NUMBERPLAYERS = 2
)

// Handlers manages the set of user endpoints.
type Handlers struct {
	Table Table
}

// New will return a new game with the player that made the request.
func (h Handlers) New(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var player Player
	if err := web.Decode(r, &player); err != nil {
		return fmt.Errorf("unable to decode payload: %w", err)
	}

	id := uuid.NewString()
	game := Game{
		ID:     id,
		Status: STATUSOPEN,
		Players: []Player{
			player,
		},
	}

	h.Table.Games[id] = game

	return web.Respond(ctx, w, game, http.StatusOK)
}

// List returns a list of all games.
func (h Handlers) List(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return web.Respond(ctx, w, h.Table.Games, http.StatusOK)
}

// Status will return information about the game.
func (h Handlers) Status(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	uuid := web.Param(r, "uuid")
	if uuid == "" {
		return v1Web.NewRequestError(fmt.Errorf("empty uuid"), http.StatusBadRequest)
	}

	game, ok := h.Table.Games[uuid]
	if !ok {
		return v1Web.NewRequestError(fmt.Errorf("invalid game uuid [%s]", uuid), http.StatusBadRequest)
	}

	return web.Respond(ctx, w, game, http.StatusOK)
}

// Join adds the given player to the given game.
func (h Handlers) Join(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	p := struct {
		ID     string `json:"id"`
		Wallet string `json:"wallet"`
	}{}

	if err := web.Decode(r, &p); err != nil {
		return fmt.Errorf("unable to decode payload: %w", err)
	}

	game, ok := h.Table.Games[p.ID]

	if !ok {
		return v1Web.NewRequestError(fmt.Errorf("game [%s] does not exist", p.ID), http.StatusBadRequest)
	}

	players := game.Players
	for _, player := range players {
		if p.Wallet == player.Wallet {
			return v1Web.NewRequestError(fmt.Errorf("player [%s] is already in the game [%s]", p.Wallet, p.ID), http.StatusBadRequest)
		}
	}

	player := Player{
		Wallet: p.Wallet,
	}

	game.Players = append(game.Players, player)

	h.Table.Games[p.ID] = game

	return web.Respond(ctx, w, game, http.StatusOK)
}

// Start checks if the game exists and can be started.
func (h Handlers) Start(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	uuid := web.Param(r, "uuid")
	if uuid == "" {
		return v1Web.NewRequestError(fmt.Errorf("empty uuid"), http.StatusBadRequest)
	}

	game, ok := h.Table.Games[uuid]
	if !ok {
		return v1Web.NewRequestError(fmt.Errorf("invalid game uuid [%s]", uuid), http.StatusBadRequest)
	}

	if game.Status != STATUSOPEN {
		return v1Web.NewRequestError(fmt.Errorf("game [%s] cannot be started", uuid), http.StatusBadRequest)
	}

	if len(game.Players) < NUMBERPLAYERS {
		return v1Web.NewRequestError(fmt.Errorf("not enough players to start game [%s]", uuid), http.StatusBadRequest)
	}

	game.Status = STATUSPLAYING

	h.Table.Games[uuid] = game

	return web.Respond(ctx, w, game, http.StatusOK)
}
