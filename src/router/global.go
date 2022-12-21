package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"XDCore/src/api"
	"XDCore/src/middlewares"
)

func InitGlobalRouter(router *gin.RouterGroup) {
	commonRouter := router.Group("common").Use(middlewares.JWTAuth())
	zap.S().Info("初始化 performance 相关 router")
	{
		commonRouter.POST("uploadFile/:folder", api.UploadFiles)
	}
}
