package user

import (
	"github.com/gin-gonic/gin"
	"godp.abdanhafidz.com/controller"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/services"
)

func Register(c *gin.Context) {
	register := services.RegisterService{}
	registerController := controller.Controller[models.RegisterRequest, models.Account, models.Account]{
		Service: &register.Service,
	}
	registerController.RequestJSON(c, func() {
		registerController.Service.Constructor.Password = registerController.Request.Password
		registerController.Service.Constructor.Email = registerController.Request.Email
		register.Create()
	})
}
