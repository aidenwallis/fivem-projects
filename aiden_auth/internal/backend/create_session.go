package backend

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/db/models"
	"go.uber.org/zap"
)

const (
	sessionLength = 36

	// SessionAttempts is how many times we attempt to create a session before giving up
	SessionAttempts = 10
)

var ErrAttemptsExceeded = fmt.Errorf("could not create a session after %d attempts", SessionAttempts)

// ValidateSession validates that a session exists
func (b *backendImpl) CreateSession(ctx context.Context, identifiers []string, metadata json.RawMessage) (*models.Session, string, error) {
	now := time.Now().UTC()
	expiresAt := now.Add(time.Second * time.Duration(b.sessionCfg.LifetimeSeconds)).UTC()

	for i := 0; i < SessionAttempts; i++ {
		// attempt to create session tokens up to 10 times
		token, err := RandomToken(sessionLength)
		if err != nil {
			continue
		}

		sess := &models.Session{
			TokenHash: HashToken(token),
			Metadata:  string(metadata),
			CreatedAt: now,
			ExpiresAt: expiresAt,
		}

		if err := b.db.CreateSession(ctx, sess, identifiers); err != nil {
			b.log.Error("failed to create session", zap.Error(err))
			continue
		}

		// send back the session, unhashed token (this will be the only place it exists, to the end client), and no err
		// by us only having the hashed token, we can verify that their token exists without needing to hold secrets in
		// db
		return sess, token, nil
	}

	return nil, "", ErrAttemptsExceeded
}
