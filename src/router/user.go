package router

import (
	"XDCore/src/api"
	"XDCore/src/middlewares"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitUserRouter(router *gin.RouterGroup) {
	router.POST("login", api.PasswordLogin)

	userRouter := router.Group("user").Use(middlewares.JWTAuth())
	zap.S().Info("初始化 user 相关 router")
	{
		userRouter.GET("list", api.GetUserList)
	}
}
