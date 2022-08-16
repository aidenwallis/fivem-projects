package backend_test

import (
	"testing"

	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/backend"
	"github.com/stretchr/testify/assert"
)

func TestRandomToken(t *testing.T) {
	t.Parallel()

	const iterations = 100

	// ensure token is still unique after 100 passes
	values := make(map[string]struct{}, iterations)

	for i := 0; i < iterations; i++ {
		token, err := backend.RandomToken(24)
		assert.NoError(t, err, token, i)

		_, ok := values[token]
		assert.False(t, ok, token, i)

		values[token] = struct{}{}
	}
}

func TestHashToken(t *testing.T) {
	t.Parallel()
	assert.Equal(t, hashedFakeToken, backend.HashToken(fakeToken))
}
