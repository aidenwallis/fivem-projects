package privateapi

import (
	"net/http"

	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/backend"
	v1 "github.com/aidenwallis/fivem-projects/aiden_auth/internal/privateapi/internal/v1"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

// NewServer creates a new privateapi server
func NewServer(backendImpl backend.Backend, log *zap.Logger) http.Handler {
	r := chi.NewRouter()

	r.Get("/healthz", healthcheck(backendImpl))
	r.Route("/v1", v1.NewVersion(backendImpl, log))

	return r
}
