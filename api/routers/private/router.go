package privateRouter

import (
	"github.com/a3510377/control-panel/container"
	"github.com/a3510377/control-panel/routers/private/instance"
	"github.com/gin-gonic/gin"
)

func PrivateRouter(container *container.Container, app *gin.RouterGroup) {
	instanceRouter := app.Group("/instances")
	instance.AddHandler(container, instanceRouter)
}
