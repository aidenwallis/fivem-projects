package backend

import (
	"context"

	"go.uber.org/zap"
)

// DropSession drops all sessions for a given set of identifiers
func (b *backendImpl) DropSession(ctx context.Context, identifiers []string) error {
	count, err := b.db.DropSession(ctx, identifiers)
	if err != nil {
		return err
	}

	b.log.Debug("dropped sessions for identifiers", zap.Int("dropped-count", count))
	return nil
}
