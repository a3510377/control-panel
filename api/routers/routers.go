package routers

import (
	privateRouter "github.com/a3510377/control-panel/routers/private"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	privateRouter.PrivateRouter(api)

	return router
}
