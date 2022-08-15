package backend

import (
	"context"

	"go.uber.org/zap"
)

// IsHealthy checks whether the backend is healthy
func (b *backendImpl) IsHealthy(ctx context.Context) bool {
	if err := b.db.Ping(ctx); err != nil {
		b.log.Error("failed to ping database?", zap.Error(err))
		return false
	}
	return true
}
