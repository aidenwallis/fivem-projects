package utils

import (
	"net/http"

	"go.uber.org/zap"
)

// WithRequestLogger returns a logger that catches the request metadata in the zap field
func WithRequestLogger(log *zap.Logger, req *http.Request) *zap.Logger {
	return log.With(zap.String("request", req.Method+" "+req.URL.Path))
}
