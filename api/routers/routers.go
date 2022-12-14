package routers

import (
	"github.com/a3510377/control-panel/container"
	privateRouter "github.com/a3510377/control-panel/routers/private"
	"github.com/gin-gonic/gin"
)

func Routers(db *container.Container) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	privateRouter.PrivateRouter(db, api)

	return router
}
