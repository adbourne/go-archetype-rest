package config

import (
	"os"
	"log"
	"strconv"
)

const (
	EnvVarPort = "REST_ARCHETYPE_PORT"
	EnvVarRandomSeed = "REST_ARCHETYPE_RANDOM_SEED"
)


// The application configuration
type AppConfig struct {
	// The port to run the rest server on
	Port       int

	// The seed to use when generating randomness
	RandomSeed int64
}

// Constructs a new AppConfig
func NewAppConfig() *AppConfig {
	return &AppConfig{
		Port: loadEnvVarAsInt(EnvVarPort, 8080),
		RandomSeed: int64(loadEnvVarAsInt(EnvVarRandomSeed, 1)),
	}
}

// Utility function for loading an environment variable or use a default
func loadEnvVarAsInt(envVarName string, defaultVal int) int {
	ev, isFound := os.LookupEnv(envVarName)
	if !isFound {
		log.Printf("Environment variable '%s' not found, using default '%d'", envVarName, defaultVal)
		return defaultVal
	}

	evi, err := strconv.Atoi(ev)
	if err != nil {
		log.Fatalf("Environment variable '%s' was not a number", envVarName)
		return defaultVal
	}

	log.Printf("Environment variable '%s' found", envVarName)
	return evi
}