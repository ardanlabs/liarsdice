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
	v1 "github.com/ardanlabs/liarsdice/business/web/v1"
	"github.com/ardanlabs/liarsdice/business/web/v1/auth"
	"github.com/ardanlabs/liarsdice/business/web/v1/mid"
	"github.com/ardanlabs/liarsdice/foundation/logger"
	"github.com/ardanlabs/liarsdice/foundation/web"
	"github.com/ethereum/go-ethereum/common"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/websocket"
)

type handlers struct {
	converter      *currency.Converter
	bank           *bank.Bank
	log            *logger.Logger
	ws             websocket.Upgrader
	activeKID      string
	auth           *auth.Auth
	anteUSD        float64
	bankTimeout    time.Duration
	connectTimeout time.Duration
}

// connect is used to return a game token for API usage.
func (h *handlers) connect(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	address, err := validateSignature(r, h.connectTimeout)
	if err != nil {
		return v1.NewTrustedError(err, http.StatusBadRequest)
	}

	token, err := generateToken(h.auth, h.activeKID, address)
	if err != nil {
		return v1.NewTrustedError(err, http.StatusBadRequest)
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

// events handles a web socket to provide events to a client.
func (h *handlers) events(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	// Need this to handle CORS on the websocket.
	h.ws.CheckOrigin = func(r *http.Request) bool { return true }

	// This upgrades the HTTP connection to a websocket connection.
	c, err := h.ws.Upgrade(w, r, nil)
	if err != nil {
		return err
	}

	h.log.Info(ctx, "websocket open", "path", "/v1/game/events")

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
	subjectID := mid.GetSubject(ctx).String()
	ch := evts.acquire(subjectID)
	defer func() {
		evts.release(subjectID)
		h.log.Info(ctx, "evts.release", "account", subjectID)
	}()

	h.log.Info(ctx, "evts.acquire", "account", subjectID)

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
			h.log.Info(ctx, "*********> socket read", "path", "/v1/game/events", "message", message, "p", string(p))
		}
	}()

	defer func() {
		wg.Wait()
		h.log.Info(ctx, "websocket closed", "path", "/v1/game/events")
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
				h.log.Info(ctx, "websocket write", "path", "/v1/game/events", "ERROR", err)
				return nil
			}

			h.log.Info(ctx, "evts.send", "msg", msg)

		case <-pingSend.C:
			if err := c.WriteMessage(websocket.PingMessage, []byte("ping")); err != nil {
				h.log.Info(ctx, "websocket ping", "path", "/v1/game/events", "ERROR", err)
				return nil
			}
		}
	}
}

// configuration returns the basic configuration the front end needs to use.
func (h *handlers) configuration(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	info := struct {
		Network    string         `json:"network"`
		ChainID    int            `json:"chainId"`
		ContractID common.Address `json:"contractId"`
	}{
		Network:    h.bank.Client().Network(),
		ChainID:    h.bank.Client().ChainID(),
		ContractID: h.bank.ContractID(),
	}

	return web.Respond(ctx, w, info, http.StatusOK)
}

// usd2Wei converts the us dollar amount to wei based on the game engine's
// conversion rate.
func (h *handlers) usd2Wei(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	usd, err := strconv.ParseFloat(web.Param(ctx, "usd"), 64)
	if err != nil {
		return v1.NewTrustedError(fmt.Errorf("converting usd: %s", err), http.StatusBadRequest)
	}

	wei := h.converter.USD2Wei(big.NewFloat(usd))

	data := struct {
		USD float64  `json:"usd"`
		WEI *big.Int `json:"wei"`
	}{
		USD: usd,
		WEI: wei,
	}

	return web.Respond(ctx, w, data, http.StatusOK)
}

// tables returns current set of existing tables.
func (h *handlers) tables(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	info := struct {
		GameIDs []string `json:"gameIDs"`
	}{
		GameIDs: game.Tables.Active(),
	}

	return web.Respond(ctx, w, info, http.StatusOK)
}

// status will return information about the game.
func (h *handlers) status(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	claims := mid.GetClaims(ctx)
	address := common.HexToAddress(claims.Subject)

	gameID := web.Param(ctx, "id")

	g, err := h.getGame(gameID)
	if err != nil {
		resp := Status{
			Status:  "nogame",
			AnteUSD: h.anteUSD,
		}
		return web.Respond(ctx, w, resp, http.StatusOK)
	}

	status := g.Info(ctx)

	var cups []Cup
	for _, accountID := range status.ExistingPlayers {
		cup := status.Cups[accountID]

		// Don't share the dice information for other players.
		dice := []int{0, 0, 0, 0, 0}
		if accountID == address {
			dice = cup.Dice
		}
		cups = append(cups, Cup{Player: cup.Player, Dice: dice, Outs: cup.Outs})
	}

	var bets []Bet
	for _, bet := range status.Bets {
		bets = append(bets, Bet{Player: bet.Player, Number: bet.Number, Suit: bet.Suit})
	}

	resp := Status{
		GameID:          status.GameID,
		Status:          status.Status,
		AnteUSD:         h.anteUSD,
		PlayerLastOut:   status.PlayerLastOut,
		PlayerLastWin:   status.PlayerLastWin,
		PlayerTurn:      status.PlayerTurn,
		Round:           status.Round,
		Cups:            cups,
		ExistingPlayers: status.ExistingPlayers,
		Bets:            bets,
		Balances:        status.Balances,
	}

	return web.Respond(ctx, w, resp, http.StatusOK)
}

// newGame creates a new game if there is no game or the status of the current game
// is GameOver.
func (h *handlers) newGame(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	g, err := game.New(ctx, h.log, h.converter, h.bank, mid.GetSubject(ctx), h.anteUSD)
	if err != nil {
		return v1.NewTrustedError(fmt.Errorf("unable to create game: %w", err), http.StatusBadRequest)
	}

	subjectID := mid.GetSubject(ctx).String()

	h.log.Info(ctx, "evts.addPlayerToGame", "gameID", g.ID(), "account", subjectID)
	if err := evts.addPlayerToGame(g.ID(), subjectID); err != nil {
		h.log.Info(ctx, "evts.addPlayerToGame", "ERROR", err, "account", subjectID)
	}

	ctx = web.SetParam(ctx, "id", g.ID())

	return h.status(ctx, w, r)
}

// join adds the given player to the game.
func (h *handlers) join(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	g, err := h.getGame(web.Param(ctx, "id"))
	if err != nil {
		return err
	}

	subjectID := mid.GetSubject(ctx)

	if err := g.AddAccount(ctx, subjectID); err != nil {
		return v1.NewTrustedError(err, http.StatusBadRequest)
	}

	h.log.Info(ctx, "evts.addPlayerToGame", "gameID", g.ID(), "account", subjectID)
	if err := evts.addPlayerToGame(g.ID(), subjectID.String()); err != nil {
		h.log.Info(ctx, "evts.addPlayerToGame", "ERROR", err, "account", subjectID)
	}

	evts.send(g.ID(), fmt.Sprintf(`{"type":"join","address":%q}`, subjectID))

	return h.status(ctx, w, r)
}

// startGame changes the status of the game so players can begin to play.
func (h *handlers) startGame(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	g, err := h.getGame(web.Param(ctx, "id"))
	if err != nil {
		return err
	}

	if err := g.StartGame(ctx); err != nil {
		return v1.NewTrustedError(err, http.StatusBadRequest)
	}

	evts.send(g.ID(), fmt.Sprintf(`{"type":"start","address":%q}`, mid.GetSubject(ctx)))

	return h.status(ctx, w, r)
}

// rollDice will roll 5 dice for the given player and game.
func (h *handlers) rollDice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	g, err := h.getGame(web.Param(ctx, "id"))
	if err != nil {
		return err
	}

	if err := g.RollDice(ctx, mid.GetSubject(ctx)); err != nil {
		return v1.NewTrustedError(err, http.StatusBadRequest)
	}

	evts.send(g.ID(), fmt.Sprintf(`{"type":"rolldice","address":%q}`, mid.GetSubject(ctx)))

	return h.status(ctx, w, r)
}

// bet processes a bet made by a player in a game.
func (h *handlers) bet(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	g, err := h.getGame(web.Param(ctx, "id"))
	if err != nil {
		return err
	}

	number, err := strconv.Atoi(web.Param(ctx, "number"))
	if err != nil {
		return v1.NewTrustedError(fmt.Errorf("converting number: %s", err), http.StatusBadRequest)
	}

	suit, err := strconv.Atoi(web.Param(ctx, "suit"))
	if err != nil {
		return v1.NewTrustedError(fmt.Errorf("converting suit: %s", err), http.StatusBadRequest)
	}

	address := mid.GetSubject(ctx)

	if err := g.Bet(ctx, address, number, suit); err != nil {
		return v1.NewTrustedError(err, http.StatusBadRequest)
	}

	evts.send(g.ID(), fmt.Sprintf(`{"type":"bet","address":%q,"index":%d}`, address, g.Info(ctx).Cups[address].OrderIdx))

	return h.status(ctx, w, r)
}

// callLiar processes the claims and defines a winner and a loser for the round.
func (h *handlers) callLiar(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	g, err := h.getGame(web.Param(ctx, "id"))
	if err != nil {
		return err
	}

	if _, _, err := g.CallLiar(ctx, mid.GetSubject(ctx)); err != nil {
		return v1.NewTrustedError(err, http.StatusBadRequest)
	}

	if _, err := g.NextRound(ctx); err != nil {
		return v1.NewTrustedError(err, http.StatusBadRequest)
	}

	evts.send(g.ID(), fmt.Sprintf(`{"type":"callliar","address":%q}`, mid.GetSubject(ctx)))

	return h.status(ctx, w, r)
}

// reconcile calls the smart contract reconcile method.
func (h *handlers) reconcile(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	g, err := h.getGame(web.Param(ctx, "id"))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, h.bankTimeout)
	defer cancel()

	if _, _, err := g.Reconcile(ctx); err != nil {
		return v1.NewTrustedError(err, http.StatusInternalServerError)
	}

	evts.send(g.ID(), fmt.Sprintf(`{"type":"reconcile","address":%q}`, mid.GetSubject(ctx)))

	evts.removePlayersFromGame(g.ID())

	return h.status(ctx, w, r)
}

// balance returns the player balance from the smart contract.
func (h *handlers) balance(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ctx, cancel := context.WithTimeout(ctx, h.bankTimeout)
	defer cancel()

	balanceGWei, err := h.bank.AccountBalance(ctx, mid.GetSubject(ctx))
	if err != nil {
		return v1.NewTrustedError(err, http.StatusInternalServerError)
	}

	resp := struct {
		Balance string `json:"balance"`
	}{
		Balance: h.converter.GWei2USD(balanceGWei),
	}

	return web.Respond(ctx, w, resp, http.StatusOK)
}

// nextTurn changes the account that will make the next move.
func (h *handlers) nextTurn(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	g, err := h.getGame(web.Param(ctx, "id"))
	if err != nil {
		return err
	}

	if err := g.NextTurn(ctx); err != nil {
		return v1.NewTrustedError(err, http.StatusBadRequest)
	}

	evts.send(g.ID(), fmt.Sprintf(`{"type":"nextturn","address":%q}`, mid.GetSubject(ctx)))

	return h.status(ctx, w, r)
}

// updateOut replaces the current out amount of the player. This call is not
// part of the game flow, it is used to control when a player should be removed
// from the game.
func (h *handlers) updateOut(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	g, err := h.getGame(web.Param(ctx, "id"))
	if err != nil {
		return err
	}

	outs, err := strconv.Atoi(web.Param(ctx, "outs"))
	if err != nil {
		return v1.NewTrustedError(fmt.Errorf("converting outs: %s", err), http.StatusBadRequest)
	}

	address := mid.GetSubject(ctx)

	if err := g.ApplyOut(ctx, address, outs); err != nil {
		return v1.NewTrustedError(err, http.StatusBadRequest)
	}

	evts.send(g.ID(), fmt.Sprintf(`{"type":"outs","address":%q}`, address))

	return h.status(ctx, w, r)
}

// getGame safely returns a copy of the game pointer.
func (h *handlers) getGame(gameID string) (*game.Game, error) {
	g, err := game.Tables.Retrieve(gameID)
	if err != nil {
		return nil, v1.NewTrustedError(errors.New("no game exists"), http.StatusBadRequest)
	}

	return g, nil
}

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
