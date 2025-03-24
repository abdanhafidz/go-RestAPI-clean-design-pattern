package router

import (
	"github.com/gin-gonic/gin"
	"godp.abdanhafidz.com/config"
	"godp.abdanhafidz.com/controller"
)

func StartService() {
	router := gin.Default()
	router.GET("/", controller.HomeController)
	UserRoute(router)
	EmailRoute(router)
	router.Run(config.TCP_ADDRESS)
}
