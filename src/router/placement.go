package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"XDCore/src/api"
	"XDCore/src/middlewares"
)

func InitPlacementRouter(router *gin.RouterGroup) {
	userRouter := router.Group("placement").Use(middlewares.JWTAuth())
	zap.S().Info("初始化 placement 相关 router")
	{
		userRouter.GET("list", api.GetPlacementList)
		userRouter.POST("", api.CreateUpdatePlacement)
		userRouter.PUT("", api.CreateUpdatePlacement)
		userRouter.DELETE(":id", api.DeletePlacement)
	}
}
