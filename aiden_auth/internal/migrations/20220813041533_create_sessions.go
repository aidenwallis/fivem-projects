package migrations

import (
	"context"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		_, err := db.ExecContext(ctx, `
			CREATE TABLE sessions (
				id INT NOT NULL AUTO_INCREMENT,
				token_hash VARCHAR(255) NOT NULL,
				metadata TEXT DEFAULT NULL,
				created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
				expires_at DATETIME NOT NULL,
				
				UNIQUE (token_hash),
				INDEX sessions_created_at (created_at),
				INDEX sessions_expires_at (expires_at),
				PRIMARY KEY (id)
			);
		`)
		return err
	}, func(ctx context.Context, db *bun.DB) error {
		_, err := db.ExecContext(ctx, `
			DROP TABLE sessions;
		`)
		return err
	})
}
