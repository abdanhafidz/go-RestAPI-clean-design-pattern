package controller

import (
	"github.com/gin-gonic/gin"
	"go-dp.abdanhafidz.com/models"
	"go-dp.abdanhafidz.com/services"
)

func LoginController(c *gin.Context) {
	authentication := services.AuthenticationService{}
	loginController := Controller[models.LoginRequest, services.LoginConstructor, models.AuthenticatedUser]{
		Service: &authentication.Service,
	}
	loginController.RequestJSON(c, func() {
		loginController.Service.Constructor.Email = loginController.Request.Email
		loginController.Service.Constructor.Password = loginController.Request.Password
		authentication.Authenticate()
	})
}
