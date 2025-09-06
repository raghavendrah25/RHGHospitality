package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleHealthCheck(c *gin.Context) {
	data := map[string]any{
		"status": "OK",
	}

	c.JSON(http.StatusOK, data)
}

func HandleAbout(c *gin.Context) {
	data := map[string]any{
		"app":         "RHG Hospitality Solutions",
		"version":     "0.1.0",
		"environment": "Development	",
	}

	c.JSON(http.StatusOK, data)
}
