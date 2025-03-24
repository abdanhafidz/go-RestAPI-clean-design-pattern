package controller

import (
	"github.com/gin-gonic/gin"
	"godp.abdanhafidz.com/controller"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/services"
)

func CreateVerification(c *gin.Context) {
	emailVerification := services.EmailVerificationService{}
	emailVerificationController := controller.Controller[any, models.EmailVerification, models.EmailVerification]{
		Service: &emailVerification.Service,
	}
	emailVerificationController.HeaderParse(c, func() {
		emailVerificationController.Service.Constructor.AccountID = uint(emailVerificationController.AccountData.UserID)
		emailVerification.Create()
		emailVerificationController.Response(c)
	})

}
