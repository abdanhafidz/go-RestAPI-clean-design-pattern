package router

import (
	"github.com/gin-gonic/gin"
	"godp.abdanhafidz.com/factory"
	"godp.abdanhafidz.com/middleware"
)

func UserRoute(router *gin.Engine) {
	userModule := factory.NewUserProfileModule()
	routerGroup := router.Group("/api/v1/user")
	{
		routerGroup.GET("/me", middleware.AuthUser, userModule.GetProfile)
		routerGroup.PUT("/me", middleware.AuthUser, userModule.UpdateProfile)
	}
}
