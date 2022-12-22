package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"XDCore/src/api"
	"XDCore/src/middlewares"
)

func InitTicketFileRouter(router *gin.RouterGroup) {
	subRouter := router.Group("ticketFile").Use(middlewares.JWTAuth())
	zap.S().Info("初始化 ticketFile 相关 router")
	{
		subRouter.POST("", api.CreateUpdateTicket)
	}
}
