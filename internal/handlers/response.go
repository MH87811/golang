package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, data gin.H) {
	c.JSON(http.StatusOK, gin.H{"success": true, "data": data})
}

func Error(c *gin.Context, status int, code, message string) {
	c.JSON(status, gin.H{"success": false, "error": gin.H{"code": code, "message": message}})
}
