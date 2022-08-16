package auth_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/db/models"
	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/fakes"
	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/middleware/auth"
	"github.com/aidenwallis/go-write/write"
	"github.com/stretchr/testify/assert"
)

func TestMiddleware(t *testing.T) {
	t.Parallel()

	const fakeToken = "abc123"

	makeReq := func(t *testing.T) *http.Request {
		req, err := http.NewRequest(http.MethodGet, "http://test/middleware", nil)
		assert.NoError(t, err)
		return req
	}

	const status = http.StatusTeapot
	h := func(t *testing.T, shouldBeAuthed bool) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			assert.Equal(t, shouldBeAuthed, auth.HasSession(req.Context()))
			_ = write.New(w, status).Empty()
		})
	}

	t.Run("no header", func(t *testing.T) {
		backend := &fakes.FakeBackend{}
		log := &fakes.FakeLogger{}
		w := httptest.NewRecorder()
		req := makeReq(t)
		auth.Middleware(backend, log)(h(t, false)).ServeHTTP(w, req)
		assert.Equal(t, status, w.Result().StatusCode)
		assert.Equal(t, 0, backend.ValidateSessionCallCount())
	})

	t.Run("error", func(t *testing.T) {
		backend := &fakes.FakeBackend{}
		backend.ValidateSessionReturns(nil, errors.New("expected"))
		log := &fakes.FakeLogger{}
		w := httptest.NewRecorder()
		req := makeReq(t)
		req.Header.Set(auth.HeaderKey, fakeToken)
		auth.Middleware(backend, log)(h(t, false)).ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
		assert.Equal(t, 1, log.ErrorCallCount())
	})

	t.Run("invalid token", func(t *testing.T) {
		backend := &fakes.FakeBackend{}
		backend.ValidateSessionReturns(nil, nil)
		log := &fakes.FakeLogger{}
		w := httptest.NewRecorder()
		req := makeReq(t)
		auth.Middleware(backend, log)(h(t, false)).ServeHTTP(w, req)
		assert.Equal(t, status, w.Result().StatusCode)
		assert.Equal(t, 0, log.ErrorCallCount())
	})

	t.Run("valid token", func(t *testing.T) {
		backend := &fakes.FakeBackend{}

		w := httptest.NewRecorder()
		log := &fakes.FakeLogger{}

		req := makeReq(t)
		req.Header.Set(auth.HeaderKey, fakeToken)

		backend.ValidateSessionCalls(func(c context.Context, v string) (*models.Session, error) {
			assert.Equal(t, req.Context(), c)
			assert.Equal(t, fakeToken, v)
			return &models.Session{}, nil
		})

		auth.Middleware(backend, log)(h(t, true)).ServeHTTP(w, req)
		assert.Equal(t, status, w.Result().StatusCode)
		assert.Equal(t, 0, log.ErrorCallCount())
	})
}
