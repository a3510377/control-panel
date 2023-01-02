package instances

import (
	"github.com/a3510377/control-panel/container"
	"github.com/gin-gonic/gin"
)

func AddHandler(container *container.Container, app *gin.RouterGroup) {
	instances(container, app)
}
