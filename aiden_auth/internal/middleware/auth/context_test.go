package auth_test

import (
	"context"
	"testing"

	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/db/models"
	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/middleware/auth"
	"github.com/stretchr/testify/assert"
)

func TestContext(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	sess := &models.Session{TokenHash: "abc123"}

	assert.Nil(t, auth.GetSession(ctx))
	assert.False(t, auth.HasSession(ctx))

	ctx = auth.WithSession(ctx, sess)

	assert.Equal(t, sess, auth.GetSession(ctx))
	assert.True(t, auth.HasSession(ctx))
}
