package controller

import "github.com/gin-gonic/gin"

func HomeController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Abdan Hafidz API 2025!",
	})
}
