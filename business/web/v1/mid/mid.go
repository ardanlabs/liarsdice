// Package mid contains the set of middleware functions.
package mid

import (
	"context"

	"github.com/ardanlabs/liarsdice/business/web/v1/auth"
	"github.com/ethereum/go-ethereum/common"
)

type ctxKey int

const claimKey ctxKey = 1

func setClaims(ctx context.Context, claims auth.Claims) context.Context {
	return context.WithValue(ctx, claimKey, claims)
}

// GetClaims provides access to the claims from the context.
func GetClaims(ctx context.Context) auth.Claims {
	v, ok := ctx.Value(claimKey).(auth.Claims)
	if !ok {
		return auth.Claims{}
	}
	return v
}

// GetSubject provides access to the subject from the claims.
func GetSubject(ctx context.Context) common.Address {
	return common.HexToAddress(GetClaims(ctx).Subject)
}
