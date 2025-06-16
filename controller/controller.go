package controller

import (
	"github.com/gin-gonic/gin"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/services"
	"godp.abdanhafidz.com/utils"
)

type (
	Controller interface {
		HeaderParse(ctx *gin.Context)
		RequestJSON(ctx *gin.Context, request any)
		Response(ctx *gin.Context, res any)
	}
	controller[TService services.Service] struct {
		accountData models.AccountData
		service     TService
	}
)

func (c *controller[TService]) HeaderParse(ctx *gin.Context) {
	cParam, _ := ctx.Get("account_data")
	if cParam != nil {
		c.accountData = cParam.(models.AccountData)
	}
}

func (c *controller[TService]) RequestJSON(ctx *gin.Context, request any) {
	cParam, _ := ctx.Get("AccountData")
	if cParam != nil {
		c.accountData = cParam.(models.AccountData)
	}

	errBinding := ctx.ShouldBindJSON(&request)
	if errBinding != nil {
		utils.ResponseFAIL(ctx, 400, models.Exception{
			BadRequest: true,
			Message:    "Invalid Request!, recheck your request, there's must be some problem about required parameter or type parameter",
		})
		return
	}
}
func (c *controller[TService]) Response(ctx *gin.Context, res any) {
	switch {
	case c.service.Error() != nil:
		utils.ResponseFAIL(ctx, 500, models.Exception{
			InternalServerError: true,
			Message:             "Internal Server Error",
		})
		utils.LogError(c.service.Error())
	case c.service.Exception().DataDuplicate:
		utils.ResponseFAIL(ctx, 400, c.service.Exception())
	case c.service.Exception().Unauthorized:
		utils.ResponseFAIL(ctx, 401, c.service.Exception())
	case c.service.Exception().DataNotFound:
		utils.ResponseFAIL(ctx, 404, c.service.Exception())
	case c.service.Exception().Message != "":
		utils.ResponseFAIL(ctx, 400, c.service.Exception())
	default:
		utils.ResponseOK(ctx, res)
	}
}
