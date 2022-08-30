// Package gamegrp provides the handlers for game play.
package gamegrp

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/ardanlabs/liarsdice/business/web/auth"
	v1Web "github.com/ardanlabs/liarsdice/business/web/v1"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"

	"github.com/ardanlabs/liarsdice/business/core/bank"
	"github.com/ardanlabs/liarsdice/business/core/game"
	"github.com/ardanlabs/liarsdice/foundation/events"
	"github.com/ardanlabs/liarsdice/foundation/smart/contract"
	"github.com/ardanlabs/liarsdice/foundation/smart/currency"
	"github.com/ardanlabs/liarsdice/foundation/web"
)

// Handlers manages the set of user endpoints.
type Handlers struct {
	Converter   currency.Converter
	Bank        *bank.Bank
	Log         *zap.SugaredLogger
	WS          websocket.Upgrader
	Evts        *events.Events
	Auth        *auth.Auth
	AnteUSD     float64
	BankTimeout time.Duration

	game *game.Game
	mu   sync.RWMutex
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

// Configuration returns the basic configuration the front end needs to use.
func (h *Handlers) Configuration(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	info := struct {
		Network    string
		ChainID    int
		ContractID string
	}{
		Network:    h.Bank.Client().Network(),
		ChainID:    h.Bank.Client().ChainID(),
		ContractID: h.Bank.ContractID(),
	}

	return web.Respond(ctx, w, info, http.StatusOK)
}

// Status will return information about the game.
func (h *Handlers) Status(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	g, err := h.getGame()
	if err != nil {
		return err
	}

	status := g.Info()

	var cups []Cup
	for _, cup := range status.Cups {
		cups = append(cups, Cup{Account: cup.Account, Outs: cup.Outs})
	}

	var bets []Bet
	for _, bet := range status.Bets {
		bets = append(bets, Bet{Account: bet.Account, Number: bet.Number, Suite: bet.Suite})
	}

	resp := struct {
		Status        string   `json:"status"`
		AnteUSD       float64  `json:"ante_usd"`
		LastOutAcct   string   `json:"last_out"`
		LastWinAcct   string   `json:"last_win"`
		CurrentPlayer string   `json:"current_player"`
		CurrentCup    int      `json:"current_cup"`
		Round         int      `json:"round"`
		Cups          []Cup    `json:"cups"`
		CupsOrder     []string `json:"player_order"`
		Bets          []Bet    `json:"bets"`
	}{
		Status:        status.Status,
		AnteUSD:       h.AnteUSD,
		LastOutAcct:   status.LastOutAcct,
		LastWinAcct:   status.LastWinAcct,
		CurrentPlayer: status.CurrentPlayer,
		CurrentCup:    status.CurrentCup,
		Round:         status.Round,
		Cups:          cups,
		CupsOrder:     status.CupsOrder,
		Bets:          bets,
	}

	return web.Respond(ctx, w, resp, http.StatusOK)
}

// Connect is used to return a game token for API usage.
func (h *Handlers) Connect(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	address, err := validateSignature(r)
	if err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	token, err := generateToken(h.Auth, address)
	if err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	data := struct {
		Token string `json:"token"`
	}{
		Token: token,
	}

	return web.Respond(ctx, w, data, http.StatusOK)
}

// USD2Wei converts the us dollar amount to wei based on the game engine's
// conversion rate.
func (h *Handlers) USD2Wei(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	usd, err := strconv.ParseFloat(web.Param(r, "usd"), 64)
	if err != nil {
		return v1Web.NewRequestError(fmt.Errorf("converting usd: %s", err), http.StatusBadRequest)
	}

	wei := h.Converter.USD2Wei(big.NewFloat(usd))

	data := struct {
		USD float64  `json:"usd"`
		WEI *big.Int `json:"wei"`
	}{
		USD: usd,
		WEI: wei,
	}

	return web.Respond(ctx, w, data, http.StatusOK)
}

// NewGame creates a new game if there is no game or the status of the current game
// is GameOver.
func (h *Handlers) NewGame(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return v1Web.NewRequestError(auth.ErrForbidden, http.StatusForbidden)
	}
	address := claims.Subject

	if err := h.createGame(ctx, address); err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	h.Evts.Send(fmt.Sprintf(`{"type":"newgame","address":%q}`, address))

	return h.Status(ctx, w, r)
}

// Join adds the given player to the game.
func (h *Handlers) Join(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	g, err := h.getGame()
	if err != nil {
		return err
	}

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return v1Web.NewRequestError(auth.ErrForbidden, http.StatusForbidden)
	}
	address := claims.Subject

	if err := g.AddAccount(ctx, address); err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	h.Evts.Send(fmt.Sprintf(`{"type":"join","address":%q}`, address))

	return h.Status(ctx, w, r)
}

// StartGame changes the status of the game so players can begin to play.
func (h *Handlers) StartGame(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	g, err := h.getGame()
	if err != nil {
		return err
	}

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return v1Web.NewRequestError(auth.ErrForbidden, http.StatusForbidden)
	}
	address := claims.Subject

	if err := g.StartGame(address); err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	h.Evts.Send(fmt.Sprintf(`{"type":"start","address":%q}`, address))

	return h.Status(ctx, w, r)
}

// RollDice will roll 5 dice for the given player and game.
func (h *Handlers) RollDice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	g, err := h.getGame()
	if err != nil {
		return err
	}

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return v1Web.NewRequestError(auth.ErrForbidden, http.StatusForbidden)
	}
	address := claims.Subject

	if err := g.RollDice(address); err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	status := g.Info()
	cup, exists := status.Cups[address]
	if !exists {
		return v1Web.NewRequestError(errors.New("address not found"), http.StatusBadRequest)
	}

	h.Evts.Send(fmt.Sprintf(`{"type":"rolldice","address":%q}`, address))

	resp := struct {
		Dice []int `json:"dice"`
	}{
		Dice: cup.Dice,
	}

	return web.Respond(ctx, w, resp, http.StatusOK)
}

// Bet processes a bet made by a player in a game.
func (h *Handlers) Bet(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	g, err := h.getGame()
	if err != nil {
		return err
	}

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return v1Web.NewRequestError(auth.ErrForbidden, http.StatusForbidden)
	}
	address := claims.Subject

	number, err := strconv.Atoi(web.Param(r, "number"))
	if err != nil {
		return v1Web.NewRequestError(fmt.Errorf("converting number: %s", err), http.StatusBadRequest)
	}

	suite, err := strconv.Atoi(web.Param(r, "suite"))
	if err != nil {
		return v1Web.NewRequestError(fmt.Errorf("converting suite: %s", err), http.StatusBadRequest)
	}

	if err := g.Bet(address, number, suite); err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	h.Evts.Send(fmt.Sprintf(`{"type":"bet","address":%q}`, address))

	return h.Status(ctx, w, r)
}

// CallLiar processes the claims and defines a winner and a loser for the round.
func (h *Handlers) CallLiar(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	g, err := h.getGame()
	if err != nil {
		return err
	}

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return v1Web.NewRequestError(auth.ErrForbidden, http.StatusForbidden)
	}
	address := claims.Subject

	if _, _, err := g.CallLiar(address); err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	if _, err := g.NextRound(); err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	h.Evts.Send(fmt.Sprintf(`{"type":"callliar","address":%q}`, address))

	return h.Status(ctx, w, r)
}

// Reconcile calls the smart contract reconcile method.
func (h *Handlers) Reconcile(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	g, err := h.getGame()
	if err != nil {
		return err
	}

	v, err := web.GetValues(ctx)
	if err != nil {
		return web.NewShutdownError("missing web context values")
	}

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return v1Web.NewRequestError(auth.ErrForbidden, http.StatusForbidden)
	}
	address := claims.Subject

	ctx, cancel := context.WithTimeout(ctx, h.BankTimeout)
	defer cancel()

	tx, receipt, err := g.Reconcile(ctx, address)
	if err != nil {
		return v1Web.NewRequestError(err, http.StatusInternalServerError)
	}

	h.Log.Infow("reconcile results", "traceid", v.TraceID, "tx", h.Converter.CalculateTransactionDetails(tx), "receipt", h.Converter.CalculateReceiptDetails(receipt, tx.GasPrice()))

	h.Evts.Send(fmt.Sprintf(`{"type":"reconcile","address":%q}`, address))

	return h.Status(ctx, w, r)
}

// Balance returns the player balance from the smart contract.
func (h *Handlers) Balance(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return v1Web.NewRequestError(auth.ErrForbidden, http.StatusForbidden)
	}
	address := claims.Subject

	ctx, cancel := context.WithTimeout(ctx, h.BankTimeout)
	defer cancel()

	balanceGWei, err := h.Bank.AccountBalance(ctx, address)
	if err != nil {
		return v1Web.NewRequestError(err, http.StatusInternalServerError)
	}

	resp := struct {
		Balance string `json:"balance"`
	}{
		Balance: h.Converter.GWei2USD(balanceGWei),
	}

	return web.Respond(ctx, w, resp, http.StatusOK)
}

// NextTurn changes the account that will make the next move.
func (h *Handlers) NextTurn(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	g, err := h.getGame()
	if err != nil {
		return err
	}

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return v1Web.NewRequestError(auth.ErrForbidden, http.StatusForbidden)
	}
	address := claims.Subject

	if err := g.NextTurn(address); err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	h.Evts.Send(fmt.Sprintf(`{"type":"nextturn","address":%q}`, address))

	return h.Status(ctx, w, r)
}

// UpdateOut replaces the current out amount of the player. This call is not
// part of the game flow, it is used to control when a player should be removed
// from the game.
func (h *Handlers) UpdateOut(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	g, err := h.getGame()
	if err != nil {
		return err
	}

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return v1Web.NewRequestError(auth.ErrForbidden, http.StatusForbidden)
	}
	address := claims.Subject

	outs, err := strconv.Atoi(web.Param(r, "outs"))
	if err != nil {
		return v1Web.NewRequestError(fmt.Errorf("converting outs: %s", err), http.StatusBadRequest)
	}

	if err := g.ApplyOut(address, outs); err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	h.Evts.Send(fmt.Sprintf(`{"type":"outs","address":%q}`, address))

	return h.Status(ctx, w, r)
}

// =============================================================================

// SetGame safely sets a game pointer.
func (h *Handlers) createGame(ctx context.Context, address string) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.game != nil {
		status := h.game.Info()
		if status.Status != game.StatusGameOver {
			return errors.New("game is currently being played")
		}
	}

	g, err := game.New(ctx, h.Converter, h.Bank, address, h.AnteUSD)
	if err != nil {
		return fmt.Errorf("unable to create game: %w", err)
	}

	h.game = g

	return nil
}

// GetGame safely returns a copy of the game pointer.
func (h *Handlers) getGame() (*game.Game, error) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if h.game == nil {
		return nil, v1Web.NewRequestError(errors.New("no game exists"), http.StatusBadRequest)
	}

	return h.game, nil
}

// =============================================================================

func validateSignature(r *http.Request) (string, error) {
	var dt struct {
		DateTime  string `json:"date_time"` // YYYYMMDDHHMMSS
		Signature string `json:"sig"`
	}

	if err := web.Decode(r, &dt); err != nil {
		return "", fmt.Errorf("unable to decode payload: %w", err)
	}

	data := struct {
		DateTime string `json:"date_time"`
	}{
		DateTime: dt.DateTime,
	}

	address, err := contract.FromAddress(data, dt.Signature)
	if err != nil {
		return "", fmt.Errorf("unable to extract address: %w", err)
	}

	return address, nil
}

func generateToken(a *auth.Auth, address string) (string, error) {
	claims := auth.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   address,
			Issuer:    "service project",
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		},
	}

	token, err := a.GenerateToken(claims)
	if err != nil {
		return "", fmt.Errorf("generating token: %w", err)
	}

	return token, nil
}
