package utils

import (
	"github.com/gin-gonic/gin"
	"go-dp.abdanhafidz.com/models"
)

func GetAccount(c *gin.Context) models.AccountData {
	cParam, _ := c.Get("accountData")
	return cParam.(models.AccountData)
}
