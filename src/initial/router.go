package initial

import (
	"github.com/gin-gonic/gin"

	"XDCore/src/router"
)

func InitRouters() *gin.Engine {
	Router := gin.Default()
	apiGroup := Router.Group("/api")

	router.InitUserRouter(apiGroup)

	return Router
}
