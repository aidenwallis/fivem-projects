package db

import (
	"context"
	"database/sql"

	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/db/models"
)

// Session returns a session instance by it's token hash
func (d *dbImpl) Session(ctx context.Context, tokenHash string) (*models.Session, error) {
	var resp models.Session
	err := d.db.NewSelect().
		Model(&resp).
		Relation("Identifiers").
		Where("token_hash = ?", tokenHash).
		Limit(1).
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &resp, nil
}
