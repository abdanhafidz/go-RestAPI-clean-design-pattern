package router

import (
	"github.com/gin-gonic/gin"
	UserController "godp.abdanhafidz.com/controller/user"
	"godp.abdanhafidz.com/middleware"
)

func UserRoute(router *gin.Engine) {
	routerGroup := router.Group("/api/v1/user")
	{
		routerGroup.POST("/login", UserController.Login)
		routerGroup.POST("/register", UserController.Register)
		routerGroup.PUT("/change-password", middleware.AuthUser, UserController.ChangePassword)
		routerGroup.GET("/me", middleware.AuthUser, UserController.Profile)
		routerGroup.PUT("/me", middleware.AuthUser, UserController.UpdateProfile)
	}
}
