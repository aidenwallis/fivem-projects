package db

import (
	"context"
	"database/sql"

	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/db/models"
	"github.com/aidenwallis/go-utils/utils"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// CreateSession creates a session instance and all the identifier rows, or fails/rolls back everything
func (d *dbImpl) CreateSession(ctx context.Context, session *models.Session, identifiers []string) error {
	// as we are inserting rows into multiple tables at once, we'll want to make sure that the attached session also
	// has all identiifers before returning, or fail everything.
	tx, err := d.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return errors.Wrap(err, "begin tx")
	}

	// rollback if the func fails at any point, since we commit at the end, this won't matter if it all worked fine.
	defer func() {
		if err := tx.Rollback(); err != nil && err != sql.ErrTxDone {
			d.log.Error("failed to rollback transaction", zap.Error(err))
		}
	}()

	sessionResp, err := tx.NewInsert().Model(session).Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "inseting session")
	}

	sessionID, err := sessionResp.LastInsertId()
	if err != nil {
		return errors.Wrap(err, "resolving session id")
	}

	session.ID = int(sessionID)
	identifierModels := utils.SliceMap(identifiers, func(identifier string) *models.SessionIdentifier {
		return &models.SessionIdentifier{
			SessionID:  session.ID,
			Identifier: identifier,
		}
	})

	_, err = tx.NewInsert().Model(&identifierModels).Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "inserting identiifers")
	}

	session.Identifiers = identifierModels

	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "committing")
	}

	return nil
}
