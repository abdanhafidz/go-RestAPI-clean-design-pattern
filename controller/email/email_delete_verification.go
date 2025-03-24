package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"godp.abdanhafidz.com/controller"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/services"
)

func DeleteVerification(c *gin.Context) {
	emailVerification := services.EmailVerificationService{}
	emailVerificationController := controller.Controller[any, models.EmailVerification, models.EmailVerification]{
		Service: &emailVerification.Service,
	}
	query, _ := c.GetQuery("account_id")
	accountId, _ := strconv.Atoi(query)
	emailVerificationController.Service.Constructor.AccountID = uint(accountId)
	emailVerification.Delete()
	emailVerificationController.Response(c)
}
