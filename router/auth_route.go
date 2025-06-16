package router

import (
	"github.com/gin-gonic/gin"
	"godp.abdanhafidz.com/factory"
)

func AuthRoute(router *gin.Engine) {
	authModule := factory.NewAuthenticationModule()
	routerGroup := router.Group("/api/v1/auth")
	{
		routerGroup.POST("/login", authModule.Login)
		routerGroup.POST("/register", authModule.Register)
		// routerGroup.PUT("/change-password", authModule.ChangePassword)
	}
}
