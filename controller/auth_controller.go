package controller

import (
	"github.com/gin-gonic/gin"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/services"
)

type AuthController interface {
	Controller
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	// ChangePassword(ctx *gin.Context)
}

type authController struct {
	*controller[services.AuthenticationService]
}

func NewAuthController(authenticationService services.AuthenticationService) AuthController {
	controller := authController{}
	controller.service = authenticationService
	return &controller
}
func (c *authController) Login(ctx *gin.Context) {
	var loginRequest models.LoginRequest
	c.RequestJSON(ctx, &loginRequest)
	createdAccount := c.service.Authenticate(ctx.Request.Context(), loginRequest.Email, loginRequest.Password)
	c.Response(ctx, createdAccount)
}

func (c *authController) Register(ctx *gin.Context) {
	var registerRequest models.RegisterRequest
	c.RequestJSON(ctx, &registerRequest)
	createdAccount := c.service.Create(ctx.Request.Context(), registerRequest.Email, registerRequest.Password)
	c.Response(ctx, createdAccount)
}

// func (c *authController) ChangePassword(ctx *gin.Context) {
// 	var changePasswordRequest models.ChangePasswordRequest
// 	c.RequestJSON(ctx, &changePasswordRequest)
// 	c.Response(ctx)
// }
