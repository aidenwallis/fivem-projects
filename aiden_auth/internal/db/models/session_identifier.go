package models

import "github.com/uptrace/bun"

// SessionIdentifier represents a row from the session_identifiers table
type SessionIdentifier struct {
	bun.BaseModel `bun:"session_identifiers"`

	SessionID  int    `bun:"session_id,pk"`
	Identifier string `bun:"identifier,pk"`
}
