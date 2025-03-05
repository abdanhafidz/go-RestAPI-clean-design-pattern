package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SendJSON200 sends a JSON response with HTTP status code 200
func SendJSON200(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": data})
}

// SendJSON400 sends a JSON response with HTTP status code 400
func SendJSON400(c *gin.Context, error_status *string, message *string) {
	c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error-status": error_status, "message": message})
}

// SendJSON401 sends a JSON response with HTTP status code 401
func SendJSON401(c *gin.Context, error_status *string, message *string) {
	c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "error-status": error_status, "message": message})
}

// SendJSON403 sends a JSON response with HTTP status code 403
func SendJSON403(c *gin.Context, message *string) {
	c.JSON(http.StatusForbidden, gin.H{"status": "error", "message": message})
}

// SendJSON404 sends a JSON response with HTTP status code 404
func SendJSON404(c *gin.Context, message *string) {
	c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": message})
}

// SendJSON500 sends a JSON response with HTTP status code 500
func SendJSON500(c *gin.Context, error_status *string, message *string) {
	c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error-status": error_status, "message": message})
}

// JSONResponseMiddleware is a middleware that provides functions for sending JSON responses
