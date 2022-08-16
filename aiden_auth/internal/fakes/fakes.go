package fakes

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

// This package contains all fakes we generate for this service.

//counterfeiter:generate -o backend.go github.com/aidenwallis/fivem-projects/aiden_auth/internal/backend.Backend
//counterfeiter:generate -o db.go github.com/aidenwallis/fivem-projects/aiden_auth/internal/db.DB
//counterfeiter:generate -o logger.go github.com/aidenwallis/fivem-projects/aiden_auth/internal/config.Logger
