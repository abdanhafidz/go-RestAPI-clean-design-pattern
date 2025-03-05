package utils

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"go-dp.abdanhafidz.com/models"
	"go-dp.abdanhafidz.com/services"
)

func ResponseOK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Request Success!", "data": data})
}

func ResponseFAIL(c *gin.Context, status int, exception models.Exception) {
	c.JSON(status, gin.H{"status": "error", "message": exception.Message, "error": exception})
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
