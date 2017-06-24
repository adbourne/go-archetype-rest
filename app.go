package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"log"
	"fmt"
	"github.com/adbourne/go-archetype-rest/config"
	"github.com/adbourne/go-archetype-rest/contollers"
)

func main() {
	printBanner()

	// Load the application config
	appConfig := config.NewAppConfig()

	// Create the application context
	appContext := NewAppContext(appConfig)

	// Build and run the gin router
	router := buildGinRouter(appContext)
	router.Run(fmt.Sprintf(":%d", appConfig.Port))
}

func printBanner() {
	log.Println(
		`
 ___ ___ ___ _____     _          _        _
| _ \ __/ __|_   _|   /_\  _ _ __| |_  ___| |_ _  _ _ __  ___
|   / _|\__ \ | |    / _ \| '_/ _| ' \/ -_)  _| || | '_ \/ -_)
|_|_\___|___/ |_|   /_/ \_\_| \__|_||_\___|\__|\_, | .__/\___|
                                               |__/|_|
		`)
}

func buildGinRouter(appContext *config.AppContext) *gin.Engine {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()
	router.HandleMethodNotAllowed = true

	// Add the health check endpoint
	controllers.HealthEndpoint(router)

	// Add the random number endpoint
	controllers.RandomNumberEndpoint(router, appContext)

	return router
}