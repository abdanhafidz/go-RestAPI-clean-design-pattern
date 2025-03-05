package controller

import (
	"github.com/gin-gonic/gin"
	"go-dp.abdanhafidz.com/models"
	"go-dp.abdanhafidz.com/services"
	"go-dp.abdanhafidz.com/utils"
)

type (
	Controllers interface {
		RequestJSON(c *gin.Context)
		Response(c *gin.Context)
	}
	Controller[T1 any, T2 any, T3 any] struct {
		AccountData models.AccountData
		Request     T1
		Service     *services.Service[T2, T3]
	}
)

func (controller *Controller[T1, T2, T3]) RequestJSON(c *gin.Context) {
	cParam, _ := c.Get("accountData")
	if cParam != nil {
		controller.AccountData = cParam.(models.AccountData)
	}
	c.ShouldBindJSON(&controller.Request)
}
func (controller *Controller[T1, T2, T3]) Response(c *gin.Context) {
	switch {
	case controller.Service.Error != nil:
		utils.ResponseFAIL(c, 500, models.Exception{InternalServerError: true})
		utils.LogError(controller.Service.Error)
	case controller.Service.Exception.Unauthorized:
		utils.ResponseFAIL(c, 401, controller.Service.Exception)
	case controller.Service.Exception.DataNotFound:
		utils.ResponseFAIL(c, 404, controller.Service.Exception)
	case controller.Service.Exception.Message != "":
		utils.ResponseFAIL(c, 400, controller.Service.Exception)
	default:
		utils.ResponseOK(c, controller.Service.Result)
	}
}
