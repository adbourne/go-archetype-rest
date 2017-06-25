package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"net/http"
	"github.com/adbourne/go-archetype-rest/config"
	"github.com/adbourne/go-archetype-rest/services"
)

func TestHealthEndpointReturns200Success(t *testing.T) {
	router := BuildGinRouter(buildTestAppContext())

	req, _ := http.NewRequest("GET", "/health", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	expectedMessage := "{\"status\":\"ok\"}"
	assert.Equal(t, resp.Code, 200)
	assert.Contains(t, resp.Body.String(), expectedMessage)
}

func buildTestAppContext() *config.AppContext {
	return &config.AppContext{
		AppConfig: buildTestAppConfig(),
		RandomNumberService: &services.DefaultRandomNumberService{
			Seed: 0,
		},
	}
}

func buildTestAppConfig() *config.AppConfig {
	return &config.AppConfig{
		Port: 0000,
		RandomSeed: int64(1),
	}
}
