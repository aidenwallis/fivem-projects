package backend_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsHealthy(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	t.Run("error", func(t *testing.T) {
		t.Parallel()
		b, db, log := newTestEnvironment()
		db.PingReturns(errors.New("expected"))
		assert.False(t, b.IsHealthy(ctx))
		assert.Equal(t, 1, log.ErrorCallCount())
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		b, db, log := newTestEnvironment()
		db.PingCalls(func(c context.Context) error {
			assert.Equal(t, ctx, c)
			return nil
		})
		assert.True(t, b.IsHealthy(ctx))
		assert.Equal(t, 0, log.ErrorCallCount())
	})
}
