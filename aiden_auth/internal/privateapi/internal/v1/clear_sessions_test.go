package v1_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/config"
	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/fakes"
	v1 "github.com/aidenwallis/fivem-projects/aiden_auth/internal/privateapi/internal/v1"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestClearSessions(t *testing.T) {
	t.Parallel()

	t.Run("error", func(t *testing.T) {
		t.Parallel()

		w, req := makeReq(t, http.MethodPost, "clear-sessions", nil)
		backend := &fakes.FakeBackend{}

		backend.ClearSessionsReturns(errors.New("unexpected"))
		r := chi.NewRouter()
		v1.NewVersion(backend, &config.NoopLogger{})(r)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		w, req := makeReq(t, http.MethodPost, "clear-sessions", nil)
		backend := &fakes.FakeBackend{}
		backend.ClearSessionsReturns(nil)

		r := chi.NewRouter()
		v1.NewVersion(backend, &config.NoopLogger{})(r)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})
}
