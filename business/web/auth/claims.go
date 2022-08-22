package auth

import (
	"context"
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

// Claims represents the authorization claims transmitted via a JWT.
type Claims struct {
	jwt.RegisteredClaims
}

// ctxKey represents the type of value for the context key.
type ctxKey int

// key is used to store/retrieve a Claims value from a context.Context.
const key ctxKey = 1

// SetClaims stores the claims in the context.
func SetClaims(ctx context.Context, claims Claims) context.Context {
	return context.WithValue(ctx, key, claims)
}

// GetClaims returns the claims from the context.
func GetClaims(ctx context.Context) (Claims, error) {
	v, ok := ctx.Value(key).(Claims)
	if !ok {
		return Claims{}, errors.New("claim value missing from context")
	}
	return v, nil
}
