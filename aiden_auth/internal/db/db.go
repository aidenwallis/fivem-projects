package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/config"
	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/db/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"go.uber.org/zap"
)

// DB defines the methods exposed for the database layer.
type DB interface {
	Bun() *bun.DB
	Ping(ctx context.Context) error

	ClearSessions(ctx context.Context, before time.Time) (int, error)
	CreateSession(ctx context.Context, session *models.Session, identifiers []string) error
	DropSession(ctx context.Context, identifiers []string) (int, error)
	ExpireSessions(context.Context) (int, error)
	Session(ctx context.Context, tokenHash string) (*models.Session, error)
}

type dbImpl struct {
	db  *bun.DB
	log *zap.Logger
}

func NewDB(cfg *config.DatabaseConfig, log *zap.Logger) (DB, error) {
	conn, err := sql.Open("mysql", cfg.URL)
	if err != nil {
		return nil, errors.Wrap(err, "connecting to database")
	}

	return &dbImpl{
		db:  bun.NewDB(conn, mysqldialect.New()),
		log: log,
	}, nil
}

func (d *dbImpl) Bun() *bun.DB {
	return d.db
}

func (d *dbImpl) Ping(ctx context.Context) error {
	return d.db.PingContext(ctx)
}
