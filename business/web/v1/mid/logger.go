package mid

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ardanlabs/liarsdice/foundation/logger"
	"github.com/ardanlabs/liarsdice/foundation/web"
)

// Logger writes some information about the request to the logs in the
// format: TraceID : (200) GET /foo -> IP ADDR (latency)
func Logger(log *logger.Logger) web.MidHandler {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			v := web.GetValues(ctx)

			path := r.URL.Path
			if r.URL.RawQuery != "" {
				path = fmt.Sprintf("%s?%s", path, r.URL.RawQuery)
			}

			log.Info(ctx, "request started", "method", r.Method, "path", path, "remoteaddr", r.RemoteAddr)

			err := handler(ctx, w, r)

			log.Info(ctx, "request completed", "method", r.Method, "path", path, "remoteaddr", r.RemoteAddr,
				"statuscode", v.StatusCode, "since", time.Since(v.Now).String())

			return err
		}

		return h
	}

	return m
}
