package privateapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/fakes"
	"github.com/stretchr/testify/assert"
)

func TestHealthcheck(t *testing.T) {
	t.Parallel()

	t.Run("unhealthy", func(t *testing.T) {
		backend := &fakes.FakeBackend{}
		backend.IsHealthyReturns(false)
		w := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "http://test/healthcheck", nil)
		assert.NoError(t, err)
		healthcheck(backend).ServeHTTP(w, req)
		assert.Equal(t, http.StatusServiceUnavailable, w.Result().StatusCode)
	})

	t.Run("healthy", func(t *testing.T) {
		t.Parallel()
		backend := &fakes.FakeBackend{}
		w := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "http://test/healthcheck", nil)
		assert.NoError(t, err)
		backend.IsHealthyCalls(func(c context.Context) bool {
			assert.Equal(t, req.Context(), c)
			return true
		})
		healthcheck(backend).ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})
}
