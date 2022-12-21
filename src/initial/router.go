package initial

import (
	"github.com/gin-gonic/gin"

	"XDCore/src/middlewares"
	"XDCore/src/router"
)

func InitRouters() *gin.Engine {
	Router := gin.Default()
	Router.Use(middlewares.Cors())

	apiGroup := Router.Group("/api")
	router.InitUserRouter(apiGroup)
	router.InitPlacementRouter(apiGroup)
	router.InitTagRouter(apiGroup)
	router.InitPerformanceRouter(apiGroup)

	return Router
}
