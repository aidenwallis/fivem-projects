package backend

import (
	"context"
	"time"

	"go.uber.org/zap"
)

// ClearSessions will clear all sessions in the DB
func (b *backendImpl) ClearSessions(ctx context.Context) error {
	count, err := b.db.ClearSessions(ctx, time.Now().UTC())
	if err != nil {
		return err
	}

	b.log.Info("Reset sessions.", zap.Int("deleted-count", count))
	return nil
}
