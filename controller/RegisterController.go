package controller

import (
	"github.com/gin-gonic/gin"
	"go-dp.abdanhafidz.com/models"
	"go-dp.abdanhafidz.com/services"
)

func RegisterController(c *gin.Context) {
	register := services.RegisterService{}
	registerController := Controller[models.RegisterRequest, models.Account, models.Account]{
		Service: &register.Service,
	}
	registerController.RequestJSON(c)
	registerController.Service.Constructor.Username = registerController.Request.Username
	registerController.Service.Constructor.Password = registerController.Request.Password
	registerController.Service.Constructor.Name = registerController.Request.Name
	registerController.Service.Constructor.Email = registerController.Request.Email
	registerController.Service.Constructor.PhoneNumber = registerController.Request.Phone

	register.Create()
	registerController.Response(c)
}
