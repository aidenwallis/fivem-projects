package backend

import (
	"context"

	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/db/models"
	"github.com/pkg/errors"
)

// ValidateSession validates that a session exists
func (b *backendImpl) ValidateSession(ctx context.Context, token string) (*models.Session, error) {
	session, err := b.db.Session(ctx, HashToken(token))
	if err != nil {
		return nil, errors.Wrap(err, "fetching session")
	}
	if session == nil || session.Expired() {
		return nil, nil
	}
	return session, nil
}
