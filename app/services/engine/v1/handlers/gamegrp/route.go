package gamegrp

import (
	"net/http"
	"time"

	"github.com/ardanlabs/ethereum/currency"
	"github.com/ardanlabs/liarsdice/business/core/bank"
	"github.com/ardanlabs/liarsdice/business/web/v1/auth"
	"github.com/ardanlabs/liarsdice/business/web/v1/mid"
	"github.com/ardanlabs/liarsdice/foundation/events"
	"github.com/ardanlabs/liarsdice/foundation/logger"
	"github.com/ardanlabs/liarsdice/foundation/web"
	"github.com/gorilla/websocket"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log            *logger.Logger
	Auth           *auth.Auth
	Converter      *currency.Converter
	Bank           *bank.Bank
	Evts           *events.Events
	AnteUSD        float64
	ActiveKID      string
	BankTimeout    time.Duration
	ConnectTimeout time.Duration
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	hdl := handlers{
		converter:      cfg.Converter,
		bank:           cfg.Bank,
		log:            cfg.Log,
		evts:           cfg.Evts,
		ws:             websocket.Upgrader{},
		auth:           cfg.Auth,
		activeKID:      cfg.ActiveKID,
		anteUSD:        cfg.AnteUSD,
		bankTimeout:    cfg.BankTimeout,
		connectTimeout: cfg.ConnectTimeout,
		games:          initGames(),
	}

	app.Handle(http.MethodPost, version, "/game/connect", hdl.connect)

	app.Handle(http.MethodGet, version, "/game/events", hdl.events)
	app.Handle(http.MethodGet, version, "/game/config", hdl.configuration)
	app.Handle(http.MethodGet, version, "/game/usd2wei/:usd", hdl.usd2Wei)
	app.Handle(http.MethodGet, version, "/game/new", hdl.newGame, mid.Authenticate(cfg.Auth))
	app.Handle(http.MethodGet, version, "/game/balance", hdl.balance, mid.Authenticate(cfg.Auth))
	app.Handle(http.MethodGet, version, "/game/tables", hdl.tables, mid.Authenticate(cfg.Auth))

	app.Handle(http.MethodGet, version, "/game/:id/status", hdl.status, mid.Authenticate(cfg.Auth))
	app.Handle(http.MethodGet, version, "/game/:id/join", hdl.join, mid.Authenticate(cfg.Auth))
	app.Handle(http.MethodGet, version, "/game/:id/start", hdl.startGame, mid.Authenticate(cfg.Auth))
	app.Handle(http.MethodGet, version, "/game/:id/rolldice", hdl.rollDice, mid.Authenticate(cfg.Auth))
	app.Handle(http.MethodGet, version, "/game/:id/bet/:number/:suit", hdl.bet, mid.Authenticate(cfg.Auth))
	app.Handle(http.MethodGet, version, "/game/:id/liar", hdl.callLiar, mid.Authenticate(cfg.Auth))
	app.Handle(http.MethodGet, version, "/game/:id/reconcile", hdl.reconcile, mid.Authenticate(cfg.Auth))

	// Timeout Situations with a player
	app.Handle(http.MethodGet, version, "/game/:id/next", hdl.nextTurn, mid.Authenticate(cfg.Auth))
	app.Handle(http.MethodGet, version, "/game/:id/out/:outs", hdl.updateOut, mid.Authenticate(cfg.Auth))
}
