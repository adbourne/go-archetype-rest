package controllers

import "gopkg.in/gin-gonic/gin.v1"

// GET /health
// Health check endpoint, returns a 200 success when the service is up
func HealthEndpoint(router *gin.Engine) gin.IRoutes {
	return router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
}

