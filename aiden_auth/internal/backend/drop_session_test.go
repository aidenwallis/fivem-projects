package backend_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDropSession(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	identifiers := []string{"steam:123", "abc:123"}

	t.Run("error", func(t *testing.T) {
		t.Parallel()
		err := errors.New("expected")
		b, db, log := newTestEnvironment()
		db.DropSessionReturns(0, err)
		assert.Equal(t, err, b.DropSession(ctx, identifiers))
		assert.Equal(t, 0, log.DebugCallCount())
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		b, db, log := newTestEnvironment()
		db.DropSessionCalls(func(c context.Context, v []string) (int, error) {
			assert.Equal(t, ctx, c)
			assert.Equal(t, identifiers, v)
			return 123, nil
		})
		assert.NoError(t, b.DropSession(ctx, identifiers))
		assert.Equal(t, 1, log.DebugCallCount())
	})
}
