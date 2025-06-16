package controller

import (
	"github.com/gin-gonic/gin"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/services"
)

type UserController interface {
	GetProfile(gCtx *gin.Context)
	UpdateProfile(gCtx *gin.Context)
}

type userController struct {
	*controller[services.UserProfileService]
}

func NewUserController(userProfileService services.UserProfileService) UserController {
	controller := userController{}
	controller.service = userProfileService
	return &controller
}
func (c *userController) GetProfile(ctx *gin.Context) {
	c.HeaderParse(ctx)
	userProfile := c.service.Retrieve(ctx.Request.Context(), uint(c.accountData.UserID))
	c.Response(ctx, userProfile)
}

func (c *userController) UpdateProfile(ctx *gin.Context) {
	var updateProfileRequest models.AccountDetails
	c.RequestJSON(ctx, updateProfileRequest)
	updatedProfile := c.service.Update(ctx.Request.Context(), uint(c.accountData.UserID), updateProfileRequest)
	c.Response(ctx, updatedProfile)
}
