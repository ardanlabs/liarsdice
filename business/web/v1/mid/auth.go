package mid

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/ardanlabs/liarsdice/business/web/auth"
	v1Web "github.com/ardanlabs/liarsdice/business/web/v1"
	"github.com/ardanlabs/liarsdice/foundation/web"
	"go.uber.org/zap"
)

// Authenticate validates a JWT from the `Authorization` header.
func Authenticate(log *zap.SugaredLogger, a *auth.Auth) web.Middleware {

	// This is the actual middleware function to be executed.
	m := func(handler web.Handler) web.Handler {

		// Create the handler that will be attached in the middleware chain.
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

			// Expecting: bearer <token>
			authStr := r.Header.Get("authorization")

			// Parse the authorization header.
			parts := strings.Split(authStr, " ")
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				err := errors.New("expected authorization header format: bearer <token>")
				return v1Web.NewRequestError(err, http.StatusUnauthorized)
			}

			// Validate the token is signed by us.
			claims, err := a.ValidateToken(parts[1])
			if err != nil {
				return v1Web.NewRequestError(err, http.StatusUnauthorized)
			}

			// Add claims to the context, so they can be retrieved later.
			ctx = auth.SetClaims(ctx, claims)

			log.Infow("request auth", "traceid", web.GetTraceID(ctx), "address", claims.Subject)

			// Call the next handler.
			return handler(ctx, w, r)
		}

		return h
	}

	return m
}
