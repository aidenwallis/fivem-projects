package auth

import (
	"context"

	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/db/models"
)

type contextKey int

const (
	sessionKey contextKey = iota
)

// GetSession returns the current session in context, or nil
func GetSession(ctx context.Context) *models.Session {
	if v, ok := ctx.Value(sessionKey).(*models.Session); ok {
		return v
	}
	return nil
}

// HasSession returns whether a session exists in context
func HasSession(ctx context.Context) bool {
	return GetSession(ctx) != nil
}

// WithSession returns a new context with the session included
func WithSession(ctx context.Context, session *models.Session) context.Context {
	return context.WithValue(ctx, sessionKey, session)
}
