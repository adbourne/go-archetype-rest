package config

import "github.com/adbourne/go-archetype-rest/services"

// The application context
type AppContext struct {
	// The application config
	AppConfig           *AppConfig

	// The random number service
	RandomNumberService services.RandomNumberService
}
