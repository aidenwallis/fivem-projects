package privateapi

import (
	"net/http"

	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/backend"
	"github.com/aidenwallis/go-write/write"
)

// healthcheck is a HTTP handler which provides healthchecks to downstream dependencies
func healthcheck(backendImpl backend.Backend) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if backendImpl.IsHealthy(req.Context()) {
			_ = write.OK(w).Text("OK")
			return
		}

		_ = write.ServiceUnavailable(w).Text("Unhealthy")
	})
}
