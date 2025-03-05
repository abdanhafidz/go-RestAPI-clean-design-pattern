package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go-dp.abdanhafidz.com/models"
	"go-dp.abdanhafidz.com/services"
)

func LoginController(c *gin.Context) {
	authentication := services.AuthenticationService{}
	loginController := Controller[models.LoginRequest, services.LoginService, models.Account]{
		Service: &authentication.Service,
	}
	loginController.RequestJSON(c)
	fmt.Println(loginController.Request)
	loginController.Service.Constructor.Username = loginController.Request.Username
	loginController.Service.Constructor.Password = loginController.Request.Password
	authentication.Authenticate()
	loginController.Response(c)
}
