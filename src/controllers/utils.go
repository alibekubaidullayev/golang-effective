package controllers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func readJSON(c *gin.Context, obj interface{}) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		slog.Error("Error binding JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error binding JSON"})
		return false
	}
	return true
}
