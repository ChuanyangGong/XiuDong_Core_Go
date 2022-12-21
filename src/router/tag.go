package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"XDCore/src/api"
	"XDCore/src/middlewares"
)

func InitTagRouter(router *gin.RouterGroup) {
	userRouter := router.Group("tag").Use(middlewares.JWTAuth())
	zap.S().Info("初始化 tag 相关 router")
	{
		userRouter.GET("list", api.GetTagList)
		userRouter.POST("", api.CreateUpdateTag)
		userRouter.PUT("", api.CreateUpdateTag)
		userRouter.DELETE(":id", api.DeleteTag)
	}
}
