package privateRouter

import (
	"github.com/a3510377/control-panel/container"
	"github.com/a3510377/control-panel/routers/private/instance"
	"github.com/gin-gonic/gin"
)

func PrivateRouter(container *container.Container, app *gin.RouterGroup) {
	// instances
	instance.AddHandler(container, app.Group("/instances"))
}
