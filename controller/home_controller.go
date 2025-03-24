package controller

import "github.com/gin-gonic/gin"

func HomeController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Api Qobiltu 2025!",
	})
}
