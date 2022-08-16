package backend_test

import (
	"context"
	"testing"
	"time"

	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/db/models"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestValidateSession(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	t.Run("error", func(t *testing.T) {
		t.Parallel()

		b, db, _ := newTestEnvironment()
		expectedErr := errors.New("expected")
		db.SessionReturns(nil, expectedErr)
		resp, err := b.ValidateSession(ctx, fakeToken)
		assert.Nil(t, resp)
		assert.Equal(t, expectedErr, errors.Cause(err))
	})

	t.Run("expired", func(t *testing.T) {
		t.Parallel()

		testCases := map[string]*models.Session{
			"nil session": nil,
			"expired session": {
				ExpiresAt: time.Now().Add(-time.Hour * 24).UTC(),
			},
		}

		for name, testCase := range testCases {
			testCase := testCase

			t.Run(name, func(t *testing.T) {
				t.Parallel()
				b, db, _ := newTestEnvironment()
				db.SessionReturns(testCase, nil)
				resp, err := b.ValidateSession(ctx, fakeToken)
				assert.Nil(t, resp)
				assert.NoError(t, err)
			})
		}
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		b, db, _ := newTestEnvironment()
		sess := &models.Session{
			ExpiresAt: time.Now().Add(time.Hour * 24).UTC(),
		}
		db.SessionCalls(func(c context.Context, v string) (*models.Session, error) {
			assert.Equal(t, ctx, c)
			assert.Equal(t, hashedFakeToken, v)
			return sess, nil
		})
		resp, err := b.ValidateSession(ctx, fakeToken)
		assert.Equal(t, sess, resp)
		assert.NoError(t, err)
	})
}
