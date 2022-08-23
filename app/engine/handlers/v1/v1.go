// Package v1 contains the full set of handler functions and routes
// supported by the v1 web api.
package v1

import (
	"net/http"

	"github.com/ardanlabs/liarsdice/app/engine/handlers/v1/gamegrp"
	"github.com/ardanlabs/liarsdice/business/core/game"
	"github.com/ardanlabs/liarsdice/business/web/auth"
	"github.com/ardanlabs/liarsdice/business/web/v1/mid"
	"github.com/ardanlabs/liarsdice/foundation/events"
	"github.com/ardanlabs/liarsdice/foundation/web"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log    *zap.SugaredLogger
	Auth   *auth.Auth
	Banker game.Banker
	Evts   *events.Events
}

// Routes binds all the version 1 routes.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	// Register group endpoints.
	ggh := gamegrp.Handlers{
		Banker: cfg.Banker,
		Evts:   cfg.Evts,
		WS:     websocket.Upgrader{},
		Auth:   cfg.Auth,
	}

	app.Handle(http.MethodGet, version, "/game/events", ggh.Events)
	app.Handle(http.MethodGet, version, "/game/status", ggh.Status)

	app.Handle(http.MethodPost, version, "/game/connect", ggh.Connect)

	app.Handle(http.MethodGet, version, "/game/new/:ante", ggh.NewGame, mid.Authenticate(cfg.Auth))
	app.Handle(http.MethodGet, version, "/game/join", ggh.Join, mid.Authenticate(cfg.Auth))

	app.Handle(http.MethodGet, version, "/game/start", ggh.StartGame, mid.Authenticate(cfg.Auth))
	app.Handle(http.MethodGet, version, "/game/rolldice", ggh.RollDice, mid.Authenticate(cfg.Auth))

	app.Handle(http.MethodGet, version, "/game/claim/:number/:suite", ggh.Claim, mid.Authenticate(cfg.Auth))
	app.Handle(http.MethodGet, version, "/game/liar", ggh.CallLiar, mid.Authenticate(cfg.Auth))

	app.Handle(http.MethodGet, version, "/game/reconcile", ggh.Reconcile, mid.Authenticate(cfg.Auth))
	app.Handle(http.MethodGet, version, "/game/balance", ggh.Balance, mid.Authenticate(cfg.Auth))

	// Timeout Situations with a player
	app.Handle(http.MethodGet, version, "/game/next", ggh.NextTurn, mid.Authenticate(cfg.Auth))
	app.Handle(http.MethodGet, version, "/game/out/:outs", ggh.UpdateOut, mid.Authenticate(cfg.Auth))
}
