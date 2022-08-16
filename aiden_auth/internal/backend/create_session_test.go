package backend_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/backend"
	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/db/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateSession(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	identifiers := []string{"example:abc123"}

	t.Run("error", func(t *testing.T) {
		t.Parallel()

		b, db, log := newTestEnvironment()
		db.CreateSessionReturns(errors.New("expected"))
		sess, token, err := b.CreateSession(ctx, identifiers, nil)
		assert.Nil(t, sess)
		assert.Empty(t, token)
		assert.Equal(t, backend.ErrAttemptsExceeded, err)
		assert.Equal(t, backend.SessionAttempts, log.ErrorCallCount())
		assert.Equal(t, backend.SessionAttempts, db.CreateSessionCallCount())
	})

	t.Run("success", func(t *testing.T) {
		b, db, log := newTestEnvironment()

		hashedTokenValue := "" // grab value from stub and store it in this scope
		db.CreateSessionCalls(func(c context.Context, sess *models.Session, is []string) error {
			assert.WithinDuration(t, time.Now().Add(time.Second*time.Duration(lifetimeSeconds)).UTC(), sess.ExpiresAt, time.Minute)
			hashedTokenValue = sess.TokenHash
			assert.Equal(t, identifiers, is)
			return nil
		})

		sess, token, err := b.CreateSession(ctx, identifiers, nil)

		assert.NotNil(t, sess)
		assert.Equal(t, backend.HashToken(token), hashedTokenValue)
		assert.NotEmpty(t, token)
		assert.NoError(t, err)
		assert.Equal(t, 0, log.ErrorCallCount())
		assert.Equal(t, 1, db.CreateSessionCallCount())
	})
}
