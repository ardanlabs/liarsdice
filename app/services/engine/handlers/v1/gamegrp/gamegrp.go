// Package gamegrp provides the handlers for game play.
package gamegrp

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	"github.com/ardanlabs/liarsdice/business/core/bank"
	"github.com/ardanlabs/liarsdice/business/core/game"
	"github.com/ardanlabs/liarsdice/business/web/auth"
	v1Web "github.com/ardanlabs/liarsdice/business/web/v1"
	"github.com/ardanlabs/liarsdice/foundation/events"
	"github.com/ardanlabs/liarsdice/foundation/web"
	"github.com/ethereum/go-ethereum/common"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

// Handlers manages the set of user endpoints.
type Handlers struct {
	Converter      *currency.Converter
	Bank           *bank.Bank
	Log            *zap.SugaredLogger
	WS             websocket.Upgrader
	Evts           *events.Events
	ActiveKID      string
	Auth           *auth.Auth
	AnteUSD        float64
	BankTimeout    time.Duration
	ConnectTimeout time.Duration

	game *game.Game
	mu   sync.RWMutex
}

// Connect is used to return a game token for API usage.
func (h *Handlers) Connect(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	address, err := validateSignature(r, h.ConnectTimeout)
	if err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	token, err := generateToken(h.Auth, h.ActiveKID, address)
	if err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	data := struct {
		Token   string `json:"token"`
		Address string `json:"address"`
	}{
		Token:   token,
		Address: address,
	}

	return web.Respond(ctx, w, data, http.StatusOK)
}

// Events handles a web socket to provide events to a client.
func (h *Handlers) Events(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	v := web.GetValues(ctx)

	// Need this to handle CORS on the websocket.
	h.WS.CheckOrigin = func(r *http.Request) bool { return true }

	// This upgrades the HTTP connection to a websocket connection.
	c, err := h.WS.Upgrade(w, r, nil)
	if err != nil {
		return err
	}

	h.Log.Infow("websocket open", "path", "/v1/game/events", "traceid", v.TraceID)

	// Set the timeouts for the ping to identify if a web socket
	// connection is broken.
	pongWait := 15 * time.Second
	pingPeriod := (pongWait * 9) / 10

	c.SetReadDeadline(time.Now().Add(pongWait))

	// Setup the pong handler to log the receiving of a pong.
	f := func(appData string) error {
		c.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	}
	c.SetPongHandler(f)

	// This provides a channel for receiving events from the blockchain.
	ch := h.Evts.Acquire(v.TraceID)
	defer h.Evts.Release(v.TraceID)

	// Starting a ticker to send a ping message over the websocket.
	pingSend := time.NewTicker(pingPeriod)

	// Set up the ability to receive chat messages.
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		// This supports the ability to add a chat system and receive a client
		// message.
		for {
			message, p, err := c.ReadMessage()
			if err != nil {
				return
			}
			h.Log.Infow("*********> socket read", "path", "/v1/game/events", "message", message, "p", string(p))
		}
	}()

	defer func() {
		wg.Wait()
		h.Log.Infow("websocket closed", "path", "/v1/game/events", "traceid", v.TraceID)
	}()
	defer c.Close()

	// Send game engine events back to the connected client.
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

		case <-pingSend.C:
			if err := c.WriteMessage(websocket.PingMessage, []byte("ping")); err != nil {
				return nil
			}
		}
	}
}

// Configuration returns the basic configuration the front end needs to use.
func (h *Handlers) Configuration(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	info := struct {
		Network    string         `json:"network"`
		ChainID    int            `json:"chainId"`
		ContractID common.Address `json:"contractId"`
	}{
		Network:    h.Bank.Client().Network(),
		ChainID:    h.Bank.Client().ChainID(),
		ContractID: h.Bank.ContractID(),
	}

	return web.Respond(ctx, w, info, http.StatusOK)
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

// Status will return information about the game.
func (h *Handlers) Status(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	claims := auth.GetClaims(ctx)
	address := common.HexToAddress(claims.Subject)

	g, err := h.getGame()
	if err != nil {
		resp := Status{
			Status:  "nogame",
			AnteUSD: h.AnteUSD,
		}
		return web.Respond(ctx, w, resp, http.StatusOK)
	}

	status := g.Info(ctx)

	var cups []Cup
	for _, accountID := range status.CupsOrder {
		cup := status.Cups[accountID]

		// Don't share the dice information for other players.
		dice := []int{0, 0, 0, 0, 0}
		if accountID == address {
			dice = cup.Dice
		}
		cups = append(cups, Cup{AccountID: cup.AccountID, Dice: dice, LastBet: Bet(cup.LastBet), Outs: cup.Outs})
	}

	var bets []Bet
	for _, bet := range status.Bets {
		bets = append(bets, Bet{AccountID: bet.AccountID, Number: bet.Number, Suite: bet.Suite})
	}

	resp := Status{
		Status:        status.Status,
		AnteUSD:       h.AnteUSD,
		LastOutAcctID: status.LastOutAcctID,
		LastWinAcctID: status.LastWinAcctID,
		CurrentAcctID: status.CurrentAcctID,
		Round:         status.Round,
		Cups:          cups,
		CupsOrder:     status.CupsOrder,
		Bets:          bets,
		Balances:      status.Balances,
	}

	return web.Respond(ctx, w, resp, http.StatusOK)
}

// NewGame creates a new game if there is no game or the status of the current game
// is GameOver.
func (h *Handlers) NewGame(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	claims := auth.GetClaims(ctx)
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

	claims := auth.GetClaims(ctx)
	address := common.HexToAddress(claims.Subject)

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

	claims := auth.GetClaims(ctx)
	address := claims.Subject

	if err := g.StartGame(ctx); err != nil {
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

	claims := auth.GetClaims(ctx)
	address := common.HexToAddress(claims.Subject)

	if err := g.RollDice(ctx, address); err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	h.Evts.Send(fmt.Sprintf(`{"type":"rolldice","address":%q}`, address))

	return h.Status(ctx, w, r)
}

// Bet processes a bet made by a player in a game.
func (h *Handlers) Bet(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	g, err := h.getGame()
	if err != nil {
		return err
	}

	claims := auth.GetClaims(ctx)
	address := common.HexToAddress(claims.Subject)

	number, err := strconv.Atoi(web.Param(r, "number"))
	if err != nil {
		return v1Web.NewRequestError(fmt.Errorf("converting number: %s", err), http.StatusBadRequest)
	}

	suite, err := strconv.Atoi(web.Param(r, "suite"))
	if err != nil {
		return v1Web.NewRequestError(fmt.Errorf("converting suite: %s", err), http.StatusBadRequest)
	}

	if err := g.Bet(ctx, address, number, suite); err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	h.Evts.Send(fmt.Sprintf(`{"type":"bet","address":%q,"index":%d}`, address, g.Info(ctx).Cups[address].OrderIdx))

	return h.Status(ctx, w, r)
}

// CallLiar processes the claims and defines a winner and a loser for the round.
func (h *Handlers) CallLiar(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	g, err := h.getGame()
	if err != nil {
		return err
	}

	claims := auth.GetClaims(ctx)
	address := common.HexToAddress(claims.Subject)

	if _, _, err := g.CallLiar(ctx, address); err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	if _, err := g.NextRound(ctx); err != nil {
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

	claims := auth.GetClaims(ctx)
	address := common.HexToAddress(claims.Subject)

	ctx, cancel := context.WithTimeout(ctx, h.BankTimeout)
	defer cancel()

	if _, _, err := g.Reconcile(ctx, address); err != nil {
		return v1Web.NewRequestError(err, http.StatusInternalServerError)
	}

	h.Evts.Send(fmt.Sprintf(`{"type":"reconcile","address":%q}`, address))

	return h.Status(ctx, w, r)
}

// Balance returns the player balance from the smart contract.
func (h *Handlers) Balance(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	claims := auth.GetClaims(ctx)
	address := claims.Subject

	ctx, cancel := context.WithTimeout(ctx, h.BankTimeout)
	defer cancel()

	balanceGWei, err := h.Bank.AccountBalance(ctx, common.HexToAddress(address))
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

	claims := auth.GetClaims(ctx)
	address := common.HexToAddress(claims.Subject)

	if err := g.NextTurn(ctx); err != nil {
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

	claims := auth.GetClaims(ctx)
	address := common.HexToAddress(claims.Subject)

	outs, err := strconv.Atoi(web.Param(r, "outs"))
	if err != nil {
		return v1Web.NewRequestError(fmt.Errorf("converting outs: %s", err), http.StatusBadRequest)
	}

	if err := g.ApplyOut(ctx, address, outs); err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	h.Evts.Send(fmt.Sprintf(`{"type":"outs","address":%q}`, address))

	return h.Status(ctx, w, r)
}

// =============================================================================

// SetGame resets the existing game. At this time we let this happen at any
// time regardless of game state.
func (h *Handlers) createGame(ctx context.Context, address string) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	g, err := game.New(ctx, h.Log, h.Converter, h.Bank, common.HexToAddress(address), h.AnteUSD)
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

func validateSignature(r *http.Request, timeout time.Duration) (string, error) {
	var dt struct {
		Address   string `json:"address"`
		DateTime  string `json:"dateTime"` // YYYYMMDDHHMMSS
		Signature string `json:"sig"`
	}

	if err := web.Decode(r, &dt); err != nil {
		return "", fmt.Errorf("unable to decode payload: %w", err)
	}

	t, err := time.Parse("20060102150405", dt.DateTime)
	if err != nil {
		return "", fmt.Errorf("parse time: %w", err)
	}

	if d := time.Since(t); d > timeout {
		return "", fmt.Errorf("data is too old, %v", d.Seconds())
	}

	data := struct {
		Address  string `json:"address"`
		DateTime string `json:"dateTime"`
	}{
		Address:  dt.Address,
		DateTime: dt.DateTime,
	}

	address, err := ethereum.FromAddressAny(data, dt.Signature)
	if err != nil {
		return "", fmt.Errorf("unable to extract address: %w", err)
	}

	if !strings.EqualFold(strings.ToLower(address), strings.ToLower(data.Address)) {
		return "", fmt.Errorf("invalid address match, got[%s] exp[%s]", address, data.Address)
	}

	return address, nil
}

func generateToken(a *auth.Auth, kid string, address string) (string, error) {
	claims := auth.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   address,
			Issuer:    "liar's project",
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		},
	}

	token, err := a.GenerateToken(kid, claims)
	if err != nil {
		return "", fmt.Errorf("generating token: kid: %s: %w", kid, err)
	}

	return token, nil
}
