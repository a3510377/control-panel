package routers

import (
	"strings"

	"github.com/a3510377/control-panel/container"
	privateRouter "github.com/a3510377/control-panel/routers/private"
	publicRouter "github.com/a3510377/control-panel/routers/public"
	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	NoRouteHandlers []gin.HandlerFunc
}

func Routers(container *container.Container, config RouterConfig) *gin.Engine {
	router := gin.Default()

	router.NoRoute(append([]gin.HandlerFunc{
		func(c *gin.Context) {
			path := c.Request.URL.Path
			if strings.HasPrefix(path, "/api") {
				c.JSON(404, gin.H{"error": "not found"})
				c.Abort()
			}
		},
	}, config.NoRouteHandlers...)...)

	api := router.Group("/api")
	private := api.Group("/", container.Account.CheckFromRequest)

	privateRouter.PrivateRouter(container, private)
	publicRouter.PublicRouter(container, api)

	return router
}
