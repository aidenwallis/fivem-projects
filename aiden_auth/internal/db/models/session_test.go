package models_test

import (
	"testing"
	"time"

	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/db/models"
	"github.com/stretchr/testify/assert"
)

func TestSessionExpired(t *testing.T) {
	t.Parallel()

	// expired ts
	sess := &models.Session{ExpiresAt: time.Now().Add(-time.Hour * 24).UTC()}
	assert.True(t, sess.Expired())

	// not expired
	sess.ExpiresAt = time.Now().Add(time.Hour * 24).UTC()
	assert.False(t, sess.Expired())
}
