package auth

import (
	"net/http"

	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/backend"
	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/schema"
	"github.com/aidenwallis/go-write/write"
	"go.uber.org/zap"
)

const (
	headerKey = "x-fivem-auth"
)

// Middleware is the auth middleware, it allows most traffic through, but will authorize requests
// as it passes through.
func Middleware(backendImpl backend.Backend, log *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			header := req.Header.Get(headerKey)
			if header == "" {
				next.ServeHTTP(w, req)
				return
			}

			ctx := req.Context()
			session, err := backendImpl.ValidateSession(ctx, header)
			if err != nil {
				log.Error("failed to validate session", zap.Error(err))
				_ = write.InternalServerError(w).JSON(schema.UnknownError)
				return
			}

			if session != nil {
				ctx = WithSession(ctx, session)
			}

			next.ServeHTTP(w, req.WithContext(ctx))
		})
	}
}
