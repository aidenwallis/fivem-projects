package backend_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClearSessions(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	t.Run("error", func(t *testing.T) {
		t.Parallel()
		b, db, log := newTestEnvironment()
		err := errors.New("expected")
		db.ClearSessionsReturns(0, err)
		assert.Equal(t, err, b.ClearSessions(ctx))
		assert.Equal(t, 1, db.ClearSessionsCallCount())
		assert.Equal(t, 0, log.InfoCallCount())
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		b, db, log := newTestEnvironment()
		db.ClearSessionsCalls(func(c context.Context, ts time.Time) (int, error) {
			assert.Equal(t, ctx, c)
			assert.WithinDuration(t, time.Now().UTC(), ts, time.Minute)
			return 123, nil
		})
		assert.NoError(t, b.ClearSessions(ctx))
		assert.Equal(t, 1, db.ClearSessionsCallCount())
		assert.Equal(t, 1, log.InfoCallCount())
	})
}
