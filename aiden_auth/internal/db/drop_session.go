package db

import (
	"context"

	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/db/models"
	"github.com/uptrace/bun"
)

// DropSession will drop all attached sessions to the given identifiers
func (d *dbImpl) DropSession(ctx context.Context, identifiers []string) (int, error) {
	if len(identifiers) == 0 {
		return 0, nil
	}

	resp, err := d.db.NewDelete().
		Model(&models.Session{}).
		Where("id IN (SELECT session_id FROM session_identifiers WHERE identifier IN (?))", bun.In(identifiers)).
		Exec(ctx)
	if err != nil {
		return 0, err
	}

	count, err := resp.RowsAffected()
	return int(count), err
}
