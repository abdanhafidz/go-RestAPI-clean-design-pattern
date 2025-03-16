package router

import (
	"github.com/gin-gonic/gin"
	"go-dp.abdanhafidz.com/config"
)

func StartService() {
	router := gin.Default()
	UserRoutes(router)
	router.Run(config.TCP_ADDRESS)
}
