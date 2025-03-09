package utils

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"go-dp.abdanhafidz.com/models"
	"go-dp.abdanhafidz.com/services"
)

func ResponseOK(c *gin.Context, data any) {
	res := models.SuccessResponse{
		Status:   "success",
		Message:  "Data retrieved successfully!",
		Data:     data,
		MetaData: c.Request.Body,
	}
	c.JSON(http.StatusOK, res)
	return
}

func ResponseFAIL(c *gin.Context, status int, exception models.Exception) {
	message := exception.Message
	exception.Message = ""
	res := models.ErrorResponse{
		Status:   "error",
		Message:  message,
		Errors:   exception,
		MetaData: c.Request.Body,
	}
	c.AbortWithStatusJSON(status, res)
	return
}

func SendResponse(c *gin.Context, data services.Service[any, any]) {
	if reflect.ValueOf(data.Exception).IsNil() {
		ResponseOK(c, data)
	} else {
		if data.Exception.Unauthorized {
			ResponseFAIL(c, 401, data.Exception)
		} else if data.Exception.BadRequest {
			ResponseFAIL(c, 400, data.Exception)
		} else if data.Exception.DataNotFound {
			ResponseFAIL(c, 404, data.Exception)
		} else if data.Exception.InternalServerError {
			ResponseFAIL(c, 500, data.Exception)
		} else {
			ResponseFAIL(c, 403, data.Exception)
		}
	}
}
