package publicapi

import (
	"net/http"

	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/backend"
	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/config"
	v1 "github.com/aidenwallis/fivem-projects/aiden_auth/internal/publicapi/internal/v1"
	"github.com/go-chi/chi/v5"
)

func NewServer(backendImpl backend.Backend, log config.Logger) http.Handler {
	r := chi.NewRouter()

	r.Route("/v1", v1.NewVersion(backendImpl, log))

	return r
}
