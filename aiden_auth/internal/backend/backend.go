package backend

import (
	"context"
	"encoding/json"

	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/config"
	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/db"
	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/db/models"
)

// Backend represents all functions available from the backend.
type Backend interface {
	ClearSessions(context.Context) error
	CreateSession(ctx context.Context, identifiers []string, metadata json.RawMessage) (*models.Session, string, error)
	DropSession(ctx context.Context, identifiers []string) error
	IsHealthy(context.Context) bool
	ValidateSession(ctx context.Context, token string) (*models.Session, error)
}

// backendImpl implements backend
type backendImpl struct {
	sessionCfg *config.SessionsConfig
	db         db.DB
	log        config.Logger
}

func NewBackend(dbImpl db.DB, log config.Logger, sessionCfg *config.SessionsConfig) Backend {
	return &backendImpl{
		db:         dbImpl,
		log:        log,
		sessionCfg: sessionCfg,
	}
}
