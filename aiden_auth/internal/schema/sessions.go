package schema

import (
	"encoding/json"
	"time"

	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/db/models"
	"github.com/aidenwallis/go-utils/utils"
)

// Session defines the session schema
//
// swagger:model
type Session struct {
	Identifiers []string               `json:"identifiers"`
	Metadata    map[string]interface{} `json:"metadata"`
	CreatedAt   time.Time              `json:"created_at"`
	ExpiresAt   time.Time              `json:"expires_at"`
}

// NewSession creates a new Session schema based off the session model
func NewSession(sess *models.Session) *Session {
	return &Session{
		Identifiers: utils.SliceMap(sess.Identifiers, func(i *models.SessionIdentifier) string {
			return i.Identifier
		}),
		Metadata:  resolveMetadata(sess.Metadata),
		CreatedAt: sess.CreatedAt,
		ExpiresAt: sess.ExpiresAt,
	}
}

// SessionInput defines the create session input schema
//
// swagger:model
type SessionInput struct {
	Identifiers []string        `json:"identifiers" validate:"min=1"`
	Metadata    json.RawMessage `json:"metadata"`
}

// DropSessionInput defines the drop session input schema
//
// swagger:model
type DropSessionInput struct {
	Identifiers []string `json:"identifiers" validate:"min=1"`
}

func resolveMetadata(v string) map[string]interface{} {
	if v == "" {
		return nil
	}

	var r map[string]interface{}
	if err := json.Unmarshal([]byte(v), &r); err != nil {
		return nil
	}
	return r
}

// CreateSessionResponse is the response sent to clients when a new session is created
//
// swagger:model
type CreateSessionResponse struct {
	*Session
	Token string `json:"token"`
}

// NewCreateSessionResponse creates a new instance of CreateSessionResponse
func NewCreateSessionResponse(sess *models.Session, token string) *CreateSessionResponse {
	return &CreateSessionResponse{
		Session: NewSession(sess),
		Token:   token,
	}
}
