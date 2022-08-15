package config

import "go.uber.org/zap"

// NewLogger returns a logger based on the environment that is set
func NewLogger(v Environment) (*zap.Logger, error) {
	if v == EnvDevelopment {
		return zap.NewDevelopment()
	}
	return zap.NewProduction()
}
