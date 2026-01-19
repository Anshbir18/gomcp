package handlers

import (
	"github.com/gin-gonic/gin"
)

func RegisterHealthRoutes(router *gin.Engine){
	router.GET("/health", HealthCheck)
}

// HealthCheck handles GET /health
func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}
