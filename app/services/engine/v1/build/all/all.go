// Package all binds all the routes into the specified app.
package all

import (
	"github.com/ardanlabs/liarsdice/app/services/engine/v1/handlers/checkgrp"
	"github.com/ardanlabs/liarsdice/app/services/engine/v1/handlers/gamegrp"
	"github.com/ardanlabs/liarsdice/business/web/v1/mux"
	"github.com/ardanlabs/liarsdice/foundation/web"
)

// Routes constructs the add value which provides the implementation of
// of RouteAdder for specifying what routes to bind to this instance.
func Routes() add {
	return add{}
}

type add struct{}

// Add implements the RouterAdder interface.
func (add) Add(app *web.App, cfg mux.Config) {
	checkgrp.Routes(app, checkgrp.Config{
		Build: cfg.Build,
		Log:   cfg.Log,
	})

	gamegrp.Routes(app, gamegrp.Config{
		Log:            cfg.Log,
		Auth:           cfg.Auth,
		Converter:      cfg.Converter,
		Bank:           cfg.Bank,
		Evts:           cfg.Evts,
		AnteUSD:        cfg.AnteUSD,
		ActiveKID:      cfg.ActiveKID,
		BankTimeout:    cfg.BankTimeout,
		ConnectTimeout: cfg.ConnectTimeout,
	})
}
