package router

import (
	"github.com/gin-gonic/gin"
	"go-dp.abdanhafidz.com/controller"
)

func UserRoutes(router *gin.Engine) {
	routerGroup := router.Group("/api/user")
	{
		routerGroup.GET("/", controller.HomeController)
		routerGroup.POST("/login", controller.LoginController)
		routerGroup.POST("/register", controller.RegisterController)
	}
}
