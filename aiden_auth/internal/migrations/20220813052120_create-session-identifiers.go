package migrations

import (
	"context"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		_, err := db.ExecContext(ctx, `
			CREATE TABLE session_identifiers (
				identifier VARCHAR(255) NOT NULL,
				session_id INT NOT NULL,

				PRIMARY KEY (session_id, identifier),
				INDEX session_identifiers_identifier (identifier),

				FOREIGN KEY (session_id) REFERENCES sessions(id) ON DELETE CASCADE
			);
		`)
		return err
	}, func(ctx context.Context, db *bun.DB) error {
		_, err := db.ExecContext(ctx, `
			DROP TABLE session_identifiers;
		`)
		return err
	})
}
