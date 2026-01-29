package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gomcp/internal/handlers"
)

func main() {
	// Create a Gin router with default middleware:
	// logger and recovery (crash-free) middleware

	router := gin.Default()

	// Register health check route
	handlers.RegisterHealthRoutes(router)
	// Register summarize route
	handlers.RegisterSummarizeRoutes(router)

	// Start the server on port 8080
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
