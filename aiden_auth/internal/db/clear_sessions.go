package db

import (
	"context"
	"time"

	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/db/models"
)

// ClearSessions will revoke all sessions that were created before a given time.
func (d *dbImpl) ClearSessions(ctx context.Context, before time.Time) (int, error) {
	resp, err := d.db.NewDelete().
		Model(&models.Session{}).
		Where("created_at <= ?", before).
		Exec(ctx)
	if err != nil {
		return 0, err
	}
	count, err := resp.RowsAffected()
	return int(count), err
}
