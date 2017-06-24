package main

import (
	"github.com/adbourne/go-archetype-rest/services"
	"github.com/adbourne/go-archetype-rest/config"
)

// Creates the application context
func NewAppContext(appConfig *config.AppConfig) *config.AppContext {
	return &config.AppContext{
		RandomNumberService: newRandomNumberService(appConfig),
	}
}

// Creates a random number service
func newRandomNumberService(appConfig *config.AppConfig) *services.DefaultRandomNumberService {
	return &services.DefaultRandomNumberService{
	}
}