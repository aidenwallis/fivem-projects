package config

import (
	"encoding/json"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

// AppConfig defines the application wide config
type AppConfig struct {
	Environment Environment     `json:"environment" validate:"required"`
	Database    *DatabaseConfig `json:"database" validate:"required"`
	Servers     *ServersConfig  `json:"servers" validate:"required"`
	Sessions    *SessionsConfig `json:"sessions" validate:"required"`
}

// DatabaseConfig defines the configuration for the database
type DatabaseConfig struct {
	URL string `json:"url" validate:"required"`
}

// SessionsConfig defines the configuration for sessions
type SessionsConfig struct {
	LifetimeSeconds int `json:"lifetime_seconds" validate:"min=1,max=1209600"`
}

// ServersConfig defines the configuration for HTTP servers
type ServersConfig struct {
	Private *ServerConfig `json:"private" validate:"required"`
	Public  *ServerConfig `json:"public" validate:"required"`
}

// ServerConfig defines the configuration for a HTTP server
type ServerConfig struct {
	Transport string `json:"transport" validate:"required"`
	Addr      string `json:"addr" validate:"required"`
}

// NewAppConfig creates a new instance of app config
func NewAppConfig(filePath string) (*AppConfig, error) {
	pipe, err := os.Open(filePath)
	if err != nil {
		return nil, errors.Wrap(err, "opening config file")
	}
	defer pipe.Close()

	var cfg AppConfig
	if err := json.NewDecoder(pipe).Decode(&cfg); err != nil {
		return nil, errors.Wrap(err, "unmarshalling json")
	}

	if err := validator.New().Struct(cfg); err != nil {
		return nil, errors.Wrap(err, "validating config")
	}

	// ensures we are using the correct env type, it'll force it to prod if it's invalid
	cfg.Environment = resolveEnvironment(cfg.Environment)

	return &cfg, nil
}
