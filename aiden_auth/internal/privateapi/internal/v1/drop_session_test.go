package v1_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/config"
	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/fakes"
	v1 "github.com/aidenwallis/fivem-projects/aiden_auth/internal/privateapi/internal/v1"
	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/schema"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestDropSession(t *testing.T) {
	t.Parallel()

	body := schema.DropSessionInput{
		Identifiers: []string{"example:123"},
	}

	t.Run("validation error", func(t *testing.T) {
		t.Parallel()
		w, req := makeReq(t, http.MethodPost, "drop-session", &schema.DropSessionInput{
			Identifiers: []string{},
		})

		r := chi.NewRouter()
		v1.NewVersion(nil, nil)(r)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()
		w, req := makeReq(t, http.MethodPost, "drop-session", body)

		backend := &fakes.FakeBackend{}
		backend.DropSessionReturns(errors.New("expected"))

		r := chi.NewRouter()
		v1.NewVersion(backend, &config.NoopLogger{})(r)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		w, req := makeReq(t, http.MethodPost, "drop-session", body)

		backend := &fakes.FakeBackend{}
		backend.DropSessionCalls(func(_ context.Context, identifiers []string) error {
			assert.Equal(t, body.Identifiers, identifiers)
			return nil
		})

		r := chi.NewRouter()
		v1.NewVersion(backend, &config.NoopLogger{})(r)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.Equal(t, 1, backend.DropSessionCallCount())
	})
}
