package web

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dimfeld/httptreemux/v5"
)

// Param returns the web call parameters from the request.
func Param(ctx context.Context, key string) string {
	m := httptreemux.ContextParams(ctx)
	return m[key]
}

// SetParam sets a param into the context
func SetParam(ctx context.Context, key string, value string) context.Context {
	m := httptreemux.ContextParams(ctx)
	m[key] = value
	return httptreemux.AddParamsToContext(ctx, m)
}

type validator interface {
	Validate() error
}

// Decode reads the body of an HTTP request looking for a JSON document. The
// body is decoded into the provided value.
// If the provided value is a struct then it is checked for validation tags.
// If the value implements a validate function, it is executed.
func Decode(r *http.Request, val any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(val); err != nil {
		return fmt.Errorf("unable to decode payload: %w", err)
	}

	if v, ok := val.(validator); ok {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("unable to validate payload: %w", err)
		}
	}

	return nil
}
