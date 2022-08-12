package gamegrp

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"time"

	v1Web "github.com/ardanlabs/liarsdice/business/web/v1"
	"github.com/gorilla/websocket"

	"github.com/ardanlabs/liarsdice/business/core/game"
	"github.com/ardanlabs/liarsdice/foundation/events"
	"github.com/ardanlabs/liarsdice/foundation/web"
)

// Handlers manages the set of user endpoints.
type Handlers struct {
	Game *game.Game
	WS   websocket.Upgrader
	Evts *events.Events
}

// Events handles a web socket to provide events to a client.
func (h Handlers) Events(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	v, err := web.GetValues(ctx)
	if err != nil {
		return web.NewShutdownError("web value missing from context")
	}

	// Need this to handle CORS on the websocket.
	h.WS.CheckOrigin = func(r *http.Request) bool { return true }

	// This upgrades the HTTP connection to a websocket connection.
	c, err := h.WS.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	defer c.Close()

	// This provides a channel for receiving events from the blockchain.
	ch := h.Evts.Acquire(v.TraceID)
	defer h.Evts.Release(v.TraceID)

	// Starting a ticker to send a ping message over the websocket.
	ticker := time.NewTicker(time.Second)

	// Block waiting for events from the blockchain or ticker.
	for {
		select {
		case msg, wd := <-ch:

			// If the channel is closed, release the websocket.
			if !wd {
				return nil
			}

			if err := c.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
				return err
			}

		case <-ticker.C:
			if err := c.WriteMessage(websocket.PingMessage, []byte("ping")); err != nil {
				return nil
			}
		}
	}
}

// Start starts the game.
func (h Handlers) Start(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if err := h.Game.Start(); err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	resp := struct {
		Status        string   `json:"status"`
		Round         int      `json:"round"`
		CurrentPlayer string   `json:"current_player,omitempty"`
		PlayerOrder   []string `json:"player_order"`
	}{
		Status:        h.Game.Status,
		Round:         h.Game.Round,
		CurrentPlayer: h.Game.CurrentPlayer,
		PlayerOrder:   h.Game.CupsOrder,
	}

	h.Evts.Send("start")
	return web.Respond(ctx, w, resp, http.StatusOK)
}

// Status will return information about the game.
func (h Handlers) Status(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return web.Respond(ctx, w, gameToResponse(h.Game), http.StatusOK)
}

// Join adds the given player to the game.
func (h Handlers) Join(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var p struct {
		Wallet string `json:"wallet"`
	}

	if err := web.Decode(r, &p); err != nil {
		return fmt.Errorf("unable to decode payload: %w", err)
	}

	if err := h.Game.AddAccount(p.Wallet); err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	resp := struct {
		Status        string   `json:"status"`
		Round         int      `json:"round"`
		CurrentPlayer string   `json:"current_player,omitempty"`
		PlayerOrder   []string `json:"player_order"`
	}{
		Status:        h.Game.Status,
		Round:         h.Game.Round,
		CurrentPlayer: h.Game.CurrentPlayer,
		PlayerOrder:   h.Game.CupsOrder,
	}

	h.Evts.Send("join")
	return web.Respond(ctx, w, resp, http.StatusOK)
}

// RollDice will roll 5 dice for the given player and game.
func (h Handlers) RollDice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	wallet := web.Param(r, "wallet")

	if err := h.Game.RollDice(wallet); err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	// Find the player to show only this player's rolled dice.
	var player Player
	for _, cup := range h.Game.Cups {
		if cup.Account == wallet {
			player.Wallet = cup.Account
			player.Dice = cup.Dice
			player.Outs = cup.Outs
			break
		}
	}

	resp := struct {
		Player Player `json:"player"`
	}{
		Player: player,
	}

	h.Evts.Send("rolldice")
	return web.Respond(ctx, w, resp, http.StatusOK)
}

// Claim processes a claim made by a player in a game.
func (h Handlers) Claim(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	wallet := web.Param(r, "wallet")

	var claim Claim
	if err := web.Decode(r, &claim); err != nil {
		return fmt.Errorf("unable to decode payload: %w", err)
	}

	businessClaim := claimToBusiness(claim)

	if err := h.Game.Claim(wallet, businessClaim); err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	h.Evts.Send("claim")
	return web.Respond(ctx, w, gameToResponse(h.Game), http.StatusOK)
}

// CallLiar processes the claims and defines a winner and a loser for the round.
func (h Handlers) CallLiar(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	wallet := web.Param(r, "wallet")

	winner, loser, err := h.Game.CallLiar(wallet)
	if err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	resp := struct {
		Winner string `json:"winner"`
		Loser  string `json:"loser"`
	}{
		Winner: winner,
		Loser:  loser,
	}

	h.Evts.Send("callliar")
	return web.Respond(ctx, w, resp, http.StatusOK)
}

// NewRound starts a new round reseting the required data.
func (h Handlers) NewRound(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	playersLeft, err := h.Game.NewRound()
	if err != nil {
		return v1Web.NewRequestError(err, http.StatusInternalServerError)
	}

	resp := struct {
		PlayersLeft int `json:"players_left"`
	}{
		PlayersLeft: playersLeft,
	}

	h.Evts.Send("newround")
	return web.Respond(ctx, w, resp, http.StatusOK)
}

// Balance returns the player balance from the smart contract.
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

// UpdateOut replaces the current out amount of the player. This call is not
// part of the game flow, it is used to control when a player should be removed
// from the game.
func (h Handlers) UpdateOut(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var p struct {
		Wallet string `json:"wallet"`
		Outs   int    `json:"outs"`
	}

	if err := web.Decode(r, &p); err != nil {
		return fmt.Errorf("unable to decode payload: %w", err)
	}

	if err := h.Game.UpdateAccountOut(p.Wallet, p.Outs); err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	h.Evts.Send("updateout")
	return web.Respond(ctx, w, "OK", http.StatusOK)
}

//==============================================================================

func gameToResponse(game *game.Game) Game {
	g := Game{
		Status:        game.Status,
		Round:         game.Round,
		CurrentPlayer: game.CurrentPlayer,
		CupsOrder:     game.CupsOrder,
	}

	g.Players = playerToResponse(game.Cups, game.Claims)

	return g
}

func playerToResponse(cups map[string]game.Cup, claims []game.Claim) []Player {
	var playerList []Player

	for _, cup := range cups {
		p := Player{
			Wallet: cup.Account,
			Outs:   cup.Outs,
			Dice:   cup.Dice,
			Claim:  claimToResponse(cup.Account, claims),
		}
		playerList = append(playerList, p)
	}

	return playerList
}

func claimToResponse(account string, claims []game.Claim) Claim {
	for _, c := range claims {
		if c.Account == account {
			return Claim{
				Wallet: account,
				Number: c.Number,
				Suite:  c.Suite,
			}
		}
	}

	return Claim{}
}

func claimToBusiness(claim Claim) game.Claim {
	return game.Claim{
		Number: claim.Number,
		Suite:  claim.Suite,
	}
}
