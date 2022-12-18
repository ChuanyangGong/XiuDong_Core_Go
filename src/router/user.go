package router

import (
	"XDCore/src/api"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitUserRouter(router *gin.RouterGroup) {
	userRouter := router.Group("user")

	zap.S().Info("初始化 user 相关 router")
	{
		userRouter.GET("list", api.GetUserList)
		userRouter.POST("login", api.PasswordLogin)
	}
}
