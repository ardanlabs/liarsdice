package gamegrp

import (
	"context"
	"fmt"
	"math/big"
	"net/http"

	v1Web "github.com/ardanlabs/liarsdice/business/web/v1"

	"github.com/ardanlabs/liarsdice/business/core/game"
	"github.com/ardanlabs/liarsdice/foundation/web"
)

const (
	STATUSPLAYING = "playing"
	STATUSOPEN    = "open"
	NUMBERPLAYERS = 2
)

// Handlers manages the set of user endpoints.
type Handlers struct {
	Game *game.Game
}

// Status will return information about the game.
func (h Handlers) Status(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return web.Respond(ctx, w, h.Game, http.StatusOK)
}

// Join adds the given player to the game.
func (h Handlers) Join(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var p struct {
		Wallet string `json:"wallet"`
	}

	if err := web.Decode(r, &p); err != nil {
		return fmt.Errorf("unable to decode payload: %w", err)
	}

	if err := h.Game.AddPlayer(p.Wallet); err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	return web.Respond(ctx, w, h.Game, http.StatusOK)
}

// Start starts the game.
func (h Handlers) Start(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if err := h.Game.StartGame(); err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	return web.Respond(ctx, w, h.Game, http.StatusOK)
}

// RollDice will roll 5 dice for the given player and game.
func (h Handlers) RollDice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	wallet := web.Param(r, "wallet")

	if err := h.Game.RollDice(wallet); err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	return web.Respond(ctx, w, h.Game, http.StatusOK)
}

// Claim processes a claim made by a player in a game.
func (h Handlers) Claim(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	wallet := web.Param(r, "wallet")

	var claim game.Claim
	if err := web.Decode(r, &claim); err != nil {
		return fmt.Errorf("unable to decode payload: %w", err)
	}

	if err := h.Game.Claim(wallet, claim); err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	return web.Respond(ctx, w, h.Game, http.StatusOK)
}

func (h Handlers) CallLiar(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	wallet := web.Param(r, "wallet")

	winner, loser, err := h.Game.CallLiar(wallet)
	if err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	resp := struct {
		Winner string
		Loser  string
	}{
		Winner: winner,
		Loser:  loser,
	}

	return web.Respond(ctx, w, resp, http.StatusOK)
}

func (h Handlers) Balance(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	wallet := web.Param(r, "wallet")

	balance, err := h.Game.PlayerBalance(ctx, wallet)
	if err != nil {
		return v1Web.NewRequestError(err, http.StatusInternalServerError)
	}

	resp := struct {
		Balance *big.Int `json:"balance"`
	}{
		Balance: balance,
	}

	return web.Respond(ctx, w, resp, http.StatusOK)
}
