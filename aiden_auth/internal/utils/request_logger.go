package utils

import (
	"net/http"

	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/config"
	"go.uber.org/zap"
)

// WithRequestLogger returns a logger that catches the request metadata in the zap field
func WithRequestLogger(log config.Logger, req *http.Request) config.Logger {
	return log.With(zap.String("request", req.Method+" "+req.URL.Path))
}
