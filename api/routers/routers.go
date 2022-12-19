package routers

import (
	"github.com/a3510377/control-panel/container"
	privateRouter "github.com/a3510377/control-panel/routers/private"
	publicRouter "github.com/a3510377/control-panel/routers/public"
	"github.com/gin-gonic/gin"
)

func Routers(container *container.Container) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	private := api.Group("/", container.Account.CheckFromRequest)

	privateRouter.PrivateRouter(container, private)
	publicRouter.PublicRouter(container, api)

	return router
}
