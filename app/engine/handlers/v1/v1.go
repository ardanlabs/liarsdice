// Package v1 contains the full set of handler functions and routes
// supported by the v1 web api.
package v1

import (
	"net/http"

	"github.com/ardanlabs/liarsdice/app/engine/handlers/v1/gamegrp"
	"github.com/ardanlabs/liarsdice/business/core/game"
	"github.com/ardanlabs/liarsdice/business/web/auth"
	"github.com/ardanlabs/liarsdice/foundation/events"
	"github.com/ardanlabs/liarsdice/foundation/web"
	"github.com/gorilla/websocket"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log  *zap.SugaredLogger
	Auth *auth.Auth
	DB   *sqlx.DB
	Game *game.Game
	Evts *events.Events
}

// Routes binds all the version 1 routes.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	// Register group endpoints.
	ggh := gamegrp.Handlers{
		Game: cfg.Game,
		Evts: cfg.Evts,
		WS:   websocket.Upgrader{},
	}

	app.Handle(http.MethodGet, version, "/game/events", ggh.Events)
	app.Handle(http.MethodGet, version, "/game/status", ggh.Status)
	app.Handle(http.MethodGet, version, "/game/rolldice/:wallet", ggh.RollDice)
	app.Handle(http.MethodGet, version, "/game/balance/:wallet", ggh.Balance)
	app.Handle(http.MethodGet, version, "/game/newround", ggh.NewRound)

	app.Handle(http.MethodPost, version, "/game/join", ggh.Join)
	app.Handle(http.MethodPost, version, "/game/start", ggh.Start)
	app.Handle(http.MethodPost, version, "/game/claim/:wallet", ggh.Claim)
	app.Handle(http.MethodPost, version, "/game/callliar/:wallet", ggh.CallLiar)
}
