package controllers

import (
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/adbourne/go-archetype-rest/config"
)

// GET /
// Random number endpoint, returns a 200 success with a random number as a body
func RandomNumberEndpoint(router *gin.Engine, appContext *config.AppContext) gin.IRoutes {
	return router.GET("/", func(c *gin.Context) {
		randomNumber := appContext.RandomNumberService.GenerateRandomNumber()
		c.JSON(200, gin.H{
			"randomNumber": randomNumber,
		})
	})
}
