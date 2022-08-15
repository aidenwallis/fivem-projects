package models

import (
	"time"

	"github.com/uptrace/bun"
)

// Session represents a row from the sessions table
type Session struct {
	bun.BaseModel `bun:"table:sessions,alias:u"`

	ID          int                  `bun:"id,pk"`
	TokenHash   string               `bun:"token_hash"`
	Metadata    string               `bun:"metadata"`
	Identifiers []*SessionIdentifier `bun:"rel:has-many,join:id=session_id"`
	CreatedAt   time.Time            `bun:"created_at"`
	ExpiresAt   time.Time            `bun:"expires_at"`
}

// Expired returns true if the session expired
func (m *Session) Expired() bool {
	return time.Now().UTC().After(m.ExpiresAt.UTC())
}
