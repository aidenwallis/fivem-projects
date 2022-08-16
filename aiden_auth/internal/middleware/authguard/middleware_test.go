package authguard_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/db/models"
	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/middleware/auth"
	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/middleware/authguard"
	"github.com/aidenwallis/go-write/write"
	"github.com/stretchr/testify/assert"
)

func TestMiddleware(t *testing.T) {
	t.Parallel()

	handler := func(t *testing.T) http.Handler {
		return authguard.Middleware(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			assert.NoError(t, write.Teapot(w).Empty())
		}))
	}

	t.Run("unauthed", func(t *testing.T) {
		t.Parallel()
		w := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "http://test/middleware", nil)
		assert.NoError(t, err)
		handler(t).ServeHTTP(w, req)
		assert.Equal(t, http.StatusUnauthorized, w.Result().StatusCode)
	})

	t.Run("authed", func(t *testing.T) {
		t.Parallel()
		w := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "http://test/middleware", nil)
		assert.NoError(t, err)
		handler(t).ServeHTTP(w, req.WithContext(auth.WithSession(req.Context(), &models.Session{})))
		assert.Equal(t, http.StatusTeapot, w.Result().StatusCode)
	})
}
