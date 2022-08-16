package v1_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/config"
	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/db/models"
	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/fakes"
	v1 "github.com/aidenwallis/fivem-projects/aiden_auth/internal/privateapi/internal/v1"
	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/schema"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestCreateSession(t *testing.T) {
	t.Parallel()

	body := schema.SessionInput{
		Identifiers: []string{"example:123"},
	}

	t.Run("validation error", func(t *testing.T) {
		t.Parallel()
		w, req := makeReq(t, http.MethodPost, "sessions", &schema.DropSessionInput{
			Identifiers: []string{},
		})

		r := chi.NewRouter()
		v1.NewVersion(nil, nil)(r)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()
		w, req := makeReq(t, http.MethodPost, "sessions", body)

		backend := &fakes.FakeBackend{}
		backend.CreateSessionReturns(nil, "", errors.New("expected"))

		r := chi.NewRouter()
		v1.NewVersion(backend, &config.NoopLogger{})(r)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		w, req := makeReq(t, http.MethodPost, "sessions", body)

		backend := &fakes.FakeBackend{}
		backend.CreateSessionCalls(func(_ context.Context, identifiers []string, _ json.RawMessage) (*models.Session, string, error) {
			assert.Equal(t, body.Identifiers, identifiers)
			return &models.Session{}, "token123", nil
		})

		r := chi.NewRouter()
		v1.NewVersion(backend, &config.NoopLogger{})(r)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.Equal(t, 1, backend.CreateSessionCallCount())
	})
}
