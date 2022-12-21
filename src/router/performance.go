package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"XDCore/src/api"
	"XDCore/src/middlewares"
)

func InitPerformanceRouter(router *gin.RouterGroup) {
	userRouter := router.Group("performance").Use(middlewares.JWTAuth())
	zap.S().Info("初始化 performance 相关 router")
	{
		userRouter.GET("list", api.GetPlacementList)
		userRouter.POST("", api.CreateUpdatePerformance)
		userRouter.PUT("", api.CreateUpdatePerformance)
		userRouter.DELETE(":id", api.DeletePerformance)
	}
}
