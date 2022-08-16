// Package gamegrp provides the handlers for game play.
package gamegrp

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"time"

	v1Web "github.com/ardanlabs/liarsdice/business/web/v1"
	"github.com/gorilla/websocket"

	"github.com/ardanlabs/liarsdice/business/core/game"
	"github.com/ardanlabs/liarsdice/foundation/events"
	"github.com/ardanlabs/liarsdice/foundation/web"
)

// Handlers manages the set of user endpoints.
type Handlers struct {
	Banker game.Banker
	WS     websocket.Upgrader
	Evts   *events.Events

	game *game.Game
}

// Events handles a web socket to provide events to a client.
func (h *Handlers) Events(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	v, err := web.GetValues(ctx)
	if err != nil {
		return v1Web.NewRequestError(errors.New("web value missing from context"), http.StatusBadRequest)
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

// Status will return information about the game.
func (h *Handlers) Status(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if h.game == nil {
		return v1Web.NewRequestError(errors.New("no game exists"), http.StatusBadRequest)
	}

	status := h.game.Info()

	var cups []Cup

	for _, cup := range status.Cups {
		cups = append(cups, Cup{Account: cup.Account, Outs: cup.Outs})
	}

	resp := struct {
		Status        string       `json:"status"`
		LastOutAcct   string       `json:"last_out"`
		LastWinAcct   string       `json:"last_win"`
		CurrentPlayer string       `json:"current_player"`
		CurrentCup    int          `json:"current_cup"`
		Round         int          `json:"round"`
		Cups          []Cup        `json:"cups"`
		CupsOrder     []string     `json:"player_order"`
		Claims        []game.Claim `json:"claims"`
	}{
		Status:        status.Status,
		LastOutAcct:   status.LastOutAcct,
		LastWinAcct:   status.LastWinAcct,
		CurrentPlayer: status.CurrentPlayer,
		CurrentCup:    status.CurrentCup,
		Round:         status.Round,
		Cups:          cups,
		CupsOrder:     status.CupsOrder,
		Claims:        status.Claims,
	}

	return web.Respond(ctx, w, resp, http.StatusOK)
}

// NewGame creates a new game if there is no game or the status of the current game
// is GameOver.
func (h *Handlers) NewGame(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if h.game != nil {
		status := h.game.Info()
		if status.Status != game.StatusGameOver {
			return v1Web.NewRequestError(errors.New("game is currently being played"), http.StatusBadRequest)
		}
	}

	ante, err := strconv.Atoi(web.Param(r, "ante"))
	if err != nil {
		return v1Web.NewRequestError(errors.New("invalid ante value"), http.StatusBadRequest)
	}

	h.game = game.New(h.Banker, ante)

	h.Evts.Send("newgame")

	return h.Status(ctx, w, r)
}

// Join adds the given player to the game.
func (h *Handlers) Join(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if h.game == nil {
		return v1Web.NewRequestError(errors.New("no game exists"), http.StatusBadRequest)
	}

	address := web.Param(r, "address")

	if err := h.game.AddAccount(ctx, address); err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	h.Evts.Send("join:" + address)

	return h.Status(ctx, w, r)
}

// Start creates a new game if there is no game or the status of the current game
// is GameOver.
func (h *Handlers) Start(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if h.game == nil {
		return v1Web.NewRequestError(errors.New("no game exists"), http.StatusBadRequest)
	}

	status := h.game.Info()
	if status.Status != game.StatusGameOver {
		return v1Web.NewRequestError(errors.New("current game status doesn't allow this call"), http.StatusBadRequest)
	}

	if err := h.game.StartPlay(); err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	h.Evts.Send("start")

	return h.Status(ctx, w, r)
}

// RollDice will roll 5 dice for the given player and game.
func (h *Handlers) RollDice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if h.game == nil {
		return v1Web.NewRequestError(errors.New("no game exists"), http.StatusBadRequest)
	}

	address := web.Param(r, "address")

	if err := h.game.RollDice(address); err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	status := h.game.Info()
	cup, exists := status.Cups[address]
	if !exists {
		return v1Web.NewRequestError(errors.New("address not found"), http.StatusBadRequest)
	}

	h.Evts.Send("rolldice:" + address)

	resp := struct {
		Dice []int `json:"dice"`
	}{
		Dice: cup.Dice,
	}

	return web.Respond(ctx, w, resp, http.StatusOK)
}

// Claim processes a claim made by a player in a game.
func (h *Handlers) Claim(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if h.game == nil {
		return v1Web.NewRequestError(errors.New("no game exists"), http.StatusBadRequest)
	}

	address := web.Param(r, "address")

	number, err := strconv.Atoi(web.Param(r, "number"))
	if err != nil {
		return v1Web.NewRequestError(fmt.Errorf("converting number: %s", err), http.StatusBadRequest)
	}

	suite, err := strconv.Atoi(web.Param(r, "suite"))
	if err != nil {
		return v1Web.NewRequestError(fmt.Errorf("converting suite: %s", err), http.StatusBadRequest)
	}

	if err := h.game.Claim(address, number, suite); err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	h.Evts.Send("claim")

	return h.Status(ctx, w, r)
}

// CallLiar processes the claims and defines a winner and a loser for the round.
func (h *Handlers) CallLiar(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if h.game == nil {
		return v1Web.NewRequestError(errors.New("no game exists"), http.StatusBadRequest)
	}

	address := web.Param(r, "address")

	winner, loser, err := h.game.CallLiar(address)
	if err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	h.Evts.Send("callliar")

	resp := struct {
		Winner string `json:"winner"`
		Loser  string `json:"loser"`
	}{
		Winner: winner,
		Loser:  loser,
	}

	return web.Respond(ctx, w, resp, http.StatusOK)
}

// NewRound starts a new round reseting the required data.
func (h *Handlers) NewRound(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if h.game == nil {
		return v1Web.NewRequestError(errors.New("no game exists"), http.StatusBadRequest)
	}

	playersLeft, err := h.game.NextRound()
	if err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	h.Evts.Send("newround")

	resp := struct {
		PlayersLeft int `json:"players_left"`
	}{
		PlayersLeft: playersLeft,
	}

	return web.Respond(ctx, w, resp, http.StatusOK)
}

// UpdateOut replaces the current out amount of the player. This call is not
// part of the game flow, it is used to control when a player should be removed
// from the game.
func (h *Handlers) UpdateOut(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if h.game == nil {
		return v1Web.NewRequestError(errors.New("no game exists"), http.StatusBadRequest)
	}

	address := web.Param(r, "address")

	outs, err := strconv.Atoi(web.Param(r, "outs"))
	if err != nil {
		return v1Web.NewRequestError(fmt.Errorf("converting outs: %s", err), http.StatusBadRequest)
	}

	if err := h.game.ApplyOut(address, outs); err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	h.Evts.Send("outs:" + address)

	return h.Status(ctx, w, r)
}

// Balance returns the player balance from the smart contract.
func (h *Handlers) Balance(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if h.game == nil {
		return v1Web.NewRequestError(errors.New("no game exists"), http.StatusBadRequest)
	}

	address := web.Param(r, "address")

	balance, err := h.game.PlayerBalance(ctx, address)
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
