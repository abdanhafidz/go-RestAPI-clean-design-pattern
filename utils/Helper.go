package utils

import (
	"github.com/gin-gonic/gin"
	"godp.abdanhafidz.com/models"
)

func GetAccount(c *gin.Context) models.AccountData {
	cParam, _ := c.Get("accountData")
	return cParam.(models.AccountData)
}
