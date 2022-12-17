package initial

import (
	"XDCore/src/router"

	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	Router := gin.Default()
	apiGroup := Router.Group("/api")

	router.InitUserRouter(apiGroup)

	return Router
}
