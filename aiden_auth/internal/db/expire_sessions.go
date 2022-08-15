package db

import (
	"context"
	"time"

	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/db/models"
)

// ExpireSession evicts all tokens that have passed their expiry
func (d *dbImpl) ExpireSessions(ctx context.Context) (int, error) {
	resp, err := d.db.NewDelete().Model(&models.Session{}).Where("expires_at <= ?", time.Now().UTC()).Exec(ctx)
	if err != nil {
		return 0, err
	}

	count, err := resp.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(count), nil
}
