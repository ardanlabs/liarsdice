package gamegrp

import (
	"context"
	"fmt"
	"math/rand"
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
		ID:      id,
		Status:  STATUSOPEN,
		Players: []Player{player},
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

// RollDices will roll 5 dices for the given player and game.
func (h Handlers) RollDices(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	uuid := web.Param(r, "uuid")
	if uuid == "" {
		return v1Web.NewRequestError(fmt.Errorf("empty uuid"), http.StatusBadRequest)
	}

	wallet := web.Param(r, "wallet")
	if wallet == "" {
		return v1Web.NewRequestError(fmt.Errorf("empty wallet address"), http.StatusBadRequest)
	}

	game, ok := h.Table.Games[uuid]
	if !ok {
		return v1Web.NewRequestError(fmt.Errorf("invalid game uuid [%s]", uuid), http.StatusBadRequest)
	}

	// The game should be in the playing state.
	if game.Status != STATUSPLAYING {
		return v1Web.NewRequestError(fmt.Errorf("game [%s] is not started", uuid), http.StatusBadRequest)
	}

	for i := range game.Players {
		if game.Players[i].Wallet == wallet {

			dice := make([]int, 5)
			for i := range dice {
				dice[i] = rand.Intn(6) + 1
			}

			game.Players[i].Dices = dice
			break
		}
	}

	return web.Respond(ctx, w, game, http.StatusOK)
}

// Claim processes a claim made by a player in a game.
func (h Handlers) Claim(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	uuid := web.Param(r, "uuid")
	if uuid == "" {
		return v1Web.NewRequestError(fmt.Errorf("empty uuid"), http.StatusBadRequest)
	}

	wallet := web.Param(r, "wallet")
	if wallet == "" {
		return v1Web.NewRequestError(fmt.Errorf("empty wallet address"), http.StatusBadRequest)
	}

	game := h.Table.Games[uuid]

	if game.Players[game.CurrentPlayer].Wallet != wallet {
		return v1Web.NewRequestError(fmt.Errorf("wrong player"), http.StatusBadRequest)
	}

	var claim Claim
	if err := web.Decode(r, &claim); err != nil {
		return fmt.Errorf("unable to decode payload: %w", err)
	}

	// Validate this player have rolled the dices,
	if game.Players[game.CurrentPlayer].Dices == nil {
		return v1Web.NewRequestError(fmt.Errorf("player [%s] didn't roll dices yet", wallet), http.StatusBadRequest)
	}

	// If this is not the first claim, validate the claim.
	if len(game.Claims) != 0 {
		lastClaim := game.Claims[len(game.Claims)-1]

		if claim.Number < lastClaim.Number {
			return v1Web.NewRequestError(fmt.Errorf("claim number must be greater or equal to the last claim"), http.StatusBadRequest)
		}

		if claim.Number == lastClaim.Number && claim.Suite <= lastClaim.Suite {
			return v1Web.NewRequestError(fmt.Errorf("claim suite must be greater that the last claim"), http.StatusBadRequest)
		}
	}

	claim.Wallet = wallet

	// Add the claim to the set of claims for this round.
	game.Claims = append(game.Claims, claim)

	// Increment the next player index.
	game.CurrentPlayer++
	game.CurrentPlayer = game.CurrentPlayer % len(game.Players)

	h.Table.Games[uuid] = game

	return web.Respond(ctx, w, game, http.StatusOK)
}

func (h Handlers) CallLiar(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	uuid := web.Param(r, "uuid")
	if uuid == "" {
		return v1Web.NewRequestError(fmt.Errorf("empty uuid"), http.StatusBadRequest)
	}

	wallet := web.Param(r, "wallet")
	if wallet == "" {
		return v1Web.NewRequestError(fmt.Errorf("empty wallet address"), http.StatusBadRequest)
	}

	game := h.Table.Games[uuid]

	if game.Players[game.CurrentPlayer].Wallet != wallet {
		return v1Web.NewRequestError(fmt.Errorf("wrong player"), http.StatusBadRequest)
	}

	// NOTE: do we need a ROUNDOVER status?

	dice := make([]int, 7) // The position 0 of the list will never be used.
	for _, player := range game.Players {
		for _, suite := range player.Dices {
			dice[suite]++
		}
	}

	lastClaim := game.Claims[len(game.Claims)-1]

	if dice[lastClaim.Suite] < lastClaim.Number {
		// the player that made the last claim, gets an `out``.
		var loser string
		for i := range game.Players {
			if game.Players[i].Wallet == lastClaim.Wallet {
				game.Players[i].Outs++
				loser = game.Players[i].Wallet
				break
			}
		}

		r := struct {
			ID     string
			Winner string
			Loser  string
		}{
			ID:     uuid,
			Winner: wallet,
			Loser:  loser,
		}

		return web.Respond(ctx, w, r, http.StatusOK)
	}

	resp := struct {
		ID     string
		Winner string
		Loser  string
	}{
		ID:     uuid,
		Winner: lastClaim.Wallet,
		Loser:  wallet,
	}

	return web.Respond(ctx, w, resp, http.StatusOK)
}
