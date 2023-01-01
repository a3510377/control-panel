package routers

import (
	"github.com/a3510377/control-panel/container"
	"github.com/a3510377/control-panel/routers/account"
	"github.com/gin-gonic/gin"
)

func loadRoutes(container *container.Container, app *gin.RouterGroup) {
	account.AddHandler(container, app)
}
