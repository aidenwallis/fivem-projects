package v1_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/db/models"
	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/fakes"
	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/middleware/auth"
	v1 "github.com/aidenwallis/fivem-projects/aiden_auth/internal/publicapi/internal/v1"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestGetSession(t *testing.T) {
	t.Parallel()

	t.Run("unauthed", func(t *testing.T) {
		t.Parallel()

		req, err := http.NewRequest(http.MethodGet, "http://test/sessions", nil)
		assert.NoError(t, err)

		r := chi.NewRouter()
		v1.NewVersion(nil, nil)(r)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Result().StatusCode)
	})

	t.Run("authed", func(t *testing.T) {
		t.Parallel()

		req, err := http.NewRequest(http.MethodGet, "http://test/sessions", nil)
		assert.NoError(t, err)

		req.Header.Set(auth.HeaderKey, "abc123")

		backend := &fakes.FakeBackend{}
		backend.ValidateSessionReturns(&models.Session{}, nil)

		r := chi.NewRouter()
		v1.NewVersion(backend, nil)(r)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})
}
