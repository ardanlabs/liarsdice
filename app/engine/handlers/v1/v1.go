// Package v1 contains the full set of handler functions and routes
// supported by the v1 web api.
package v1

import (
	"net/http"

	"github.com/ardanlabs/liarsdice/app/engine/handlers/v1/gamegrp"
	"github.com/ardanlabs/liarsdice/business/web/auth"
	"github.com/ardanlabs/liarsdice/foundation/web"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log  *zap.SugaredLogger
	Auth *auth.Auth
	DB   *sqlx.DB
}

// Routes binds all the version 1 routes.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	// Register group endpoints.
	ggh := gamegrp.Handlers{
		Table: gamegrp.NewTable(),
	}

	app.Handle(http.MethodGet, version, "/game/list", ggh.List)
	app.Handle(http.MethodGet, version, "/game/status/:uuid", ggh.Status)

	app.Handle(http.MethodPost, version, "/game/new", ggh.New)
	app.Handle(http.MethodPost, version, "/game/join", ggh.Join)
	app.Handle(http.MethodPost, version, "/game/start/:uuid", ggh.Start)

	// app.Handle(http.MethodGet, version, "/users/:page/:rows", ugh.Query, authen, admin)
	// app.Handle(http.MethodGet, version, "/users/:id", ugh.QueryByID, authen)
	// app.Handle(http.MethodPost, version, "/users", ugh.Create, authen, admin)
	// app.Handle(http.MethodPut, version, "/users/:id", ugh.Update, authen, admin)
	// app.Handle(http.MethodDelete, version, "/users/:id", ugh.Delete, authen, admin)
}
