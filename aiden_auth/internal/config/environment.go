package config

// Environment defines whether the app is running in development or production
type Environment string

const (
	// EnvDevelopment is the value of Environment when the app is running in development mode
	EnvDevelopment = Environment("development")

	// EnvProduction is the value of Environment when the app is running in production mode
	EnvProduction = Environment("production")
)

// resolveEnvironment attempts to set the app into dev mode, but if any kind of invalid environment
// is passed, it fails to prod mode (which logs less, and thus is less risky).
func resolveEnvironment(v Environment) Environment {
	if v == EnvDevelopment {
		return EnvDevelopment
	}
	return EnvProduction
}
