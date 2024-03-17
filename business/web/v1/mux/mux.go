// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"net/http"
	"os"
	"time"

	"github.com/ardanlabs/ethereum/currency"
	"github.com/ardanlabs/liarsdice/business/core/bank"
	"github.com/ardanlabs/liarsdice/business/web/v1/auth"
	"github.com/ardanlabs/liarsdice/business/web/v1/mid"
	"github.com/ardanlabs/liarsdice/foundation/logger"
	"github.com/ardanlabs/liarsdice/foundation/web"
	"github.com/jmoiron/sqlx"
)

// Options represent optional parameters.
type Options struct {
	corsOrigin []string
}

// WithCORS provides configuration options for CORS.
func WithCORS(origins []string) func(opts *Options) {
	return func(opts *Options) {
		opts.corsOrigin = origins
	}
}

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Build          string
	Shutdown       chan os.Signal
	Log            *logger.Logger
	Auth           *auth.Auth
	Converter      *currency.Converter
	Bank           *bank.Bank
	DB             *sqlx.DB
	AnteUSD        float64
	ActiveKID      string
	BankTimeout    time.Duration
	ConnectTimeout time.Duration
}

// RouteAdder defines behavior that sets the routes to bind for an instance
// of the service.
type RouteAdder interface {
	Add(app *web.App, cfg Config)
}

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI(cfg Config, routeAdder RouteAdder, options ...func(opts *Options)) http.Handler {
	var opts Options
	for _, option := range options {
		option(&opts)
	}

	app := web.NewApp(
		cfg.Shutdown,
		mid.Logger(cfg.Log),
		mid.Errors(cfg.Log),
		mid.Metrics(),
		mid.Panics(),
		mid.Cors(opts.corsOrigin),
	)

	if len(opts.corsOrigin) > 0 {
		app.EnableCORS(mid.Cors(opts.corsOrigin))
	}

	routeAdder.Add(app, cfg)

	return app
}
