package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gomcp/internal/models"
	"gomcp/internal/services"
)

// RegisterSummarizeRoutes registers summarize routes
func RegisterSummarizeRoutes(router *gin.Engine) {
	router.POST("/summarize", Summarize)
}

// Summarize handles POST /summarize
func Summarize(c *gin.Context) {
	var req models.SummarizeRequest

	// Bind JSON body to struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid JSON request",
		})
		return
	}

	// Basic validation
	if req.Text == "" && req.URL == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "either text or url must be provided",
		})
		return
	}

	// // Temporary response (we’ll replace this later)
	// response := models.SummarizeResponse{
	// 	Summary: "Summarization will be added later",
	// }

	// c.JSON(http.StatusOK, response)

	// If URL is provided, fetch content

	if req.URL != "" {
		htmlContent, err := services.FetchHTML(req.URL)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "failed to fetch URL content: " + err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"message":      "html fetched successfully",
			"html_length": len(htmlContent),
		})
		return
	}
	// Text path (temporary)
	c.JSON(http.StatusOK, models.SummarizeResponse{
		Summary: "text received, summarization coming soon",
	})
}
